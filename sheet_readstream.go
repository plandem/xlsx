// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/format/conditional"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"github.com/plandem/xlsx/types/options/sheet"
	"math"
)

type sheetReadStream struct {
	*sheetInfo
	stream     *ooxml.StreamFileReader
	mergedRows map[int]*Row
	currentRow *ml.Row
}

var _ Sheet = (*sheetReadStream)(nil)

func (s *sheetReadStream) Cell(colIndex, rowIndex int) *Cell {
	var data *ml.Cell

	colIndex, rowIndex, _ = s.mergedCells.Resolve(colIndex, rowIndex)
	row := s.Row(rowIndex)
	data = row.ml.Cells[colIndex]

	if data == nil {
		data = &ml.Cell{
			Ref: types.CellRefFromIndexes(colIndex, rowIndex),
		}
	}

	return &Cell{ml: data, sheet: s.sheetInfo}
}

func (s *sheetReadStream) Row(index int) *Row {
	//if index from cached merged rows, then use it
	if row, ok := s.mergedRows[index]; ok {
		return row
	}

	//skip rows till required index
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

		t, _ := s.stream.Token()
		if t == nil {
			break
		}

		//no more rows?
		if start, ok := t.(xml.StartElement); !ok || !s.loadRow(&start) {
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

func (s *sheetReadStream) loadRow(start *xml.StartElement) bool {
	if start.Name.Local == "row" {
		row := &ml.Row{}
		_ = s.stream.DecodeElement(row, start)

		//expand row dimension to required width
		width, _ := s.Dimension()
		cells := make([]*ml.Cell, width)
		for _, c := range row.Cells {
			//add cell info
			if !isCellEmpty(c) {
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
	_ = s.stream.Close()
}

func (s *sheetReadStream) emptyDataRow(indexRef int) *ml.Row {
	width, _ := s.Dimension()
	return &ml.Row{
		Ref:   indexRef,
		Cells: make([]*ml.Cell, width),
	}
}

//nolint
//afterOpen loads worksheet data and initializes it if required
func (s *sheetReadStream) afterOpen() {
	//adds a styles for types
	s.workbook.doc.styleSheet.addTypedStylesIfRequired()

	s.currentRow = nil
	ignoreDimension := (s.sheetMode & SheetModeIgnoreDimension) != 0
	multiPhase := (s.sheetMode & SheetModeMultiPhase) != 0
	conditionalsInited := false

	if !multiPhase && ignoreDimension {
		panic(fmt.Errorf("to ignore dimension, streaming should be multiphased"))
	}

	if stream, err := s.file.ReadStream(); err != nil {
		panic(err)
	} else {
		s.stream = stream
	}

	var maxCol, maxRow float64

	stop := false

	//first phase
	for {
		if stop {
			break
		}

		t, _ := s.stream.Token()
		if t == nil {
			break
		}

		if start, ok := t.(xml.StartElement); ok {
			switch start.Name.Local {
			case "dimension":
				if !ignoreDimension {
					s.ml.Dimension = &ml.SheetDimension{}
					_ = s.stream.DecodeElement(s.ml.Dimension, &start)
				}
			case "hyperlinks":
				s.hyperlinks = newHyperlinks(s.sheetInfo)
			case "conditionalFormatting":
				if !conditionalsInited {
					//N.B.: conditionalFormatting is not nested, so we have to use flag to init once only
					s.conditionals = newConditionals(s.sheetInfo)
					s.conditionals.initIfRequired()
					conditionalsInited = true
				}
			case "mergeCells":
				s.mergedCells = newMergedCells(s.sheetInfo)
			case "mergeCell":
				cell := &ml.MergeCell{}
				_ = s.stream.DecodeElement(cell, &start)
				s.ml.MergeCells.Items = append(s.ml.MergeCells.Items, cell)
			case "row":
				if !multiPhase {
					stop = true
					s.loadRow(&start)
					break
				}

				row := &ml.Row{}
				_ = s.stream.DecodeElement(row, &start)
				for _, cell := range row.Cells {
					colIndex, _ := types.CellRef(cell.Ref).ToIndexes()
					maxCol = math.Max(maxCol, float64(colIndex))
				}
			}
		}
	}

	// multi phased?
	if multiPhase {
		if ignoreDimension {
			s.ml.Dimension = &ml.SheetDimension{Bounds: types.BoundsFromIndexes(0, 0, int(maxCol), int(maxRow))}
		}

		//skipToFistRow is func to skip any info till first row
		skipToFistRow := func() {
			//close previous opened stream
			_ = s.stream.Close()

			//re-open stream again and cache skip any info till first row
			if stream, err := s.file.ReadStream(); err != nil {
				panic(err)
			} else {
				s.stream = stream
			}

			for {
				t, _ := s.stream.Token()
				if t == nil {
					break
				}

				if start, ok := t.(xml.StartElement); ok && start.Name.Local == "row" {
					s.loadRow(&start)
					break
				}
			}
		}

		//phase2 - reset pointer to rows
		skipToFistRow()

		//cache merged rows
		s.mergedRows = make(map[int]*Row)
		for rows := s.Rows(); rows.HasNext(); {
			rIdx, row := rows.Next()

			for _, mc := range s.ml.MergeCells.Items {
				if rIdx >= mc.Bounds.FromRow && rIdx <= mc.Bounds.ToRow {
					s.mergedRows[rIdx] = row
					break
				}
			}
		}

		//phase3 - reset pointer to rows and clear current row info, because it was polluted while loaded merged cells
		skipToFistRow()
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

func (s *sheetReadStream) SetOptions(o *options.Info) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) SetName(name string) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) MergeRows(fromIndex, toIndex int) error {
	panic(errorNotSupported)
}

func (s *sheetReadStream) MergeCols(fromIndex, toIndex int) error {
	panic(errorNotSupported)
}

func (s *sheetReadStream) SplitRows(fromIndex, toIndex int) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) SplitCols(fromIndex, toIndex int) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) AddConditional(conditional *conditional.Info, refs ...types.Ref) error {
	panic(errorNotSupported)
}

func (s *sheetReadStream) DeleteConditional(refs ...types.Ref) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) AutoFilter(ref types.Ref, settings ...interface{}) {
	panic(errorNotSupported)
}

func (s *sheetReadStream) AddFilter(colIndex int, settings ...interface{}) error {
	panic(errorNotSupported)
}

func (s *sheetReadStream) DeleteFilter(colIndex int) {
	panic(errorNotSupported)
}
