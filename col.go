package xlsx

import (
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
)

//Col is a higher level object that wraps ml.Col with functionality
type Col struct {
	ml    *ml.Col
	sheet *Sheet
	index int
}

//walk walks for col range and call callback for each cell
func (c *Col) walk(cb rangeWalkerFn) {
	for _, row := range c.sheet.ml.SheetData {
		cell := row.Cells[c.index]
		if cell != nil {
			cb(cell)
		}
	}
}

//Reset resets each cell data into zero state
func (c *Col) Reset() {
	c.walk(rangeResetFn)
}

//Clear clears each cell value
func (c *Col) Clear() {
	c.walk(rangeClearFn)
}

//Cells returns all cells in col
func (c *Col) Cells() []*Cell {
	c.sheet.resolveDimension(false)
	_, toRef := c.sheet.ml.Dimension.Ref.ToCellRefs()
	_, maxHeight := toRef.ToIndexes()

	cells := make([]*Cell, maxHeight + 1)
	for rid := 0; rid <= maxHeight; rid++ {
		cells[rid] = c.sheet.Cell(c.index, rid)
	}

	return cells
}

//Values returns values for all cells in col
func (c *Col) Values() []string {
	cells := c.Cells()
	values := make([]string, len(cells))
	for i, cell := range cells {
		values[i] = cell.Value()
	}

	return values
}

//Set sets options for column
func (c *Col) Set(options *internal.ColumnOptions) {
	if options.Width > 0 {
		c.ml.Width = options.Width
		c.ml.CustomWidth = true
	}

	if options.StyleID > 0 {
		c.ml.Style = options.StyleID
	}

	c.ml.OutlineLevel = options.OutlineLevel
	c.ml.Hidden = options.Hidden
	c.ml.Collapsed = options.Collapsed
	c.ml.Phonetic = options.Phonetic
}
