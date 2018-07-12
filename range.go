package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/types"
)

//Range is a object that provides some functionality for cells inside of range. E.g.: A1:D12
type Range struct {
	*bounds
	sheet *Sheet
}

//newRangeFromRef create and returns Range for requested ref
func newRangeFromRef(sheet *Sheet, ref types.Ref) *Range {
	return &Range{
		newBoundsFromRef(ref),
		sheet,
	}
}

//newRange create and returns Range for requested 0-based indexes
func newRange(sheet *Sheet, fromCol, toCol, fromRow, toRow int) *Range {
	return &Range{
		newBounds(fromCol, toCol, fromRow, toRow),
		sheet,
	}
}

//Reset resets each cell data into zero state
func (r *Range) Reset() {
	r.Walk(func(c *Cell) { c.Reset() })
}

//Clear clears each cell value in range
func (r *Range) Clear() {
	r.Walk(func(c *Cell) { c.Clear() })
}

//Cells returns all cells in range
func (r *Range) Cells() []*Cell {
	width, height := r.bounds.Size()
	cells := make([]*Cell, 0, width*height)

	r.Walk(func(c *Cell) {
		cells = append(cells, c)
	})

	return cells
}

//Values returns values for all cells in range
func (r *Range) Values() []string {
	width, height := r.bounds.Size()
	values := make([]string, 0, width*height)

	r.Walk(func(c *Cell) {
		values = append(values, c.Value())
	})

	return values
}

//SetFormatting sets style format to all cells in range
func (r *Range) SetFormatting(styleRef format.StyleRefID) {
	r.Walk(func(c *Cell) {
		c.SetFormatting(styleRef)
	})
}

//Walk calls callback cb for each Cell in range
func (r *Range) Walk(cb func(c *Cell)) {
	for cells := r.Iterator(); cells.HasNext(); {
		_, _, cell := cells.Next()
		cb(cell)
	}
}

//Iterator returns iterator for all cells in range
func (r *Range) Iterator() RangeIterator {
	return newRangeIterator(r)
}
