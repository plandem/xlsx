// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
)

//workbook is a higher level object that wraps ml.Workbook with functionality
type workbook struct {
	ml   ml.Workbook
	doc  *Spreadsheet
	file *ooxml.PackageFile
}

func newWorkbook(f interface{}, doc *Spreadsheet) *workbook {
	wb := &workbook{
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
