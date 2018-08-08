package xlsx

import (
	"archive/zip"
	"fmt"
	"github.com/plandem/ooxml"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/options"
	"math"
	"reflect"
)

type sheetInfo struct {
	ml            ml.Worksheet
	workbook      *Workbook
	isInitialized bool
	index         int
	file          *ooxml.PackageFile
	mergedCells   *mergedCellManager
	sheet         Sheet
	sheetMode     sheetMode
}

//isCellEmpty checks if cell is empty - has no value and any formatting
func isCellEmpty(c *ml.Cell) bool {
	if c != nil && (*c != ml.Cell{Ref: c.Ref}) {
		return false
	}

	return true
}

//isRowEmpty checks if row is empty (supposed that only non empty cells here) - has no cells
func isRowEmpty(r *ml.Row) bool {
	return r == nil || (len(r.Cells) == 0 && reflect.DeepEqual(r, &ml.Row{}))
}

//newSheetInfo creates a new sheetInfo and link it with workbook
func newSheetInfo(f interface{}, doc *Spreadsheet) *sheetInfo {
	index := -1

	//not initialized yet?
	if doc.sheets == nil {
		doc.sheets = make([]*sheetInfo, len(doc.workbook.ml.Sheets))
	}

	//is it existing sheet?
	if zf, ok := f.(*zip.File); ok && zf != nil {
		//get RID for an existing sheet
		if rid := doc.relationships.GetIdByTarget(zf.Name); rid != "" {
			for idx, sheet := range doc.workbook.ml.Sheets {
				if sheet.RID == sharedML.RID(rid) {
					index = idx
					break
				}
			}
		}
	} else if fileName, ok := f.(string); ok && len(fileName) > 3 {
		_, rid := doc.relationships.AddFile(internal.RelationTypeWorksheet, fileName)

		//get the next SheetID for a new sheet
		var sheetID uint
		for _, sheet := range doc.workbook.ml.Sheets {
			sheetID = uint(math.Max(float64(sheetID), float64(sheet.SheetID)))
		}

		//get index for a new sheet
		index = len(doc.workbook.ml.Sheets)

		//insert ml.Sheet
		sheetID++
		doc.workbook.ml.Sheets = append(doc.workbook.ml.Sheets, &ml.Sheet{
			RID:     rid,
			SheetID: sheetID,
			Name:    fmt.Sprintf("Sheet%d", sheetID), //temporary name for sheet
		})
	}

	//Link sheet with workbook
	var sheet *sheetInfo
	if index >= 0 {
		sheet = &sheetInfo{
			index:    index,
			workbook: doc.workbook,
		}

		//link worksheet
		if index >= len(doc.sheets) {
			doc.sheets = append(doc.sheets, sheet)
		} else {
			doc.sheets[index] = sheet
		}

		sheet.file = ooxml.NewPackageFile(doc.pkg, f, &sheet.ml, sheet)
		sheet.mergedCells = newMergedCellManager(sheet)
	}

	return sheet
}

func (s *sheetInfo) mode() sheetMode {
	return s.sheetMode
}

//Name returns name of sheet
func (s *sheetInfo) Name() string {
	return s.workbook.ml.Sheets[s.index].Name
}

//SetName sets a name for sheet
func (s *sheetInfo) SetName(name string) {
	s.workbook.ml.Sheets[s.index].Name = ooxml.UniqueName(name, s.workbook.doc.GetSheetNames(), sheetNameLimit)
	s.workbook.file.MarkAsUpdated()
}

//Set sets options for sheet
func (s *sheetInfo) Set(o *options.SheetOptions) {
	if o.Visibility >= options.VisibilityTypeVisible && o.Visibility <= options.VisibilityTypeVeryHidden {
		s.workbook.ml.Sheets[s.index].State = o.Visibility
		s.workbook.file.MarkAsUpdated()
	}
}

//SetActive sets the sheet as active
func (s *sheetInfo) SetActive() {
	//set activate from workbook side
	if s.workbook.ml.BookViews == nil || len(*s.workbook.ml.BookViews) == 0 {
		s.workbook.ml.BookViews = &[]*ml.BookView{{
			ActiveTab: s.index,
		}}
	} else {
		(*s.workbook.ml.BookViews)[0].ActiveTab = s.index
	}

	//set active from worksheet side
	if s.ml.SheetViews != nil && len(s.ml.SheetViews.SheetView) > 0 {
		s.ml.SheetViews.SheetView[0].TabSelected = true
	}

	s.workbook.file.MarkAsUpdated()
}

//Dimension returns total number of cols and rows in sheet
func (s *sheetInfo) Dimension() (cols int, rows int) {
	if s.ml.Dimension == nil || s.ml.Dimension.Bounds.IsEmpty() {
		return 0, 0
	}

	//we can't use dimension of bounds, because it depends on fromCol, fromRow, but in case of sheet we need maximum dimension to fit content
	cols, rows = s.ml.Dimension.Bounds.ToCol+1, s.ml.Dimension.Bounds.ToRow+1
	return
}

//Close frees allocated by sheet resources
func (s *sheetInfo) Close() {

}

//afterOpen is callback that will be called right after requesting an already existing sheet. By default, it does nothing
func (s *sheetInfo) afterOpen() {
}

//afterCreate is callback that will be called right after creating a new sheet. By default, it registers sheet at spreadsheet
func (s *sheetInfo) afterCreate(name string) {
	if len(name) > 0 {
		s.SetName(name)
	}

	s.file.MarkAsUpdated()
	s.workbook.file.MarkAsUpdated()
	s.workbook.doc.pkg.ContentTypes().RegisterContent(s.file.FileName(), internal.ContentTypeWorksheet)
}

func (s *sheetInfo) BeforeMarshalXML() interface{} {
	if prep, ok := s.sheet.(ooxml.MarshalPreparation); ok {
		return prep.BeforeMarshalXML()
	}

	return s.ml
}

func (s *sheetInfo) AfterMarshalXML(content []byte) []byte {
	if fix, ok := s.sheet.(ooxml.MarshalFixation); ok {
		return fix.AfterMarshalXML(content)
	}

	return content
}
