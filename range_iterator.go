package xlsx

//RangeIterator is a interface for iterating cells inside of range
type RangeIterator interface {
	//Next returns next Cell in range and corresponding indexes
	Next() (cIdx int, rIdx int, cell *Cell)

	//HasNext returns true if there are cells to iterate or false in other case
	HasNext() bool
}

//rangeIterator is object that holds required information for range's iterator
type rangeIterator struct {
	r    *Range
	cIdx int
	rIdx int
}

var _ RangeIterator = (*rangeIterator)(nil)

func newRangeIterator(r *Range) RangeIterator {
	return &rangeIterator{
		r:    r,
		cIdx: r.bounds.FromCol,
		rIdx: r.bounds.FromRow,
	}
}

//Next returns next Cell in range and corresponding indexes
func (i *rangeIterator) Next() (cIdx int, rIdx int, cell *Cell) {
	cIdx, rIdx, cell = i.cIdx, i.rIdx, i.r.sheet.Cell(i.cIdx, i.rIdx)

	i.cIdx++
	if i.cIdx > i.r.bounds.ToCol {
		i.cIdx = i.r.bounds.FromCol
		i.rIdx++
	}

	return
}

//HasNext returns true if there are cells to iterate or false in other case
func (i *rangeIterator) HasNext() bool {
	return i.rIdx <= i.r.bounds.ToRow
}
