package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/row_option"
)

//Row is a higher level object that wraps ml.Row with functionality
type Row struct {
	ml *ml.Row
	*Range
}

//Set sets options for row
func (r *Row) Set(options *rowOption.RowOptions) {
	if options.Height > 0 {
		r.ml.Height = options.Height
		r.ml.CustomHeight = true
	}

	r.ml.OutlineLevel = options.OutlineLevel
	r.ml.Hidden = options.Hidden
	r.ml.Collapsed = options.Collapsed
	r.ml.Phonetic = options.Phonetic
}

//SetFormatting sets default style for the row. Affects cells not yet allocated in the row. In other words, this style applies to new cells.
func (r *Row) SetFormatting(styleRef format.StyleRefID) {
	r.ml.CustomFormat = true
	r.ml.Style = styleRef
}
