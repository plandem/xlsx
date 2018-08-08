package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//SharedStrings is a higher level object that wraps ml.SharedStrings with functionality
type SharedStrings struct {
	ml    ml.SharedStrings
	index map[string]int //TODO: need optimization, currently we holds 2 version in memory ('slice' at ml + 'map' for indexes)
	doc   *Spreadsheet
	file  *ooxml.PackageFile
}

func newSharedStrings(f interface{}, doc *Spreadsheet) *SharedStrings {
	ss := &SharedStrings{
		doc:   doc,
		index: make(map[string]int),
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
		ss.index[string(s.Text)] = i
	}
}

//get returns string stored at index
func (ss *SharedStrings) get(index int) string {
	ss.file.LoadIfRequired(ss.afterLoad)

	var value string
	if index < len(ss.ml.StringItem) {
		value = string(ss.ml.StringItem[index].Text)
	}

	return value
}

//add add a new value and return index for it
func (ss *SharedStrings) add(value string) int {
	ss.file.LoadIfRequired(ss.afterLoad)

	//return sid if already exists
	if sid, ok := ss.index[value]; ok {
		return sid
	}

	//add a new one if there is no such string
	sid := len(ss.ml.StringItem)
	ss.ml.StringItem = append(ss.ml.StringItem, &ml.StringItem{Text: primitives.Text(value)})
	ss.index[value] = sid

	ss.file.MarkAsUpdated()

	return sid
}
