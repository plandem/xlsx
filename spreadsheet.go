package xlsx

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"regexp"
)

//Spreadsheet is a higher level object that wraps OOXML package with XLSX functionality
type Spreadsheet struct {
	ooxml.Package
	pkg           *ooxml.PackageInfo
	workbook      *Workbook
	sheets        []*Sheet
	relationships *ooxml.Relationships
	sharedStrings *SharedStrings
	styleSheet    *StyleSheet
}

//newSpreadsheet creates an object that implements XLSX functionality
func newSpreadsheet(pkg *ooxml.PackageInfo) (interface{}, error) {
	xlDoc := &Spreadsheet{
		pkg:     pkg,
		Package: pkg,
	}

	pkg.Validator = xlDoc.IsValid

	if pkg.IsNew() {
		xlDoc.createSpreadsheet()
	} else {
		xlDoc.readSpreadsheet()
	}

	return xlDoc, nil
}

//GetSheetNames returns a names of all sheets
func (xl *Spreadsheet) GetSheetNames() []string {
	sheetNames := make([]string, len(xl.sheets))

	for id := range xl.sheets {
		sheetNames[id] = xl.workbook.ml.Sheets[id].Name
	}

	return sheetNames
}

//Sheet returns a sheet by 0-based index
func (xl *Spreadsheet) Sheet(i int) *Sheet {
	var sheet *Sheet

	if i < len(xl.sheets) {
		sheet = xl.sheets[i]
		sheet.file.LoadIfRequired(sheet.expandOnInit)
		sheet.file.MarkAsUpdated()
	}

	return sheet
}

//Sheets returns iterator for all sheets of Spreadsheet
func (xl *Spreadsheet) Sheets() SheetIterator {
	return newSheetIterator(xl)
}

//DeleteSheet deletes the sheet with required 0-based index
func (xl *Spreadsheet) DeleteSheet(i int) {
	if i < len(xl.sheets) {
		sheet := xl.sheets[i]
		if sheet != nil {
			rid := xl.workbook.ml.Sheets[i].RID

			//remove from document
			xl.sheets = append(xl.sheets[:i], xl.sheets[i+1:]...)

			//remove from workbook
			xl.workbook.ml.Sheets = append(xl.workbook.ml.Sheets[:i], xl.workbook.ml.Sheets[i+1:]...)
			xl.workbook.file.MarkAsUpdated()

			//remove relation
			xl.relationships.Remove(rid)

			//remove file
			xl.pkg.Remove(sheet.file.FileName())
		}
	}
}

//AddSheet adds a new sheet with name to document
func (xl *Spreadsheet) AddSheet(name string) *Sheet {
	var sheet *Sheet

	if sheet = newSheet(fmt.Sprintf("xl/worksheets/sheet%d.xml", len(xl.workbook.ml.Sheets)+1), xl); sheet != nil {
		if len(name) > 0 {
			sheet.SetName(name)
		}

		sheet.file.MarkAsUpdated()
		xl.workbook.file.MarkAsUpdated()
		xl.pkg.ContentTypes().RegisterContent(sheet.file.FileName(), internal.ContentTypeWorksheet)
	}

	sheet.expandOnInit()
	return sheet
}

//SetActive sets the sheet with 0-based index as active
func (xl *Spreadsheet) SetActive(i int) {
	if i >= len(xl.sheets) {
		return
	}

	//set activate from workbook side
	if xl.workbook.ml.BookViews == nil || len(*xl.workbook.ml.BookViews) == 0 {
		xl.workbook.ml.BookViews = &[]*ml.BookView{{
			ActiveTab: i,
		}}
	} else {
		(*xl.workbook.ml.BookViews)[0].ActiveTab = i
	}

	//set active from worksheet side
	for id := range xl.sheets {
		sheet := xl.Sheet(id)
		if sheet.ml.SheetViews != nil && len(sheet.ml.SheetViews.SheetView) > 0 {
			sheet.ml.SheetViews.SheetView[0].TabSelected = i == id
		}
	}

	xl.workbook.file.MarkAsUpdated()
}

//AddFormatting adds a new style formatting to document and return related ID that can be used lately
func (xl *Spreadsheet) AddFormatting(style *format.StyleFormat) format.StyleRefID {
	return xl.styleSheet.addXF(style)
}

//IsValid validates document and return error if there is any error. Using right before saving.
func (xl *Spreadsheet) IsValid() error {
	if len(xl.sheets) == 0 {
		return errors.New("Spreadsheet requires at least one worksheet.")
	}

	return nil
}

//readSpreadsheet reads required information from XLSX
func (xl *Spreadsheet) readSpreadsheet() {
	files := xl.pkg.Files()
	for _, file := range files {
		if f, ok := file.(*zip.File); ok {
			switch {
			case f.Name == "xl/workbook.xml":
				xl.workbook = newWorkbook(f, xl)
			case f.Name == "xl/_rels/workbook.xml.rels":
				xl.relationships = ooxml.NewRelationships(f, xl.pkg)
			case f.Name == "xl/sharedStrings.xml":
				xl.sharedStrings = newSharedStrings(f, xl)
			case f.Name == "xl/styles.xml":
				xl.styleSheet = newStyleSheet(f, xl)
			}
		}
	}

	//we need populated 'relationships' to resolve index for sheet
	reSheet := regexp.MustCompile(`xl/worksheets/[[:alpha:]]+[\d]+\.xml`)
	for _, file := range files {
		if f, ok := file.(*zip.File); ok {
			if reSheet.MatchString(f.Name) {
				newSheet(f, xl)
			}
		}
	}
}

//createSpreadsheet initialize a new XLSX document
func (xl *Spreadsheet) createSpreadsheet() {
	xl.relationships = ooxml.NewRelationships("xl/_rels/workbook.xml.rels", xl.pkg)
	xl.workbook = newWorkbook("xl/workbook.xml", xl)
	xl.sharedStrings = newSharedStrings("xl/sharedStrings.xml", xl)
	xl.styleSheet = newStyleSheet("xl/styles.xml", xl)
}
