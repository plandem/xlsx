package xlsx

//RowIterator is a interface for iterating rows inside of sheet
type RowIterator interface {
	//Next returns next Row in sheet and corresponding index
	Next() (idx int, row *Row)

	//HasNext returns true if there are rows to iterate or false in other case
	HasNext() bool
}

//rowIterator is object that holds required information for common row's iterator
type rowIterator struct {
	idx   int
	max   int
	sheet Sheet
}

var _ RowIterator = (*rowIterator)(nil)

func newRowIterator(sheet Sheet) RowIterator {
	_, rows := sheet.Dimension()
	return &rowIterator{
		idx:   -1,
		max:   rows - 1,
		sheet: sheet,
	}
}

//Next returns next Row in sheet and corresponding index
func (i *rowIterator) Next() (int, *Row) {
	i.idx++
	return i.idx, i.sheet.Row(i.idx)
}

//HasNext returns true if there are rows to iterate or false in other case
func (i *rowIterator) HasNext() bool {
	return i.idx < i.max
}
