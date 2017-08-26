package xlsx

import (
	"archive/zip"
	"fmt"
	"github.com/plandem/ooxml"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"math"
	"reflect"
)

//max length of excel's sheet name
const sheetNameLimit = 31

//Sheet is a higher level object that wraps ml.Worksheet with functionality
type Sheet struct {
	ml            ml.Worksheet
	workbook      *Workbook
	mergedRanges  []*rangeInfo
	isInitialized bool
	index         int
	file          *shared.PackageFile
}

//newSheet creates a new sheet and link it with workbook
func newSheet(f interface{}, doc *Spreadsheet) *Sheet {
	index := -1

	//not initialized yet?
	if doc.sheets == nil {
		doc.sheets = make([]*Sheet, len(doc.workbook.ml.Sheets))
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
	var sheet *Sheet
	if index >= 0 {
		sheet = &Sheet{
			index:    index,
			workbook: doc.workbook,
		}

		//link worksheet
		if index >= len(doc.sheets) {
			doc.sheets = append(doc.sheets, sheet)
		} else {
			doc.sheets[index] = sheet
		}

		sheet.file = shared.NewPackageFile(doc.pkg, f, &sheet.ml, sheet)
	}

	return sheet
}

//Name returns name of sheet
func (s *Sheet) Name() string {
	return s.workbook.ml.Sheets[s.index].Name
}

//SetName sets a name for sheet
func (s *Sheet) SetName(name string) {
	s.workbook.ml.Sheets[s.index].Name = shared.UniqueName(name, s.workbook.doc.GetSheetNames(), sheetNameLimit)
	s.workbook.file.MarkAsUpdated()
}

//SetState sets a visibility state for sheet
func (s *Sheet) SetState(state types.VisibilityType) {
	s.workbook.ml.Sheets[s.index].State = state
	s.workbook.file.MarkAsUpdated()
}

//Cell returns a cell for 0-based indexes
func (s *Sheet) Cell(colIndex, rowIndex int) *Cell {
	s.expandIfRequired(colIndex, rowIndex)
	s.resolveMergedIfRequired(false)

	//is merged cell?
	for _, mergedRange := range s.mergedRanges {
		if mergedRange.Contains(colIndex, rowIndex) {
			colIndex, rowIndex = mergedRange.fromCol, mergedRange.fromRow
			break
		}
	}

	data := s.ml.SheetData[rowIndex].Cells[colIndex]

	//if there is no any data for this cell, then create it
	if data == nil {
		data = &ml.Cell{
			Ref: types.CellRefFromIndexes(colIndex, rowIndex),
		}

		s.ml.SheetData[rowIndex].Cells[colIndex] = data
	}

	return &Cell{ml: data, sheet: s}
}

//CellByRef returns a cell for ref
func (s *Sheet) CellByRef(cellRef types.CellRef) *Cell {
	cid, rid := cellRef.ToIndexes()
	return s.Cell(cid, rid)
}

//Row returns a row for 0-based index
func (s *Sheet) Row(index int) *Row {
	s.expandIfRequired(0, index)

	data := s.ml.SheetData[index]
	return &Row{ml: data, sheet: s}
}

//refreshRefs update refs for all rows/cells starting from row with 0-based index
func (s *Sheet) refreshAllRefs(index int) {
	for iRow, rowMax := index, len(s.ml.SheetData); iRow < rowMax; iRow++ {
		row := s.ml.SheetData[iRow]
		row.Ref = iRow + 1

		for iCol, cell := range row.Cells {
			if !s.isCellEmpty(cell) {
				cell.Ref = types.CellRefFromIndexes(iCol, int(row.Ref-1))
			}
		}
	}
}

//refreshColRefs update refs only for cells at row 0-based index rowIndex and starting 0-based index colIndex
func (s *Sheet) refreshColRefs(colIndex, rowIndex int) {
	for iCol, colMax := colIndex, len(s.ml.SheetData[rowIndex].Cells); iCol < colMax; iCol++ {
		cell := s.ml.SheetData[rowIndex].Cells[iCol]
		if !s.isCellEmpty(cell) {
			cell.Ref = types.CellRefFromIndexes(iCol, rowIndex)
		}
	}
}

//InsertRow inserts a row at 0-based index and returns it. Using to insert a row between other rows.
func (s *Sheet) InsertRow(index int) *Row {
	//getting a current height
	_, currRef := s.ml.Dimension.Ref.ToCellRefs()
	_, curHeight := currRef.ToIndexes()

	//expand to a new height
	s.expandIfRequired(0, curHeight+1)

	//copy previous info
	copy(s.ml.SheetData[index+1:], s.ml.SheetData[index:])

	//clear previous info at this index
	for iCol := range s.ml.SheetData[index].Cells {
		s.ml.SheetData[index].Cells[iCol] = nil
	}

	//refresh refs
	s.refreshAllRefs(index)

	return s.Row(index)
}

//DeleteRow deletes a row at 0-based index
func (s *Sheet) DeleteRow(index int) {
	s.expandIfRequired(0, index)

	s.ml.SheetData = append(s.ml.SheetData[:index], s.ml.SheetData[index+1:]...)

	//now we must updated refs
	s.refreshAllRefs(index)

	//update dimension for a new size
	_, currRef := s.ml.Dimension.Ref.ToCellRefs()
	curWidth, curHeight := currRef.ToIndexes()
	s.ml.Dimension.Ref = types.Ref(types.CellRefFromIndexes(curWidth, curHeight - 1))
}

//Col returns a col for 0-based index
func (s *Sheet) Col(index int) *Col {
	s.expandIfRequired(index, 0)

	if s.ml.Cols == nil {
		s.ml.Cols = &[]*ml.Col{}
	}

	var data *ml.Col

	//Cols has 1-based index, but we are using 0-based to unify all indexes at library
	index++

	for _, c := range *s.ml.Cols {
		if c.Min == c.Max && c.Min == index {
			data = c
			break
		}
	}

	if data == nil {
		data = &ml.Col{
			Min: index,
			Max: index,
		}

		*s.ml.Cols = append(*s.ml.Cols, data)
	}

	return &Col{ml: data, sheet: s, index: index - 1}
}

//InsertCol inserts a col at 0-based index and returns it. Using to insert a col between other cols.
func (s *Sheet) InsertCol(index int) *Col {
	//getting current width
	_, currRef := s.ml.Dimension.Ref.ToCellRefs()
	curWidth, _ := currRef.ToIndexes()

	//expand to a new width
	s.expandIfRequired(curWidth+1, 0)

	for iRow, row := range s.ml.SheetData {
		//copy previous info
		copy(s.ml.SheetData[iRow].Cells[index+1:], s.ml.SheetData[iRow].Cells[index:])

		//clear previous info at this index
		s.ml.SheetData[iRow].Cells[index] = nil

		//refresh refs
		s.refreshColRefs(index, row.Ref-1)
	}

	return s.Col(index)
}

//DeleteCol deletes a col at 0-based index
func (s *Sheet) DeleteCol(index int) {
	s.expandIfRequired(index, 0)

	if s.ml.Cols != nil {
		for _, c := range *s.ml.Cols {
			if c.Min == c.Max && c.Min == index {
				*s.ml.Cols = append((*s.ml.Cols)[:index], (*s.ml.Cols)[index+1:]...)
			}
		}
	}

	for iRow, row := range s.ml.SheetData {
		//delete col
		s.ml.SheetData[iRow].Cells = append(s.ml.SheetData[iRow].Cells[:index], s.ml.SheetData[iRow].Cells[index+1:]...)

		//refresh refs
		s.refreshColRefs(index, row.Ref-1)
	}

	//update dimension for a new size
	_, currRef := s.ml.Dimension.Ref.ToCellRefs()
	curWidth, curHeight := currRef.ToIndexes()
	s.ml.Dimension.Ref = types.Ref(types.CellRefFromIndexes(curWidth - 1, curHeight))
}

//Range returns a range for ref
func (s *Sheet) Range(ref types.Ref) *Range {
	return &Range{
		newRangeInfo(ref),
		s,
	}
}

//TotalCols returns total number of cols in grid
func (s *Sheet) TotalCols() int {
	_, currRef := s.ml.Dimension.Ref.ToCellRefs()
	curWidth, _ := currRef.ToIndexes()
	return curWidth + 1
}

//TotalRows returns total number of rows in grid
func (s *Sheet) TotalRows() int {
	_, currRef := s.ml.Dimension.Ref.ToCellRefs()
	_, curHeight := currRef.ToIndexes()
	return curHeight + 1
}

//resolveDimension check if there is a 'dimension' information(optional) and if there is no any, then calculate it from existing data
func (s *Sheet) resolveDimension(force bool) {
	if !force && (s.ml.Dimension != nil && s.ml.Dimension.Ref != "") {
		return
	}

	var (
		maxWidth float64
		minWidth float64 = math.MaxFloat64

		maxHeight float64
		minHeight float64 = math.MaxFloat64
	)

	//supposed that grid holds rows/cells with valid refs
	for _, row := range s.ml.SheetData {
		maxHeight = math.Max(maxHeight, float64(row.Ref))
		minHeight = math.Min(minHeight, float64(row.Ref))

		for _, cell := range row.Cells {
			colIndex, _ := types.CellRef(cell.Ref).ToIndexes()
			maxWidth = math.Max(maxWidth, float64(colIndex))
			minWidth = math.Min(minWidth, float64(colIndex))
		}
	}

	var dimension types.Ref
	fromRef := types.CellRefFromIndexes(int(math.Min(minWidth, 0)), int(math.Min(minHeight, 0)))
	toRef := types.CellRefFromIndexes(int(maxWidth), int(maxHeight))

	if fromRef == toRef {
		dimension = types.Ref(fromRef)
	} else {
		dimension = types.RefFromCellRefs(fromRef, toRef)
	}

	s.ml.Dimension = &ml.SheetDimension{Ref: dimension}
}

//expandOnInit expands grid to required dimension and copy existing data
func (s *Sheet) expandOnInit() {
	s.resolveDimension(false)

	//during initialize phase we need to do hard work first time - expand grid to required size and copy it with existing data
	_, nextRef := s.ml.Dimension.Ref.ToCellRefs()
	nextWidth, nextHeight := nextRef.ToIndexes()

	//convert indexes to size
	nextWidth++
	nextHeight++

	//expand grid
	grid := make([]*ml.Row, nextHeight)
	for iRow := 0; iRow < nextHeight; iRow++ {
		grid[iRow] = &ml.Row{
			Ref:   iRow + 1,
			Cells: make([]*ml.Cell, nextWidth),
		}
	}

	//fill grid with data
	for _, row := range s.ml.SheetData {
		iRow := int(row.Ref - 1)
		for _, cell := range row.Cells {
			//add cell info
			if !s.isCellEmpty(cell) {
				iCellCol, iCellRow := cell.Ref.ToIndexes()
				grid[iCellRow].Cells[iCellCol] = cell
			}
		}

		//add row info
		row := row
		row.Cells = grid[iRow].Cells
		grid[iRow] = row
	}

	s.ml.SheetData = grid
	s.ml.Dimension.Ref = types.Ref(types.CellRefFromIndexes(nextWidth-1, nextHeight-1))
	s.isInitialized = true
}

//expandIfRequired expands grid to required dimension
func (s *Sheet) expandIfRequired(colIndex, rowIndex int) {
	if !s.isInitialized {
		s.expandOnInit()
	}

	s.resolveDimension(false)

	//during expand phase we need to increase grid to new size only, without copying any info - it's already in place
	_, currRef := s.ml.Dimension.Ref.ToCellRefs()
	curWidth, curHeight := currRef.ToIndexes()
	nextWidth, nextHeight := colIndex, rowIndex

	//shrink is not supported here, so fix for current size if required
	if nextWidth < curWidth {
		nextWidth = curWidth
	}

	if nextHeight < curHeight {
		nextHeight = curHeight
	}

	//if size is fit to current, then ignore
	if curWidth >= nextWidth && curHeight >= nextHeight {
		return
	}

	//convert indexes to size
	curWidth++
	curHeight++
	nextWidth++
	nextHeight++

	//TODO: think about optimizing - use incremental step to decrease number of allocations for +1 step case, e.g.: size = (size * 3) / 2 + 1
	//step to expand width
	widthStep := nextWidth - curWidth
	if widthStep > 0 {
		for iRow, row := range s.ml.SheetData {
			s.ml.SheetData[iRow].Cells = append(row.Cells, make([]*ml.Cell, widthStep)...)
		}
	}

	//step to expand height
	heightStep := nextHeight - curHeight
	if heightStep > 0 {
		s.ml.SheetData = append(s.ml.SheetData, make([]*ml.Row, heightStep)...)
		for iRow := curHeight; iRow < nextHeight; iRow++ {
			s.ml.SheetData[iRow] = &ml.Row{
				Ref:   iRow + 1,
				Cells: make([]*ml.Cell, nextWidth),
			}
		}
	}

	//update dimension for a new size
	s.ml.Dimension.Ref = types.Ref(types.CellRefFromIndexes(nextWidth-1, nextHeight-1))
}

//shrinkIfRequired shrinks grid to minimal size and set actual dimension. Called right before packing sheet data.
func (s *Sheet) shrinkIfRequired() {
	grid := make([]*ml.Row, 0, len(s.ml.SheetData))

	for _, row := range s.ml.SheetData {
		nextRow := &ml.Row{}
		*nextRow = *row
		nextRow.Cells = make([]*ml.Cell, 0, len(row.Cells))

		for iCol, cell := range row.Cells {
			if !s.isCellEmpty(cell) {
				cell.Ref = types.CellRefFromIndexes(iCol, int(row.Ref-1))
				nextRow.Cells = append(nextRow.Cells, cell)
			}
		}

		if !s.isRowEmpty(nextRow) {
			grid = append(grid, nextRow)
		}
	}

	s.ml.SheetData = grid
	s.resolveDimension(true)
}

//isCellEmpty checks if cell is empty - has no value and any formatting
func (s *Sheet) isCellEmpty(c *ml.Cell) bool {
	if c != nil && (*c != ml.Cell{Ref: c.Ref}) {
		return false
	}

	return true
}

//isRowEmpty checks if row is empty (supposed that only non empty cells here) - has no cells
func (s *Sheet) isRowEmpty(r *ml.Row) bool {
	return r == nil || (len(r.Cells) == 0 && reflect.DeepEqual(r, &ml.Row{}))
}

//resolveMergedIfRequired transforms merged cells into rangeInfo
func (s *Sheet) resolveMergedIfRequired(force bool) {
	if force || (s.ml.MergeCells != nil && (len(*s.ml.MergeCells) != len(s.mergedRanges))) {
		s.mergedRanges = make([]*rangeInfo, len(*s.ml.MergeCells))

		for i, mergedRef := range *s.ml.MergeCells {
			s.mergedRanges[i] = newRangeInfo(mergedRef.Ref)
		}
	}
}

//BeforeMarshalXML shrinks data to optimize output and returns related ML information for marshaling
func (s *Sheet) BeforeMarshalXML() interface{} {
	s.shrinkIfRequired()
	s.isInitialized = false
	return &s.ml
}
