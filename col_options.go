package xlsx

import (
	"github.com/plandem/xlsx/format"
)

type columnOption func(co *columnOptions)

//columnOptions is a helper type to simplify process of settings options for column
type columnOptions struct {
	outlineLevel uint8
	collapsed    bool
	phonetic     bool
	hidden       bool
	styleID      format.StyleRefID
	width        float32
}

//ColumnOption is a 'namespace' for all possible options for column
//
// Possible options are:
// OutlineLevel
// Collapsed
// Phonetic
// Hidden
// Formatting
// Width
var ColumnOption columnOption

//NewColumnOptions create and returns option set for column
func NewColumnOptions(options ...columnOption) *columnOptions {
	s := &columnOptions{}
	s.Set(options...)
	return s
}

//Set sets new options for option set
func (co *columnOptions) Set(options ...columnOption) {
	for _, o := range options {
		o(co)
	}
}

//OutlineLevel sets outline level of affected column. Range is 0 to 7.
func (o *columnOption) OutlineLevel(level uint8) columnOption {
	return func(co *columnOptions) {
		if level < 8 {
			co.outlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected column is in the collapsed state.
func (o *columnOption) Collapsed(collapsed bool) columnOption {
	return func(co *columnOptions) {
		co.collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected column of the worksheet.
func (o *columnOption) Phonetic(phonetic bool) columnOption {
	return func(co *columnOptions) {
		co.phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected column are hidden on this worksheet.
func (o *columnOption) Hidden(hidden bool) columnOption {
	return func(co *columnOptions) {
		co.hidden = hidden
	}
}

//Formatting sets default style for the affected column. Affects cells not yet allocated in the column. In other words, this style applies to new cells.
func (o *columnOption) Formatting(styleID format.StyleRefID) columnOption {
	return func(co *columnOptions) {
		co.styleID = styleID
	}
}

//Width sets the column width in the same units used by Excel which is: the number of characters in the default font. For more details: https://support.microsoft.com/en-us/kb/214123
func (o *columnOption) Width(width float32) columnOption {
	return func(co *columnOptions) {
		co.width = width
	}
}