package xlsx

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"math"
)

type sheetReadWrite struct {
	*sheetInfo
}

var _ Sheet = (*sheetReadWrite)(nil)

func (s *sheetReadWrite) setDimension(cols, rows int, resize bool) {
	if cols <= 0 {
		cols = 1
	}

	if rows <= 0 {
		rows = 1
	}

	//converts rows/cols into indexes
	cols--
	rows--

	if resize {
		s.expandIfRequired(cols, rows)
	}

	s.ml.Dimension = &ml.SheetDimension{Bounds: types.BoundsFromIndexes(0, 0, cols, rows)}
}

//SetDimension sets total number of cols and rows in sheet
func (s *sheetReadWrite) SetDimension(cols, rows int) {
	s.setDimension(cols, rows, true)
}

//Cell returns a cell for 0-based indexes
func (s *sheetReadWrite) Cell(colIndex, rowIndex int) *Cell {
	s.expandIfRequired(colIndex, rowIndex)

	colIndex, rowIndex = s.mergedCells.Resolve(colIndex, rowIndex)
	data := s.ml.SheetData[rowIndex].Cells[colIndex]

	//if there is no any data for this cell, then create it
	if data == nil {
		data = &ml.Cell{
			Ref: types.CellRefFromIndexes(colIndex, rowIndex),
		}

		s.ml.SheetData[rowIndex].Cells[colIndex] = data
	}

	return &Cell{ml: data, sheet: s.sheetInfo}
}

//CellByRef returns a cell for ref
func (s *sheetReadWrite) CellByRef(cellRef types.CellRef) *Cell {
	cid, rid := cellRef.ToIndexes()
	return s.Cell(cid, rid)
}

//Row returns a row for 0-based index
func (s *sheetReadWrite) Row(index int) *Row {
	s.expandIfRequired(0, index)

	data := s.ml.SheetData[index]
	return &Row{
		data,
		newRange(s, 0, len(data.Cells)-1, index, index),
	}
}

//refreshRefs update refs for all rows/cells starting from row with 0-based index
func (s *sheetReadWrite) refreshAllRefs(index int) {
	for iRow, rowMax := index, len(s.ml.SheetData); iRow < rowMax; iRow++ {
		row := s.ml.SheetData[iRow]
		row.Ref = iRow + 1

		for iCol, cell := range row.Cells {
			if !isCellEmpty(cell) {
				cell.Ref = types.CellRefFromIndexes(iCol, int(row.Ref-1))
			}
		}
	}
}

//refreshColRefs update refs only for cells at row 0-based index rowIndex and starting 0-based index colIndex
func (s *sheetReadWrite) refreshColRefs(colIndex, rowIndex int) {
	for iCol, colMax := colIndex, len(s.ml.SheetData[rowIndex].Cells); iCol < colMax; iCol++ {
		cell := s.ml.SheetData[rowIndex].Cells[iCol]
		if !isCellEmpty(cell) {
			cell.Ref = types.CellRefFromIndexes(iCol, rowIndex)
		}
	}
}

//InsertRow inserts a row at 0-based index and returns it. Using to insert a row between other rows.
func (s *sheetReadWrite) InsertRow(index int) *Row {
	//getting current height
	_, rows := s.Dimension()

	//expand to a new height
	s.expandIfRequired(0, rows)

	//copy previous info
	copy(s.ml.SheetData[index+1:], s.ml.SheetData[index:])

	//clear previous info at this index
	s.ml.SheetData[index] = &ml.Row{Cells: make([]*ml.Cell, len(s.ml.SheetData[index].Cells))}

	//refresh refs
	s.refreshAllRefs(index)

	return s.Row(index)
}

//DeleteRow deletes a row at 0-based index
func (s *sheetReadWrite) DeleteRow(index int) {
	s.expandIfRequired(0, index)

	s.ml.SheetData = append(s.ml.SheetData[:index], s.ml.SheetData[index+1:]...)

	//now we must updated refs
	s.refreshAllRefs(index)

	//update dimension for a new size
	cols, rows := s.Dimension()
	s.setDimension(cols, rows-1, false)
}

//Col returns a col for 0-based index
func (s *sheetReadWrite) Col(index int) *Col {
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

	_, rows := s.Dimension()

	index--
	return &Col{
		data,
		newRange(s, index, index, 0, rows-1),
	}
}

//InsertCol inserts a col at 0-based index and returns it. Using to insert a col between other cols.
func (s *sheetReadWrite) InsertCol(index int) *Col {
	//getting current width
	cols, _ := s.Dimension()

	//expand to a new width
	s.expandIfRequired(cols, 0)

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
func (s *sheetReadWrite) DeleteCol(index int) {
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
	cols, rows := s.Dimension()
	s.setDimension(cols-1, rows, false)
}

//Range returns a range for ref
func (s *sheetReadWrite) Range(ref types.Ref) *Range {
	return newRangeFromRef(s, ref)
}

//Cols returns iterator for all cols of sheet
func (s *sheetReadWrite) Cols() ColIterator {
	cols, rows := s.Dimension()
	s.expandIfRequired(cols-1, rows-1)
	return newColIterator(s)
}

//Rows returns iterator for all rows of sheet
func (s *sheetReadWrite) Rows() RowIterator {
	cols, rows := s.Dimension()
	s.expandIfRequired(cols-1, rows-1)
	return newRowIterator(s)
}

//resolveDimension check if there is a 'dimension' information(optional) and if there is no any, then calculate it from existing data
func (s *sheetReadWrite) resolveDimension(force bool) {
	if !force && (s.ml.Dimension != nil && !s.ml.Dimension.Bounds.IsEmpty()) {
		return
	}

	var (
		maxWidth float64
		minWidth = math.MaxFloat64

		maxHeight float64
		minHeight = math.MaxFloat64
	)

	//supposed that grid holds rows/cells with valid refs
	for _, row := range s.ml.SheetData {
		maxHeight = math.Max(maxHeight, float64(row.Ref)-1)
		minHeight = math.Min(minHeight, float64(row.Ref)-1)

		for _, cell := range row.Cells {
			colIndex, _ := types.CellRef(cell.Ref).ToIndexes()
			maxWidth = math.Max(maxWidth, float64(colIndex))
			minWidth = math.Min(minWidth, float64(colIndex))
		}
	}

	s.ml.Dimension = &ml.SheetDimension{Bounds: types.BoundsFromIndexes(
		int(math.Min(minWidth, 0)), int(math.Min(minHeight, 0)),
		int(maxWidth), int(maxHeight),
	)}
}

//expandOnInit expands grid to required dimension and copy existing data
func (s *sheetReadWrite) expandOnInit() {
	s.resolveDimension(false)

	//during initialize phase we need to do hard work first time - expand grid to required size and copy it with existing data
	nextWidth, nextHeight := s.Dimension()

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
			if !isCellEmpty(cell) {
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
	s.isInitialized = true
	s.setDimension(nextWidth, nextHeight, false)
}

//expandIfRequired expands grid to required dimension
func (s *sheetReadWrite) expandIfRequired(colIndex, rowIndex int) {
	if !s.isInitialized {
		s.expandOnInit()
	}

	s.resolveDimension(false)

	//during expand phase we need to increase grid to new size only, without copying any info - it's already in place
	curWidth, curHeight := s.Dimension()
	nextWidth, nextHeight := colIndex+1, rowIndex+1

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
	s.setDimension(nextWidth, nextHeight, false)
}

//shrinkIfRequired shrinks grid to minimal size and set actual dimension. Called right before packing sheet data.
func (s *sheetReadWrite) shrinkIfRequired() {
	grid := make([]*ml.Row, 0, len(s.ml.SheetData))

	for _, row := range s.ml.SheetData {
		nextRow := &ml.Row{}
		*nextRow = *row
		nextRow.Cells = make([]*ml.Cell, 0, len(row.Cells))

		for iCol, cell := range row.Cells {
			if !isCellEmpty(cell) {
				cell.Ref = types.CellRefFromIndexes(iCol, int(row.Ref-1))
				nextRow.Cells = append(nextRow.Cells, cell)
			}
		}

		if !isRowEmpty(nextRow) {
			grid = append(grid, nextRow)
		}
	}

	s.ml.SheetData = grid
	s.resolveDimension(true)
}

//BeforeMarshalXML shrinks data to optimize output and returns related ML information for marshaling
func (s *sheetReadWrite) BeforeMarshalXML() interface{} {
	s.shrinkIfRequired()
	s.isInitialized = false
	return &s.ml
}

//afterOpen is callback that will be called right after requesting an already existing sheet. By default, it does nothing
func (s *sheetReadWrite) afterOpen() {
	//make a grid
	s.file.LoadIfRequired(s.expandOnInit)

	//adds a styles for types
	s.workbook.doc.styleSheet.addTypedStylesIfRequired()

	//mark file as updated
	s.file.MarkAsUpdated()
}

//afterCreate initializes a new sheet
func (s *sheetReadWrite) afterCreate(name string) {
	//register file
	s.sheetInfo.afterCreate(name)

	//make a grid
	s.expandOnInit()

	//adds a styles for types
	s.workbook.doc.styleSheet.addTypedStylesIfRequired()
}
