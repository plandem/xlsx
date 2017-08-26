package internal

import (
	"github.com/plandem/xlsx/format"
)

type ColumnOption func(co *ColumnOptions)

type ColumnOptions struct {
	OutlineLevel uint8
	Collapsed    bool
	Phonetic     bool
	Hidden       bool
	StyleID      format.StyleRefID
	Width        float32
}

//Set sets new options for option set
func (col *ColumnOptions) Set(options ...ColumnOption) {
	for _, o := range options {
		o(col)
	}
}

//OutlineLevel sets outline level of affected column. Range is 0 to 7.
func (o *ColumnOption) OutlineLevel(level uint8) ColumnOption {
	return func(co *ColumnOptions) {
		if level < 8 {
			co.OutlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected column is in the collapsed state.
func (o *ColumnOption) Collapsed(collapsed bool) ColumnOption {
	return func(co *ColumnOptions) {
		co.Collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected column of the worksheet.
func (o *ColumnOption) Phonetic(phonetic bool) ColumnOption {
	return func(co *ColumnOptions) {
		co.Phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected column are hidden on this worksheet.
func (o *ColumnOption) Hidden(hidden bool) ColumnOption {
	return func(co *ColumnOptions) {
		co.Hidden = hidden
	}
}

//Formatting sets default style for the affected column. Affects cells not yet allocated in the column. In other words, this style applies to new cells.
func (o *ColumnOption) Formatting(styleID format.StyleRefID) ColumnOption {
	return func(co *ColumnOptions) {
		co.StyleID = styleID
	}
}

//Width sets the column width in the same units used by Excel which is: the number of characters in the default font. For more details: https://support.microsoft.com/en-us/kb/214123
func (o *ColumnOption) Width(width float32) ColumnOption {
	return func(co *ColumnOptions) {
		co.Width = width
	}
}
