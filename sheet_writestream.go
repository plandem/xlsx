package xlsx

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/format/conditional"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"github.com/plandem/xlsx/types/options/sheet"
)

type sheetWriteStream struct {
	*sheetInfo
	stream     *ooxml.StreamFileWriter
	worksheet  xml.StartElement
	sheetData  xml.StartElement
	currentRow *ml.Row
}

var _ Sheet = (*sheetWriteStream)(nil)

//expandIfRequired expands grid to required dimension
func (s *sheetWriteStream) expandIfRequired(colIndex, rowIndex int) {
	curWidth, curHeight := s.Dimension()
	nextWidth, nextHeight := colIndex+1, rowIndex+1

	if nextWidth < curWidth {
		nextWidth = curWidth
	}

	if nextHeight < curHeight {
		nextHeight = curHeight
	}

	//expand current row, if required
	if s.currentRow != nil && len(s.currentRow.Cells) < nextWidth {
		s.currentRow.Cells = append(s.currentRow.Cells, make([]*ml.Cell, nextWidth-curWidth)...)
	}

	//if size is fit to current, then ignore
	if curWidth >= nextWidth && curHeight >= nextHeight {
		return
	}

	s.ml.Dimension = &ml.SheetDimension{Bounds: types.BoundsFromIndexes(0, 0, nextWidth-1, nextHeight-1)}
}

func (s *sheetWriteStream) Cell(colIndex, rowIndex int) *Cell {
	s.expandIfRequired(colIndex, rowIndex)

	var data *ml.Cell
	row := s.Row(rowIndex)
	data = row.ml.Cells[colIndex]

	//if there is no any data for this cell, then create it
	if data == nil {
		data = &ml.Cell{
			Ref: types.CellRefFromIndexes(colIndex, rowIndex),
		}

		row.ml.Cells[colIndex] = data
	}

	return &Cell{ml: data, sheet: s.sheetInfo}
}

func (s *sheetWriteStream) emptyDataRow(indexRef int) *ml.Row {
	width, _ := s.Dimension()
	return &ml.Row{
		Ref:   indexRef,
		Cells: make([]*ml.Cell, width),
	}
}

func (s *sheetWriteStream) saveCurrentRow() {
	hasCells := false

	for i, c := range s.currentRow.Cells {
		if isCellEmpty(c) {
			s.currentRow.Cells[i] = nil
		} else {
			hasCells = true
		}
	}

	if !hasCells {
		s.currentRow.Cells = nil
	}

	if !isRowEmpty(s.currentRow) {
		if err := s.stream.EncodeElement(s.currentRow, xml.StartElement{Name: xml.Name{Local: "row"}}); err != nil {
			panic(err)
		}
	}
}

func (s *sheetWriteStream) Row(index int) *Row {
	s.expandIfRequired(0, index)
	indexRef := index + 1

	var data *ml.Row

	if s.currentRow == nil {
		data = s.emptyDataRow(indexRef)
		s.currentRow = data
	} else {
		if indexRef < s.currentRow.Ref {
			//return empty row, if request previous row
			data = s.emptyDataRow(indexRef)
		} else if indexRef == s.currentRow.Ref {
			//return current row, if same index
			data = s.currentRow
		} else {
			s.saveCurrentRow()

			//create a new row
			data = s.emptyDataRow(indexRef)
			s.currentRow = data
		}
	}

	return &Row{
		data,
		newRange(s, 0, len(data.Cells)-1, index, index),
	}
}

//afterCreate initializes a new sheet
func (s *sheetWriteStream) afterCreate(name string) {
	//register file
	s.sheetInfo.afterCreate(name)

	//open worksheet
	s.worksheet = xml.StartElement{Name: xml.Name{
		Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main",
		Local: "worksheet",
	}}

	var err error

	if s.stream, err = s.file.WriteStream(true); err == nil {
		if err = s.stream.EncodeToken(s.worksheet); err == nil {
			s.sheetData = xml.StartElement{Name: xml.Name{
				Local: "sheetData",
			}}

			err = s.stream.EncodeToken(s.sheetData)
		}
	}

	if err != nil {
		panic(err)
	}
}

//Close frees allocated by sheet resources
func (s *sheetWriteStream) Close() {
	s.saveCurrentRow()

	//close sheetData
	if err := s.stream.EncodeToken(s.sheetData.End()); err != nil {
		panic(err)
	}

	//close worksheet
	if err := s.stream.EncodeToken(s.worksheet.End()); err != nil {
		panic(err)
	}

	if err := s.stream.Close(); err != nil {
		panic(err)
	}
}

func (s *sheetWriteStream) Rows() RowIterator {
	panic(errorNotSupported)
}

//not allowed methods for stream reading mode
func (s *sheetWriteStream) Col(index int) *Col {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) Cols() ColIterator {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) InsertCol(index int) *Col {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) InsertRow(index int) *Row {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) DeleteRow(index int) {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) DeleteCol(index int) {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) SetDimension(cols, rows int) {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) SetActive() {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) SetOptions(o *options.Info) {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) SetName(name string) {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) MergeRows(fromIndex, toIndex int) error {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) MergeCols(fromIndex, toIndex int) error {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) SplitRows(fromIndex, toIndex int) {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) SplitCols(fromIndex, toIndex int) {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) AddConditional(conditional *conditional.Info, refs ...types.Ref) error {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) DeleteConditional(refs ...types.Ref) {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) AutoFilter(ref types.Ref, settings ...interface{}) {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) AddFilter(colIndex int, settings ...interface{}) error {
	panic(errorNotSupported)
}

func (s *sheetWriteStream) DeleteFilter(colIndex int) {
	panic(errorNotSupported)
}
