package xlsx

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/format"
	"regexp"
)

//Spreadsheet is a higher level object that wraps OOXML package with XLSX functionality
type Spreadsheet struct {
	ooxml.Package
	pkg           *ooxml.PackageInfo
	workbook      *Workbook
	sheets        []*sheetInfo
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
func (xl *Spreadsheet) Sheet(i int) Sheet {
	if i >= len(xl.sheets) {
		return nil
	}

	si := xl.sheets[i]
	si.sheetMode = sheetModeRead | sheetModeWrite

	sheet := &sheetReadWrite{si}
	si.sheet = sheet
	sheet.afterOpen()
	return sheet
}

//SheetReader returns a sheet by 0-based index that opened in stream reading mode
//In stream reading mode only forward reading is allowed and no updates will be applied
//For multi phase mode sheet will be iterated two times: first one to load meta information (e.g. merged cells) and another one for sheet data
func (xl *Spreadsheet) SheetReader(i int, multiPhase bool) Sheet {
	if i >= len(xl.sheets) {
		return nil
	}

	si := xl.sheets[i]
	si.sheetMode = sheetModeRead | sheetModeStream
	sheet := &sheetReadStream{sheetInfo: &(*si), multiPhase: multiPhase}
	sheet.afterOpen()
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
func (xl *Spreadsheet) AddSheet(name string) Sheet {
	if si := newSheetInfo(fmt.Sprintf("xl/worksheets/sheet%d.xml", len(xl.workbook.ml.Sheets)+1), xl); si != nil {
		sheet := &sheetReadWrite{si}
		si.sheet = sheet
		si.sheetMode = sheetModeRead | sheetModeWrite
		sheet.afterCreate(name)
		return sheet
	}

	return nil
}

//AddFormatting adds a new style formatting to document and return related ID that can be used lately
func (xl *Spreadsheet) AddFormatting(style *format.StyleFormat) format.StyleID {
	return xl.styleSheet.addStyle(style)
}

//IsValid validates document and return error if there is any error. Using right before saving.
func (xl *Spreadsheet) IsValid() error {
	if len(xl.sheets) == 0 {
		return errors.New("spreadsheet requires at least one worksheet")
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
				newSheetInfo(f, xl)
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
