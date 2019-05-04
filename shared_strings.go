package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/hash"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//SharedStrings is a higher level object that wraps ml.SharedStrings with functionality
type SharedStrings struct {
	ml    ml.SharedStrings
	index map[hash.Code]int
	doc   *Spreadsheet
	file  *ooxml.PackageFile
}

func newSharedStrings(f interface{}, doc *Spreadsheet) *SharedStrings {
	ss := &SharedStrings{
		doc:   doc,
		index: make(map[hash.Code]int),
	}

	ss.file = ooxml.NewPackageFile(doc.pkg, f, &ss.ml, nil)

	if ss.file.IsNew() {
		ss.doc.pkg.ContentTypes().RegisterContent(ss.file.FileName(), internal.ContentTypeSharedStrings)
		ss.doc.relationships.AddFile(internal.RelationTypeSharedStrings, ss.file.FileName())
		ss.file.MarkAsUpdated()
	}

	return ss
}

func (ss *SharedStrings) afterLoad() {
	for i, s := range ss.ml.StringItem {
		ss.index[hash.StringItem(s).Hash()] = i
	}
}

//get returns string item stored at index
func (ss *SharedStrings) get(index int) *ml.StringItem {
	ss.file.LoadIfRequired(ss.afterLoad)

	if index < len(ss.ml.StringItem) {
		return ss.ml.StringItem[index]
	}

	return nil
}

//addString adds a new string and return index for it
func (ss *SharedStrings) addString(value string) int {
	return ss.addText(&ml.StringItem{Text: primitives.Text(value)})
}

//addText adds a new StringItem and return index for it
func (ss *SharedStrings) addText(si *ml.StringItem) int {
	ss.file.LoadIfRequired(ss.afterLoad)

	key := hash.StringItem(si).Hash()

	//return sid if already exists
	if sid, ok := ss.index[key]; ok {
		return sid
	}

	sid := len(ss.ml.StringItem)
	ss.ml.StringItem = append(ss.ml.StringItem, si)
	ss.index[key] = sid

	ss.file.MarkAsUpdated()

	return sid
}
