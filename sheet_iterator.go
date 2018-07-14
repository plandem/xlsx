package xlsx

//SheetIterator is a interface for iterating sheets inside of Spreadsheet
type SheetIterator interface {
	//Next returns next Sheet in Spreadsheet and corresponding index
	Next() (idx int, sheet Sheet)

	//HasNext returns true if there are sheets to iterate or false in other case
	HasNext() bool
}

//sheetIterator is object that holds required information for common sheet's iterator
type sheetIterator struct {
	idx int
	max int
	xl  *Spreadsheet
}

var _ SheetIterator = (*sheetIterator)(nil)

func newSheetIterator(xl *Spreadsheet) SheetIterator {
	return &sheetIterator{
		xl:  xl,
		idx: -1,
		max: len(xl.sheets) - 1,
	}
}

//Next returns next Sheet in Spreadsheet and corresponding index
func (i *sheetIterator) Next() (int, Sheet) {
	i.idx++
	return i.idx, i.xl.Sheet(i.idx)
}

//HasNext returns true if there are sheets to iterate or false in other case
func (i *sheetIterator) HasNext() bool {
	return i.idx < i.max
}
