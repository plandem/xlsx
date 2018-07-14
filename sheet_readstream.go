package xlsx

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"log"
)

type sheetReadStream struct {
	*sheetInfo
	stream     *ooxml.StreamFileReader
	rowReader  ooxml.StreamReaderIterator
	currentRow *ml.Row
	multiPhase bool
}

var _ Sheet = (*sheetReadStream)(nil)

func (s *sheetReadStream) Cell(colIndex, rowIndex int) *Cell {
	var data *ml.Cell

	row := s.Row(rowIndex)
	data = row.ml.Cells[colIndex]

	if data == nil {
		data = &ml.Cell{
			Ref: types.CellRefFromIndexes(colIndex, rowIndex),
		}
	}

	return &Cell{ml: data, sheet: s.sheetInfo}
}

func (s *sheetReadStream) CellByRef(cellRef types.CellRef) *Cell {
	cid, rid := cellRef.ToIndexes()
	return s.Cell(cid, rid)
}

func (s *sheetReadStream) Range(ref types.Ref) *Range {
	return newRangeFromRef(s, ref)
}

func (s *sheetReadStream) Row(index int) *Row {
	if s.rowReader == nil {
		return nil
	}

	indexRef := index + 1
	var data *ml.Row

	for {
		if s.currentRow != nil {
			if indexRef == s.currentRow.Ref {
				data = s.currentRow
				break
			}

			if indexRef < s.currentRow.Ref {
				break
			}
		}

		if !s.rowReader(s.nextRow) {
			break
		}
	}

	//looks like there is no any data anymore, return empty row
	if data == nil {
		data = s.emptyDataRow(indexRef)
	}

	return &Row{
		data,
		newRange(s, 0, len(data.Cells)-1, index, index),
	}
}

func (s *sheetReadStream) nextRow(decoder *xml.Decoder, start *xml.StartElement) bool {
	if start != nil && start.Name.Local == "row" {
		row := &ml.Row{}
		decoder.DecodeElement(row, start)

		//expand row dimension to required width
		width, _ := s.Dimension()
		cells := make([]*ml.Cell, width)
		for _, c := range row.Cells {
			//add cell info
			if !s.isCellEmpty(c) {
				iCellCol, _ := c.Ref.ToIndexes()
				cells[iCellCol] = c
			}
		}

		row.Cells = cells
		s.currentRow = row
		return true
	}

	s.currentRow = nil
	return false
}

func (s *sheetReadStream) Rows() RowIterator {
	return newRowIterator(s)
}

//Close frees allocated by sheet resources
func (s *sheetReadStream) Close() {
	s.stream.Close()
}

func (s *sheetReadStream) emptyDataRow(indexRef int) *ml.Row {
	width, _ := s.Dimension()
	return &ml.Row{
		Ref:   indexRef,
		Cells: make([]*ml.Cell, width),
	}
}

//afterOpen loads worksheet data and initializes it if required
func (s *sheetReadStream) afterOpen() {
	if s.currentRow == nil {
		s.stream = s.file.ReadStream()

		//first phase
		for next, hasNext := s.stream.StartIterator(nil); hasNext; {
			hasNext = next(func(decoder *xml.Decoder, start *xml.StartElement) bool {
				switch start.Name.Local {
				case "dimension":
					s.ml.Dimension = &ml.SheetDimension{}
					decoder.DecodeElement(s.ml.Dimension, start)
				case "sheetData":
					log.Println("sheetData")
					return true
				case "mergeCell":
					//log.Println("merged cells!")
					s.ml.MergeCells = &[]*ml.MergeCell{}
					decoder.DecodeElement(s.ml.MergeCells, start)
				case "row":
					//first row found, so stop pre-loading phase
					s.rowReader, _ = s.stream.StartIterator(start)
					return s.multiPhase
				}

				return true
			})
		}

		//second phase
		if !s.multiPhase {
			return
		}
	}
}

//not allowed methods for stream reading mode
func (s *sheetReadStream) Col(index int) *Col {
	panic(errorNotSupported)
}

func (s *sheetReadStream) Cols() ColIterator {
	panic(errorNotSupported)
}

func (s *sheetReadStream) InsertCol(index int) *Col {
	panic(errorNotSupported)
}

func (s *sheetReadStream) InsertRow(index int) *Row {
	panic(errorNotSupported)
}

func (s *sheetReadStream) DeleteRow(index int) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) DeleteCol(index int) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) SetDimension(cols, rows int) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) SetActive() {
	panic(errorNotSupported)
}

func (s *sheetReadStream) SetState(state types.VisibilityType) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) SetName(name string) {
	panic(errorNotSupported)
}
