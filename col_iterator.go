package xlsx

//ColIterator is a interface for iterating cols inside of sheet
type ColIterator interface {
	//Next returns next Col in sheet and corresponding index
	Next() (*Col, int)

	//HasNext returns true if there are cols to iterate or false in other case
	HasNext() bool
}

//colIterator is object that holds required information for common col's iterator
type colIterator struct {
	idx   int
	max   int
	sheet *Sheet
}

var _ ColIterator = (*colIterator)(nil)

func newColIterator(sheet *Sheet) ColIterator {
	return &colIterator{
		idx:   -1,
		max:   sheet.TotalCols() - 1,
		sheet: sheet,
	}
}

//Next returns next Col in sheet and corresponding index
func (i *colIterator) Next() (*Col, int) {
	i.idx++
	return i.sheet.Col(i.idx), i.idx
}

//HasNext returns true if there are cols to iterate or false in other case
func (i *colIterator) HasNext() bool {
	return i.idx < i.max
}
