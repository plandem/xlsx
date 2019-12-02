// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"archive/zip"
	"fmt"
	"github.com/plandem/ooxml"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/format/conditional"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"github.com/plandem/xlsx/types/options/sheet"
	"math"
	"path/filepath"
	"reflect"
)

type sheetInfo struct {
	ml            ml.Worksheet
	workbook      *workbook
	file          *ooxml.PackageFile
	columns       *columns
	mergedCells   *mergedCells
	hyperlinks    *hyperlinks
	conditionals  *conditionals
	comments      *comments
	drawingsVML   *drawingsVML
	filters       *filters
	relationships *ooxml.Relationships
	sheet         Sheet
	sheetMode     SheetMode
	isInitialized bool
	index         int
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
	return r == nil ||
		reflect.DeepEqual(r, &ml.Row{Ref: r.Ref, Cells: []*ml.Cell{}}) ||
		reflect.DeepEqual(r, &ml.Row{Ref: r.Ref}) ||
		reflect.DeepEqual(r, &ml.Row{})
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
		sheet.columns = newColumns(sheet)
		sheet.mergedCells = newMergedCells(sheet)
		sheet.hyperlinks = newHyperlinks(sheet)
		sheet.conditionals = newConditionals(sheet)
		sheet.filters = newFilters(sheet)
		sheet.comments = newComments(sheet)
		sheet.drawingsVML = newDrawingsVML(sheet)
	}

	return sheet
}

//some private methods used objects that use Sheet implementation and have no access to internal information
func (s *sheetInfo) mode() SheetMode {
	return s.sheetMode
}

func (s *sheetInfo) info() *sheetInfo {
	return s
}

//Name returns name of sheet
func (s *sheetInfo) Name() string {
	return s.workbook.ml.Sheets[s.index].Name
}

//SetName sets a name for sheet
func (s *sheetInfo) SetName(name string) {
	s.workbook.ml.Sheets[s.index].Name = ooxml.UniqueName(name, s.workbook.doc.SheetNames(), internal.ExcelSheetNameLimit)
	s.workbook.file.MarkAsUpdated()
}

//SetOptions sets options for sheet
func (s *sheetInfo) SetOptions(o *options.Info) {
	s.workbook.ml.Sheets[s.index].State = o.Visibility
	s.workbook.file.MarkAsUpdated()
}

//SetActive sets the sheet as active
func (s *sheetInfo) SetActive() {
	//set activate from workbook side
	if len(s.workbook.ml.BookViews.Items) == 0 {
		s.workbook.ml.BookViews.Items = append(s.workbook.ml.BookViews.Items, &ml.BookView{
			ActiveTab: s.index,
		})
	} else {
		s.workbook.ml.BookViews.Items[0].ActiveTab = s.index
	}

	//set active from worksheet side
	if len(s.ml.SheetViews.Items) > 0 {
		s.ml.SheetViews.Items[0].TabSelected = true
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

//CellByRef returns a cell for ref
func (s *sheetInfo) CellByRef(cellRef types.CellRef) *Cell {
	cid, rid := cellRef.ToIndexes()
	return s.sheet.Cell(cid, rid)
}

//Range returns a range for indexes
func (s *sheetInfo) Range(fromCol, fromRow, toCol, toRow int) *Range {
	return newRange(s.sheet, fromCol, toCol, fromRow, toRow)
}

//RangeByRef returns a range for ref
func (s *sheetInfo) RangeByRef(ref types.Ref) *Range {
	return newRangeFromRef(s.sheet, ref)
}

//MergeRows merges rows between fromIndex and toIndex
func (s *sheetInfo) MergeRows(fromIndex, toIndex int) error {
	return s.Range(0, fromIndex, internal.ExcelColumnLimit, toIndex).Merge()
}

//MergeCols merges cols between fromIndex and toIndex
func (s *sheetInfo) MergeCols(fromIndex, toIndex int) error {
	return s.Range(fromIndex, 0, toIndex, internal.ExcelRowLimit).Merge()
}

//SplitRows splits rows between fromIndex and toIndex
func (s *sheetInfo) SplitRows(fromIndex, toIndex int) {
	s.Range(0, fromIndex, internal.ExcelColumnLimit, toIndex).Split()
}

//SplitCols splits cols between fromIndex and toIndex
func (s *sheetInfo) SplitCols(fromIndex, toIndex int) {
	s.Range(fromIndex, 0, toIndex, internal.ExcelRowLimit).Split()
}

//AddConditional adds a new conditional formatting with additional refs if required
func (s *sheetInfo) AddConditional(conditional *conditional.Info, refs ...types.Ref) error {
	return s.conditionals.Add(conditional, refs)
}

//DeleteConditional deletes a conditional formatting from refs
func (s *sheetInfo) DeleteConditional(refs ...types.Ref) {
	s.conditionals.Remove(refs)
}

//AutoFilter adds auto filter in provided Ref range
func (s *sheetInfo) AutoFilter(ref types.Ref, settings ...interface{}) {
	s.filters.AddAuto(ref, settings)
}

//AddFilter adds a filter to column with index
func (s *sheetInfo) AddFilter(colIndex int, settings ...interface{}) error {
	return s.filters.Add(colIndex, settings)
}

//DeleteFilter deletes a filter from column with index
func (s *sheetInfo) DeleteFilter(colIndex int) {
	s.filters.Remove(colIndex)
}

//DefineName creates a name for sheet scope and provided value
func (s *sheetInfo) DefineName(name, value string) error {
	return s.workbook.doc.definedNames.Add(name, value, s.index)
}

//Close frees allocated by sheet resources
func (s *sheetInfo) Close() {

}

//afterOpen is callback that will be called right after requesting an already existing sheet. By default, it does nothing
func (s *sheetInfo) afterOpen() {
}

func (s *sheetInfo) attachRelationshipsIfRequired() {
	if s.relationships == nil {
		fileName := s.workbook.doc.relationships.GetTargetById(string(s.workbook.ml.Sheets[s.index].RID))
		fileName = fmt.Sprintf("xl/worksheets/_rels/%s.rels", filepath.Base(fileName))

		if file := s.workbook.doc.pkg.File(fileName); file != nil {
			s.relationships = ooxml.NewRelationships(file, s.workbook.doc.pkg)
		} else {
			s.relationships = ooxml.NewRelationships(fileName, s.workbook.doc.pkg)
		}
	}
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

	return &s.ml
}

func (s *sheetInfo) AfterMarshalXML(content []byte) []byte {
	if fix, ok := s.sheet.(ooxml.MarshalFixation); ok {
		return fix.AfterMarshalXML(content)
	}

	return content
}

func (s *sheetInfo) resolveStyleID(st interface{}) styles.DirectStyleID {
	if st == nil {
		return 0
	}

	if styleID, ok := st.(styles.DirectStyleID); ok {
		return styleID
	}

	//we can update styleSheet only when sheet is in write mode, to prevent pollution of styleSheet with fake values
	if (s.mode() & sheetModeWrite) == 0 {
		panic(errorNotSupportedWrite)
	}

	var format *styles.Info
	if f, ok := st.(styles.Info); ok {
		format = &f
	} else if f, ok := st.(*styles.Info); ok {
		format = f
	} else {
		panic("only DirectStyleID or styles.Info supported as styles for cell")
	}

	return s.workbook.doc.styleSheet.addStyle(format)
}
