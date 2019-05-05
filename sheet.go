package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/options"
	"github.com/plandem/xlsx/types"
)

const errorNotSupported = "not supported"
const errorNotSupportedWrite = "not supported in read-only mode"
const errorNotSupportedStream = "not supported in stream mode"

type sheetMode byte

//List of all possible open modes for Sheet. Mode applies only once, except SheetModeStream and few modes can be combined. E.g.: SheetModeStream, SheetModeMultiPhase
const (
	sheetModeUnknown sheetMode = 0
	sheetModeRead    sheetMode = 1 << iota
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
	//MergeRows merges rows between fromIndex and toIndex
	MergeRows(fromIndex, toIndex int) error
	//MergeCols merges cols between fromIndex and toIndex
	MergeCols(fromIndex, toIndex int) error
	//SplitRows splits rows between fromIndex and toIndex
	SplitRows(fromIndex, toIndex int)
	//SplitCols splits cols between fromIndex and toIndex
	SplitCols(fromIndex, toIndex int)
	//AddConditional adds conditional formatting to sheet
	AddConditional(conditional *format.ConditionalFormat, refs ...types.Ref) error
	//DeleteConditional deletes conditional formatting for refs
	DeleteConditional(refs ...types.Ref)
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

	//private methods to use by internals only
	mode() sheetMode
	info() *sheetInfo
}
