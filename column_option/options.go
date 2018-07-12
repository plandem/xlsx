package columnOption

//columnOption is a 'namespace' for all possible options for column
//
// Possible options are:
// OutlineLevel
// Collapsed
// Phonetic
// Hidden
// Formatting
// Width

import (
	"github.com/plandem/xlsx/format"
)

//OutlineLevel sets outline level of affected column. Range is 0 to 7.
func OutlineLevel(level uint8) Option {
	return func(co *Options) {
		if level < 8 {
			co.OutlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected column is in the collapsed state.
func Collapsed(collapsed bool) Option {
	return func(co *Options) {
		co.Collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected column of the worksheet.
func Phonetic(phonetic bool) Option {
	return func(co *Options) {
		co.Phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected column are hidden on this worksheet.
func Hidden(hidden bool) Option {
	return func(co *Options) {
		co.Hidden = hidden
	}
}

//Formatting sets default style for the affected column. Affects cells not yet allocated in the column. In other words, this style applies to new cells.
func Formatting(styleID format.StyleRefID) Option {
	return func(co *Options) {
		co.StyleID = styleID
	}
}

//Width sets the column width in the same units used by Excel which is: the number of characters in the default font. For more details: https://support.microsoft.com/en-us/kb/214123
func Width(width float32) Option {
	return func(co *Options) {
		co.Width = width
	}
}
