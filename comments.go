package xlsx

import (
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"path"
	"path/filepath"
	"strings"
)

type comments struct {
	sheet        *sheetInfo
	ml           ml.Comments
	vml          vml.Excel
	fileComments *ooxml.PackageFile
	fileDrawings *ooxml.PackageFile
}

//newComments creates an object that implements comments functionality
func newComments(sheet *sheetInfo) *comments {
	return &comments{sheet: sheet}
}

func (c *comments) initIfRequired() {
	//attach sheet relations file
	c.sheet.attachRelationshipsIfRequired()
	//attach comments file
	c.attachCommentsIfRequired()
	//attach VML file
	c.attachDrawingsIfRequired()
}

func (c *comments) attachCommentsIfRequired() {
	if c.fileComments == nil {
		fileName := c.sheet.relationships.GetTargetByType(internal.RelationTypeComments)
		if fileName != "" {
			fileName = strings.Replace(fileName, "../", "xl/", 1)
		} else {
			fileName = c.sheet.workbook.doc.relationships.GetTargetById(string(c.sheet.workbook.ml.Sheets[c.sheet.index].RID))
			fileName = fmt.Sprintf("xl/comments.%s.xml", strings.TrimSuffix(filepath.Base(fileName), path.Ext(fileName)))
		}

		if file := c.sheet.workbook.doc.pkg.File(fileName); file != nil {
			c.fileComments = ooxml.NewPackageFile(c.sheet.workbook.doc.pkg, file, &c.ml, nil)
			c.fileComments.LoadIfRequired(nil)
		} else {
			//register a new comments content
			c.sheet.workbook.doc.pkg.ContentTypes().RegisterContent(fileName, internal.ContentTypeComments)

			//attach file to package
			c.fileComments = ooxml.NewPackageFile(c.sheet.workbook.doc.pkg, fileName, &c.ml, nil)

			//add file to sheet relations
			c.sheet.relationships.AddFile(internal.RelationTypeComments, fileName)
			c.fileComments.MarkAsUpdated()
		}
	}
}

func (c *comments) attachDrawingsIfRequired() {
	if c.fileDrawings == nil {
		fileName := c.sheet.relationships.GetTargetByType(internal.RelationTypeVmlDrawing)
		if fileName != "" {
			fileName = strings.Replace(fileName, "../", "xl/", 1)
		} else {
			fileName = c.sheet.workbook.doc.relationships.GetTargetById(string(c.sheet.workbook.ml.Sheets[c.sheet.index].RID))
			fileName = fmt.Sprintf("xl/drawings/vmlDrawings.%s.vml", strings.TrimSuffix(filepath.Base(fileName), path.Ext(fileName)))
		}

		if file := c.sheet.workbook.doc.pkg.File(fileName); file != nil {
			c.fileDrawings = ooxml.NewPackageFile(c.sheet.workbook.doc.pkg, file, &c.vml, nil)
			c.fileDrawings.LoadIfRequired(nil)
		} else {
			//register a VML content type, if required
			c.sheet.workbook.doc.pkg.ContentTypes().RegisterType("vml", ooxml.ContentTypeVmlDrawing)

			//attach file to package
			c.fileDrawings = ooxml.NewPackageFile(c.sheet.workbook.doc.pkg, fileName, &c.vml, nil)

			//add file to sheet relations
			_, rid := c.sheet.relationships.AddFile(internal.RelationTypeVmlDrawing, fileName)

			//add legacy drawing
			c.sheet.ml.LegacyDrawing = &ml.LegacyDrawing{RID: rid}
		}
	}
}

func (c *comments) Add(bounds types.Bounds, comment interface{}) error {
	c.initIfRequired()
	c.fileComments.MarkAsUpdated()
	c.fileDrawings.MarkAsUpdated()
	return nil
}

//Remove removes comment info for bounds
func (c *comments) Remove(bounds types.Bounds) {
	c.initIfRequired()
	c.fileComments.MarkAsUpdated()
	c.fileDrawings.MarkAsUpdated()
}
