package xlsx

import (
	"github.com/plandem/xlsx/format"
)

type rowOption func(co *rowOptions)

//rowOptions is a helper type to simplify process of settings options for row
type rowOptions struct {
	outlineLevel uint8
	collapsed    bool
	phonetic     bool
	hidden       bool
	styleID      format.StyleRefID
	height       float32
}

//RowOption is a 'namespace' for all possible options for row
//
// Possible options are:
// OutlineLevel
// Collapsed
// Phonetic
// Hidden
// Formatting
// Height
var RowOption rowOption

//NewRowOptions create and returns option set for row
func NewRowOptions(options ...rowOption) *rowOptions {
	s := &rowOptions{}
	s.Set(options...)
	return s
}

//Set sets new options for option set
func (ro *rowOptions) Set(options ...rowOption) {
	for _, o := range options {
		o(ro)
	}
}

//OutlineLevel set outlining level of the row, when outlining is on.
func (o *rowOption) OutlineLevel(level uint8) rowOption {
	return func(ro *rowOptions) {
		if level < 8 {
			ro.outlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected row is in the collapsed state.
func (o *rowOption) Collapsed(collapsed bool) rowOption {
	return func(ro *rowOptions) {
		ro.collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected row of the worksheet.
func (o *rowOption) Phonetic(phonetic bool) rowOption {
	return func(ro *rowOptions) {
		ro.phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected row are hidden on this worksheet.
func (o *rowOption) Hidden(hidden bool) rowOption {
	return func(ro *rowOptions) {
		ro.hidden = hidden
	}
}

//Formatting sets default style for the affected row. Affects cells not yet allocated in the row. In other words, this style applies to new cells.
func (o *rowOption) Formatting(styleID format.StyleRefID) rowOption {
	return func(ro *rowOptions) {
		ro.styleID = styleID
	}
}

//Height sets the row height in point size. There is no margin padding on row height.
func (o *rowOption) Height(height float32) rowOption {
	return func(ro *rowOptions) {
		ro.height = height
	}
}
