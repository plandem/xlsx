package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
)

//Workbook is a higher level object that wraps ml.Workbook with functionality
type Workbook struct {
	ml   ml.Workbook
	doc  *Spreadsheet
	file *ooxml.PackageFile
}

func newWorkbook(f interface{}, doc *Spreadsheet) *Workbook {
	wb := &Workbook{
		doc: doc,
	}

	doc.workbook = wb

	wb.file = ooxml.NewPackageFile(doc.pkg, f, &wb.ml, nil)
	wb.file.LoadIfRequired(nil)

	if wb.file.IsNew() {
		doc.pkg.ContentTypes().RegisterContent(wb.file.FileName(), internal.ContentTypeWorkbook)
		doc.pkg.Relationships().AddFile(internal.RelationTypeWorkbook, wb.file.FileName())
		wb.file.MarkAsUpdated()
	}

	return wb
}
