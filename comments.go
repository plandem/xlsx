package xlsx

import (
	"github.com/plandem/xlsx/types"
)

type comments struct {
	sheet *sheetInfo
}

//newComments creates an object that implements comments functionality
func newComments(sheet *sheetInfo) *comments {
	return &comments{sheet: sheet}
	//ss.file = ooxml.NewPackageFile(doc.pkg, f, &ss.ml, nil)
	//
	//if ss.file.IsNew() {
	//	ss.doc.pkg.ContentTypes().RegisterContent(ss.file.FileName(), internal.ContentTypeStyles)
	//	ss.doc.relationships.AddFile(internal.RelationTypeStyles, ss.file.FileName())
	//	ss.file.MarkAsUpdated()
	//	ss.addDefaults()
	//	ss.buildIndexes()
	//}

	//TODO:
	//1) add file to [Content_Types].xml
	//2) add link to that file in sheet/relations?
	//3) add link to drawing in sheet/relations?
	//4) add legacyDrawing in sheet?
}

func (c *comments) initIfRequired() {
	////attach hyperlinks object if required
	//if h.sheet.ml.Hyperlinks == nil {
	//	var links []*ml.Hyperlink
	//	h.sheet.ml.Hyperlinks = &links
	//}

	//c.sheet.attachRelationshipsIfRequired()
	//c.sheet.relationships.GetIdByTarget()
}

func (c *comments)  Add(bounds types.Bounds, comment interface{}) error {
	c.initIfRequired()

	return nil
}

//Remove removes comment info for bounds
func (c *comments) Remove(bounds types.Bounds) {
	c.initIfRequired()
}
