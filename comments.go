package xlsx

import (
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/hash"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"github.com/plandem/xlsx/types/comment"
	"strings"
)

type comments struct {
	sheet        *sheetInfo
	ml           ml.Comments
	file         *ooxml.PackageFile
	authorIndex  map[hash.Code]int
	commentIndex map[hash.Code]int
}

//newComments creates an object that implements comments functionality
func newComments(sheet *sheetInfo) *comments {
	return &comments{
		sheet:        sheet,
		authorIndex:  make(map[hash.Code]int),
		commentIndex: make(map[hash.Code]int),
	}
}

func (c *comments) initIfRequired() {
	//attach sheet relations file
	c.sheet.attachRelationshipsIfRequired()

	if c.file == nil {
		fileName := c.sheet.relationships.GetTargetByType(internal.RelationTypeComments)
		if fileName != "" {
			//transform relative path to absolute
			fileName = strings.Replace(fileName, "../", "xl/", 1)

			if file := c.sheet.workbook.doc.pkg.File(fileName); file != nil {
				c.file = ooxml.NewPackageFile(c.sheet.workbook.doc.pkg, file, &c.ml, nil)
				c.file.LoadIfRequired(nil)
				c.addDefaults()
				c.buildIndexes()
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
		c.addDefaults()
		c.buildIndexes()
	}
}

//build author and comment indexes
func (c *comments) buildIndexes() {
	for id, f := range c.ml.Authors {
		c.authorIndex[hash.Key(f).Hash()] = id
	}

	for id, f := range c.ml.CommentList {
		c.commentIndex[hash.Key(f.Ref.String()).Hash()] = id
	}
}

func (c *comments) addDefaults() {
	//attach default author if required
	if len(c.ml.Authors) == 0 {
		c.ml.Authors = append(c.ml.Authors, "")
	}
}

//Add adds a new comment info for bounds
func (c *comments) Add(bounds types.Bounds, info interface{}) error {
	c.initIfRequired()

	//check if there is comment already for these bounds
	commentKey := hash.Key(bounds.String()).Hash()
	if _, ok := c.commentIndex[commentKey]; ok {
		return fmt.Errorf("there is already comment for ref %s", bounds.String())
	}

	//resolve Info if required
	var object *comment.Info
	if text, ok := info.(string); ok {
		object = comment.New(comment.Text(text))
	} else if pointer, ok := info.(*comment.Info); ok {
		object = pointer
	} else if value, ok := info.(comment.Info); ok {
		object = &value
	} else {
		return fmt.Errorf("unsupported type of comment, only string or comment.Info is allowed")
	}

	cml := &ml.Comment{}

	//resolve id of author
	authorKey := hash.Key(object.Author).Hash()
	if id, ok := c.authorIndex[authorKey]; ok {
		cml.AuthorID = id
	} else {
		nextID := len(c.ml.Authors)
		c.authorIndex[authorKey] = nextID
		cml.AuthorID = nextID
	}

	cml.Ref = bounds
	if text, err := toRichText(object.Text...); err != nil {
		return err
	} else {
		cml.Text = text
	}

	c.commentIndex[commentKey] = len(c.ml.CommentList)
	c.ml.CommentList = append(c.ml.CommentList, cml)
	c.file.MarkAsUpdated()

	return c.sheet.drawingsVML.addComment(bounds, object)
}

//Remove removes comment info for bounds
func (c *comments) Remove(bounds types.Bounds) {
	c.initIfRequired()
	c.file.MarkAsUpdated()

	key := hash.Key(bounds.String()).Hash()
	if id, ok := c.commentIndex[key]; ok {
		c.ml.CommentList[id] = c.ml.CommentList[len(c.ml.CommentList)-1]
		c.ml.CommentList[len(c.ml.CommentList)-1] = nil //prevent memory leaks
		c.ml.CommentList = c.ml.CommentList[:len(c.ml.CommentList)-1]

		//remove VML drawings
		c.sheet.drawingsVML.removeComment(bounds)

		//clean up indexes
		delete(c.commentIndex, key)
	}
}
