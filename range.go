package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/types"
)

//Range is a object that provides some functionality for cells inside of range. E.g.: A1:D12
type Range struct {
	*bounds
	sheet Sheet
}

//newRangeFromRef create and returns Range for requested ref
func newRangeFromRef(sheet Sheet, ref types.Ref) *Range {
	return &Range{
		newBoundsFromRef(ref),
		sheet,
	}
}

//newRange create and returns Range for requested 0-based indexes
func newRange(sheet Sheet, fromCol, toCol, fromRow, toRow int) *Range {
	return &Range{
		newBounds(fromCol, toCol, fromRow, toRow),
		sheet,
	}
}

//Reset resets each cell data into zero state
func (r *Range) Reset() {
	r.Walk(func(idx, cIdx, rIdx int, c *Cell) { c.Reset() })
}

//Clear clears each cell value in range
func (r *Range) Clear() {
	r.Walk(func(idx, cIdx, rIdx int, c *Cell) { c.Clear() })
}

//Cells returns iterator for all cells in range
func (r *Range) Cells() RangeIterator {
	return newRangeIterator(r)
}

//Values returns values for all cells in range
func (r *Range) Values() []string {
	width, height := r.bounds.Dimension()
	values := make([]string, 0, width*height)

	r.Walk(func(idx, cIdx, rIdx int, c *Cell) {
		values = append(values, c.Value())
	})

	return values
}

//Walk calls callback cb for each Cell in range
func (r *Range) Walk(cb func(idx, cIdx, rIdx int, c *Cell)) {
	for idx, cells := 0, r.Cells(); cells.HasNext(); idx++ {
		iCol, iRow, cell := cells.Next()
		cb(idx, iCol, iRow, cell)
	}
}

//SetFormatting sets style format to all cells in range
func (r *Range) SetFormatting(styleRef format.StyleRefID) {
	r.Walk(func(idx, cIdx, rIdx int, c *Cell) {
		c.SetFormatting(styleRef)
	})
}

//CopyTo copies range cells into another range starting with ref
func (r *Range) CopyTo(ref types.Ref) {
	//TODO: check if sheet is opened as read stream and panic about
}
