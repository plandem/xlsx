// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/xlsx/format/conditional"
	"github.com/plandem/xlsx/types"
	"github.com/plandem/xlsx/types/options/sheet"
)

const errorNotSupported = "not supported"
const errorNotSupportedWrite = "not supported in read-only mode"
const errorNotSupportedStream = "not supported in stream mode"

type SheetMode byte

//List of all possible open modes for Sheet. Mode applies only once, except SheetModeStream and few modes can be combined. E.g.: SheetModeStream, SheetModeMultiPhase
const (
	sheetModeUnknown SheetMode = 0
	sheetModeRead    SheetMode = 1 << iota
	sheetModeWrite
	SheetModeStream          //In stream mode only forward reading/writing is allowed
	SheetModeMultiPhase      //Sheet will be iterated two times: first one to load meta information (e.g. merged cells) and another one for sheet data. Only for SheetModeStream mode.
	SheetModeIgnoreDimension //Ignore dimension information during reading or skip it during writing
)

//Sheet is interface for a higher level object that wraps ml.Worksheet with functionality
type Sheet interface {
	//Cell returns a cell for 0-based indexes
	Cell(colIndex, rowIndex int) *Cell
	//CellByRef returns a cell for ref
	CellByRef(cellRef types.CellRef) *Cell
	//Rows returns iterator for all rows of sheet
	Rows() RowIterator
	//Row returns a row for 0-based index
	Row(index int) *Row
	//Cols returns iterator for all cols of sheet
	Cols() ColIterator
	//Col returns a col for 0-based index
	Col(index int) *Col
	//Range returns a range for indexes
	Range(fromCol, fromRow, toCol, toRow int) *Range
	//RangeByRef returns a range for ref
	RangeByRef(ref types.Ref) *Range
	//Dimension returns total number of cols and rows in sheet
	Dimension() (cols int, rows int)
	//SetDimension sets total number of cols and rows in sheet
	SetDimension(cols, rows int)
	//InsertRow inserts a row at 0-based index and returns it. Using to insert a row between other rows.
	InsertRow(index int) *Row
	//DeleteRow deletes a row at 0-based index
	DeleteRow(index int)
	//InsertCol inserts a col at 0-based index and returns it. Using to insert a col between other cols.
	InsertCol(index int) *Col
	//DeleteCol deletes a col at 0-based index
	DeleteCol(index int)
	//MergeRows merges rows between 0-based fromIndex and toIndex
	MergeRows(fromIndex, toIndex int) error
	//MergeCols merges cols between 0-based fromIndex and toIndex
	MergeCols(fromIndex, toIndex int) error
	//SplitRows splits rows between 0-based fromIndex and toIndex
	SplitRows(fromIndex, toIndex int)
	//SplitCols splits cols between 0-based fromIndex and toIndex
	SplitCols(fromIndex, toIndex int)
	//AddConditional adds conditional formatting to sheet, with additional refs if required
	AddConditional(conditional *conditional.Info, refs ...types.Ref) error
	//DeleteConditional deletes conditional formatting for refs
	DeleteConditional(refs ...types.Ref)
	//AutoFilter adds auto filter in provided Ref range with additional settings if required
	AutoFilter(ref types.Ref, settings ...interface{})
	//AddFilter adds a custom filter to column with 0-based colIndex
	AddFilter(colIndex int, settings ...interface{}) error
	//DeleteFilter deletes a filter from column with 0-based colIndex
	DeleteFilter(colIndex int)
	//Name returns name of sheet
	Name() string
	//SetName sets a name for sheet
	SetName(name string)
	//Set sets options for sheet
	SetOptions(o *options.Info)
	//SetActive sets the sheet as active
	SetActive()
	//Close frees allocated by sheet resources
	Close()

	//private methods to use by internals only
	mode() SheetMode
	info() *sheetInfo
}
