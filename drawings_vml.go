package xlsx

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/drawing/vml"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type drawingsVML struct {
	sheet               *sheetInfo
	ml                  vml.Excel
	chunks              []int
	file                *ooxml.PackageFile
	initializedFile     bool
	initializedComments bool
	updated             bool
	nextShapeId         int
	nextShapeIdMax      int
}

//capacity of chunk. vml file store shapes in chunks
const vmlChunkSize = 1024

//office's internal id of vml shape type for comments
const commentShapeTypeSpt = 202

var (
	regexpDrawings = regexp.MustCompile(`xl/drawings/[[:alpha:]]+[\d]+\.vml`)
	regexpShapeID  = regexp.MustCompile(`_x0000_s([\d]+)`)
)

func newDrawingsVML(sheet *sheetInfo) *drawingsVML {
	return &drawingsVML{sheet: sheet}
}

//resolve chunks info of this VML drawings file
func (d *drawingsVML) resolveChunks() {
	if len(d.chunks) == 0 {
		d.attachFileIfRequired()

		var idMap *vml.IdMap

		//only non loaded existing files will return stream, otherwise chunks info can be gathered from existing ml info
		if stream, err := d.file.ReadStream(); err == nil {
			//locate chunks info
			for next, hasNext := stream.StartIterator(nil); hasNext; {
				hasNext = next(func(decoder *xml.Decoder, start *xml.StartElement) bool {
					if start.Name.Local == "shapelayout" {
						shapeLayout := &vml.ShapeLayout{}
						if err := decoder.DecodeElement(shapeLayout, start); err != nil {
							_ = stream.Close()
							panic(err)
						}

						idMap = shapeLayout.IdMap
						return false
					}

					return true
				})
			}

			_ = stream.Close()
		} else {
			idMap = d.ml.ShapeLayout.IdMap
		}

		//parse chunks info
		if idMap.Data != "" {
			chunks := strings.Split(idMap.Data, ",")
			for _, s := range chunks {
				if n, err := strconv.Atoi(strings.TrimSpace(s)); err != nil {
					panic(fmt.Errorf("can't load chunks info from VML idmap: %s", err))
				} else {
					d.chunks = append(d.chunks, n)
				}
			}
		}
	}
}

func (d *drawingsVML) nextChunkID() int {
	nextChunk := 0

	//get maximum chunk - each VML file can have few non serial chunks (1024 shapes per chunk), e.g.: 1,3,9
	for _, s := range d.sheet.workbook.doc.sheets {
		s.drawingsVML.resolveChunks()

		for _, chunkID := range s.drawingsVML.chunks {
			nextChunk = int(math.Max(float64(nextChunk), float64(chunkID)))
		}
	}

	nextChunk++
	d.chunks = append(d.chunks, nextChunk)
	d.nextShapeId = nextChunk * vmlChunkSize
	d.nextShapeIdMax = d.nextShapeId + vmlChunkSize

	return nextChunk
}

func (d *drawingsVML) nextShapeID() int {
	if d.nextShapeId >= d.nextShapeIdMax {
		d.nextChunkID()
	}

	d.nextShapeId++
	return d.nextShapeId
}

func (d *drawingsVML) addComment(bounds types.Bounds, comment interface{}) error {
	d.initCommentsIfRequired()

	//TODO: replace mock data
	shape := &vml.Shape{}
	shape.ID = fmt.Sprintf("_x0000_s%d", d.nextShapeID())
	shape.Type = fmt.Sprintf("#_x0000_t%d", commentShapeTypeSpt)
	shape.FillColor = "#ffffe1"
	shape.InsetMode = vml.InsetModeAuto
	shape.Style = "position:absolute;margin-left:242.25pt;margin-top:22.5pt;width:96pt;height:55.5pt;z-index:1;visibility:hidden"
	shape.Fill = &vml.Fill{Color2: "#ffffe1"}
	shape.PathSettings = &vml.Path{ConnectType: vml.ConnectTypeNone}
	shape.Shadow = &vml.Shadow{
		Color:    "black",
		On:       sharedML.TriStateTrue,
		Obscured: sharedML.TriStateTrue,
	}
	shape.ClientData = &vml.ClientData{
		Row:           bounds.FromRow,
		Column:        bounds.FromCol,
		Type:          vml.ObjectTypeNote,
		Anchor:        "3, 15, 1, 10, 5, 15, 2, 64",
		SizeWithCells: sharedML.TriStateBlankTrue(sharedML.TriStateTrue),
		MoveWithCells: sharedML.TriStateBlankTrue(sharedML.TriStateTrue),
		AutoFill:      sharedML.TriStateBlankTrue(sharedML.TriStateFalse),
	}

	//colCount = max(line length)
	//lineCount = number of lines + 1
	//Row:      xAxis,
	//Column:   yAxis,
	//Anchor: 1+yAxis, 1+xAxis, 2+yAxis+lineCount, colCount+yAxis, 2+xAxis+lineCount),

	d.ml.Shape = append(d.ml.Shape, shape)
	d.file.MarkAsUpdated()
	d.updated = true
	return nil
}

//Remove removes comment info for bounds
func (d *drawingsVML) removeComment(bounds types.Bounds) {
	d.initCommentsIfRequired()

	//TODO: remove comments
}

//load all content if required or add minimal required
func (d *drawingsVML) initIfRequired() {
	if d.initializedFile {
		return
	}

	d.attachFileIfRequired()

	if !d.file.IsNew() {
		d.file.LoadIfRequired(nil)

		//resolve chunks info
		d.resolveChunks()

		//resolve next shapeID info
		nextShapeID := 0
		for _, shape := range d.ml.Shape {
			if matched := regexpShapeID.FindSubmatch([]byte(shape.ID)); len(matched) > 0 {
				if id, err := strconv.Atoi(string(matched[1])); err != nil {
					panic(fmt.Errorf("can't get ID of shape: %s", matched))
				} else {
					//TODO: theoretically we should take maximum shape_id of the lowest chunk from file as possible.
					// But it's premature optimization, so right now we just take maximum ID of shape, even "lower" chunks still can have more shapes.
					// E.g.: file had few chunks, but later few shapes were deleted from chunk1, so theoretically we could add next shapes to chunk1 again
					nextShapeID = int(math.Max(float64(nextShapeID), float64(id)))
				}
			}
		}

		d.nextShapeId = nextShapeID
		d.nextShapeIdMax = (nextShapeID/vmlChunkSize)*vmlChunkSize + vmlChunkSize
	} else {
		d.ml.ShapeLayout = &vml.ShapeLayout{
			Ext: vml.ExtTypeEdit,
			IdMap: &vml.IdMap{
				Ext: vml.ExtTypeEdit,
			},
		}

		//our resolving mechanism of next chunk relies on non nil object, that's why we should assign it manually
		d.ml.ShapeLayout.IdMap.Data = strconv.Itoa(d.nextChunkID())
	}

	d.file.MarkAsUpdated()
	d.initializedFile = true
	d.updated = true
}

//attach LegacyDrawing info into sheet
func (d *drawingsVML) attachDrawingsRID() {
	if d.updated && (d.sheet.ml.LegacyDrawing == nil || d.sheet.ml.LegacyDrawing.RID == "") {
		fileName := d.sheet.relationships.GetTargetByType(internal.RelationTypeVmlDrawing)
		rid := d.sheet.relationships.GetIdByTarget(fileName)
		d.sheet.ml.LegacyDrawing = &ml.LegacyDrawing{RID: rid}
	}
}

//add shape type for comments if required
func (d *drawingsVML) initCommentsIfRequired() {
	if d.initializedComments {
		return
	}

	d.initIfRequired()

	//attach shape type if required
	for _, shapeType := range d.ml.ShapeType {
		if shapeType.Spt == commentShapeTypeSpt {
			return
		}
	}

	shapeType := &vml.ShapeType{}
	shapeType.ID = fmt.Sprintf("_x0000_t%d", commentShapeTypeSpt)
	shapeType.Spt = commentShapeTypeSpt
	shapeType.Path = "m,l,21600r21600,l21600,xe"
	shapeType.CoordSize = "21600,21600"

	shapeType.Stroke = &vml.Stroke{}
	shapeType.Stroke.JoinStyle = vml.StrokeJoinStyleMiter

	shapeType.PathSettings = &vml.Path{}
	shapeType.PathSettings.GradientShapeOK = sharedML.TriStateTrue
	shapeType.PathSettings.ConnectType = vml.ConnectTypeRect

	d.ml.ShapeType = append(d.ml.ShapeType, shapeType)
	d.initializedComments = true
}

//only attach files, no content is loading
func (d *drawingsVML) attachFileIfRequired() {
	//attach sheet relations file
	d.sheet.attachRelationshipsIfRequired()

	if d.file == nil {
		fileName := d.sheet.relationships.GetTargetByType(internal.RelationTypeVmlDrawing)
		if fileName != "" {
			//transform relative path to absolute
			fileName = strings.Replace(fileName, "../", "xl/", 1)

			if file := d.sheet.workbook.doc.pkg.File(fileName); file != nil {
				d.file = ooxml.NewPackageFile(d.sheet.workbook.doc.pkg, file, &d.ml, nil)
				return
			}

			panic(fmt.Sprintf("can't load VML file: %s", fileName))
		}

		totalFiles := len(d.sheet.workbook.doc.pkg.Files(regexpDrawings))
		fileName = fmt.Sprintf("xl/drawings/vmlDrawing%d.vml", totalFiles+1)

		//register a VML content type, if required
		d.sheet.workbook.doc.pkg.ContentTypes().RegisterType("vml", ooxml.ContentTypeVmlDrawing)

		//attach file to package
		d.file = ooxml.NewPackageFile(d.sheet.workbook.doc.pkg, fileName, &d.ml, nil)

		//add file to sheet relations
		d.sheet.relationships.AddFile(internal.RelationTypeVmlDrawing, fileName)
	}
}
