package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
)

//Col is a higher level object that wraps ml.Col with functionality
type Col struct {
	ml *ml.Col
	*Range
}

//Set sets options for column
func (c *Col) Set(options *columnOptions) {
	if options.width > 0 {
		c.ml.Width = options.width
		c.ml.CustomWidth = true
	}

	c.ml.OutlineLevel = options.outlineLevel
	c.ml.Hidden = options.hidden
	c.ml.Collapsed = options.collapsed
	c.ml.Phonetic = options.phonetic
}

//SetFormatting sets default style for the column. Affects cells not yet allocated in the column. In other words, this style applies to new cells.
func (c *Col) SetFormatting(styleRef format.StyleRefID) {
	c.ml.Style = styleRef
}
