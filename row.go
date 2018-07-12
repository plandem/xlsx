package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
)

//Row is a higher level object that wraps ml.Row with functionality
type Row struct {
	ml *ml.Row
	*Range
}

//Set sets options for row
func (r *Row) Set(options *rowOptions) {
	if options.height > 0 {
		r.ml.Height = options.height
		r.ml.CustomHeight = true
	}

	r.ml.OutlineLevel = options.outlineLevel
	r.ml.Hidden = options.hidden
	r.ml.Collapsed = options.collapsed
	r.ml.Phonetic = options.phonetic
}

//SetFormatting sets default style for the row. Affects cells not yet allocated in the row. In other words, this style applies to new cells.
func (r *Row) SetFormatting(styleRef format.StyleRefID) {
	r.ml.CustomFormat = true
	r.ml.Style = styleRef
}
