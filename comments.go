package xlsx

import (
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"strings"
)

type comments struct {
	sheet *sheetInfo
	ml    ml.Comments
	file  *ooxml.PackageFile
}

//newComments creates an object that implements comments functionality
func newComments(sheet *sheetInfo) *comments {
	return &comments{sheet: sheet}
}

func (c *comments) initIfRequired() {
	//attach comments file
	c.attachFileIfRequired()

	//attach default author if required
	if len(c.ml.Authors) == 0 {
		c.ml.Authors = append(c.ml.Authors, "")
	}
}

//only attach files, no content is loading
func (c *comments) attachFileIfRequired() {
	//attach sheet relations file
	c.sheet.attachRelationshipsIfRequired()

	if c.file == nil {
		fileName := c.sheet.relationships.GetTargetByType(internal.RelationTypeComments)
		if fileName != "" {
			//transform relative path to absolute
			fileName = strings.Replace(fileName, "../", "xl/", 1)

			if file := c.sheet.workbook.doc.pkg.File(fileName); file != nil {
				c.file = ooxml.NewPackageFile(c.sheet.workbook.doc.pkg, file, &c.ml, nil)
				return
			}

			panic(fmt.Sprintf("can't load comments file: %s", fileName))
		}

		totalFiles := c.sheet.workbook.doc.pkg.ContentTypes().CountTypes(internal.ContentTypeComments)
		fileName = fmt.Sprintf("xl/comments%d.xml", totalFiles+1)

		//register a new comments content
		c.sheet.workbook.doc.pkg.ContentTypes().RegisterContent(fileName, internal.ContentTypeComments)

		//attach file to package
		c.file = ooxml.NewPackageFile(c.sheet.workbook.doc.pkg, fileName, &c.ml, nil)

		//add file to sheet relations
		c.sheet.relationships.AddFile(internal.RelationTypeComments, fileName)
		c.file.MarkAsUpdated()
	}
}

//Add adds a new comment info for bounds
func (c *comments) Add(bounds types.Bounds, comment interface{}) error {
	c.initIfRequired()

	//TODO: replace mock data
	cml := &ml.Comment{}
	cml.Ref = bounds
	cml.AuthorID = 0
	cml.Text = &ml.StringItem{
		Text: "My Comment",
	}

	c.ml.CommentList = append(c.ml.CommentList, cml)
	c.file.MarkAsUpdated()

	return c.sheet.drawingsVML.addComment(bounds, comment)
}

//Remove removes comment info for bounds
func (c *comments) Remove(bounds types.Bounds) {
	c.initIfRequired()
	c.file.MarkAsUpdated()
	//remove comment

	//remove VML drawings
	c.sheet.drawingsVML.removeComment(bounds)
}
