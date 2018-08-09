package xlsx

import (
	"github.com/plandem/xlsx/options"
	"github.com/plandem/xlsx/types"
)

//max length of excel's sheet name
const sheetNameLimit = 31
const errorNotSupported = "not supported"
const errorNotSupportedWrite = "not supported in read-only mode"
const errorNotSupportedStream = "not supported in stream mode"

type sheetMode byte

const (
	_ sheetMode = 1 << iota
	sheetModeRead
	sheetModeWrite
	sheetModeStream
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
	//Range returns a range for ref
	Range(ref types.Ref) *Range
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
	//Name returns name of sheet
	Name() string
	//SetName sets a name for sheet
	SetName(name string)
	//Set sets options for sheet
	Set(o *options.SheetOptions)
	//SetActive sets the sheet as active
	SetActive()
	//Close frees allocated by sheet resources
	Close()

	//private method to use by internals only
	mode() sheetMode
}
