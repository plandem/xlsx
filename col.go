package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/options"
)

//Col is a higher level object that wraps ml.Col with functionality. Inherits functionality of Range
type Col struct {
	ml *ml.Col
	*Range
}

//Cell returns cell of col at row with rowIndex
func (c *Col) Cell(rowIndex int) *Cell {
	return c.sheet.Cell(c.bounds.FromCol, rowIndex)
}

//Set sets options for column
func (c *Col) Set(o *options.ColumnOptions) {
	if o.Width > 0 {
		c.ml.Width = o.Width
		c.ml.CustomWidth = true
	}

	c.ml.OutlineLevel = o.OutlineLevel
	c.ml.Hidden = o.Hidden
	c.ml.Collapsed = o.Collapsed
	c.ml.Phonetic = o.Phonetic
}

//SetFormatting sets default style for the column. Affects cells not yet allocated in the column. In other words, this style applies to new cells.
func (c *Col) SetFormatting(styleID format.StyleID) {
	c.ml.Style = styleID
}

//CopyTo copies col cells into another col with cIdx index.
//N.B.: Merged cells are not supported
func (c *Col) CopyTo(cIdx int, withOptions bool) {
	if withOptions {
		//TODO: copy col options
		panic(errorNotSupported)
	}

	//copy cell data
	c.Range.CopyTo(cIdx, c.Range.bounds.FromRow)
}
