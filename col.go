package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
)

//Col is a higher level object that wraps ml.Col with functionality
type Col struct {
	ml *ml.Col
	*Range
}

//Set sets options for column
func (c *Col) Set(options *internal.ColumnOptions) {
	if options.Width > 0 {
		c.ml.Width = options.Width
		c.ml.CustomWidth = true
	}

	c.ml.OutlineLevel = options.OutlineLevel
	c.ml.Hidden = options.Hidden
	c.ml.Collapsed = options.Collapsed
	c.ml.Phonetic = options.Phonetic
}

//SetFormatting sets default style for the column. Affects cells not yet allocated in the column. In other words, this style applies to new cells.
func (c *Col) SetFormatting(styleRef format.StyleRefID) {
	c.ml.Style = styleRef
}
