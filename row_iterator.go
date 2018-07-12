package xlsx

//RowIterator is a interface for iterating rows inside of sheet
type RowIterator interface {
	//Next returns next Row in sheet and corresponding index
	Next() (*Row, int)

	//HasNext returns true if there are rows to iterate or false in other case
	HasNext() bool
}

//rowIterator is object that holds required information for common row's iterator
type rowIterator struct {
	idx   int
	max   int
	sheet  *Sheet
}

var _ RowIterator = (*rowIterator)(nil)

func newRowIterator(sheet *Sheet) RowIterator {
	return &rowIterator{
		idx:  -1,
		max: sheet.TotalRows() - 1,
		sheet:  sheet,
	}
}

//Next returns next Row in sheet and corresponding index
func (i *rowIterator) Next() (*Row, int) {
	i.idx++
	return i.sheet.Row(i.idx), i.idx
}

//HasNext returns true if there are rows to iterate or false in other case
func (i *rowIterator) HasNext() bool {
	return i.idx < i.max
}
