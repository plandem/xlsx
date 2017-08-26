package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
)

//Range is a object that provides some functionality for cells inside of range. E.g.: A1:D12
type Range struct {
	*rangeInfo
	sheet *Sheet
}

type rangeWalkerFn func(c *ml.Cell)

var rangeClearFn = func(c *ml.Cell) { c.Value = "" }
var rangeResetFn = func(c *ml.Cell) { *c = ml.Cell{} }

//walk walks for range and call callback for each cell
func (r *Range) walk(cb rangeWalkerFn) {
	for _, row := range r.sheet.ml.SheetData {
		for _, c := range row.Cells {
			if c != nil && r.ContainsRef(c.Ref) {
				cb(c)
			}
		}
	}
}

//Reset resets each cell data into zero state
func (r *Range) Reset() {
	r.walk(rangeResetFn)
}

//Clear clears each cell value in range
func (r *Range) Clear() {
	r.walk(rangeClearFn)
}

//Cells returns all cells in range
func (r *Range) Cells() []*Cell {
	r.sheet.resolveDimension(false)

	width := r.toCol - r.fromCol
	height := r.toRow - r.fromRow

	cells := make([]*Cell, 0, width*height)

	for rid := r.fromRow; rid <= r.toRow; rid++ {
		for cid := r.fromCol; cid <= r.toCol; cid++ {
			cells = append(cells, r.sheet.Cell(cid, rid))
		}
	}

	return cells
}

//Values returns values for all cells in range
func (r *Range) Values() []string {
	cells := r.Cells()
	values := make([]string, len(cells))
	for i, cell := range cells {
		values[i] = cell.Value()
	}

	return values
}

//SetFormatting sets style format to all cells in range
func (r *Range) SetFormatting(styleRef format.StyleRefID) {
	for _, cell := range r.Cells() {
		cell.SetFormatting(styleRef)
	}
}
