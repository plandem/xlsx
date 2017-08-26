package internal

import (
	"github.com/plandem/xlsx/format"
)

type RowOption func(co *RowOptions)

type RowOptions struct {
	OutlineLevel uint8
	Collapsed    bool
	Phonetic     bool
	Hidden       bool
	StyleID      format.StyleRefID
	Height       float32
}

//Set sets new options for option set
func (row *RowOptions) Set(options ...RowOption) {
	for _, o := range options {
		o(row)
	}
}

//OutlineLevel set outlining level of the row, when outlining is on.
func (o *RowOption) OutlineLevel(level uint8) RowOption {
	return func(ro *RowOptions) {
		if level < 8 {
			ro.OutlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected row is in the collapsed state.
func (o *RowOption) Collapsed(collapsed bool) RowOption {
	return func(ro *RowOptions) {
		ro.Collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected row of the worksheet.
func (o *RowOption) Phonetic(phonetic bool) RowOption {
	return func(ro *RowOptions) {
		ro.Phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected row are hidden on this worksheet.
func (o *RowOption) Hidden(hidden bool) RowOption {
	return func(ro *RowOptions) {
		ro.Hidden = hidden
	}
}

//Formatting sets default style for the affected row. Affects cells not yet allocated in the row. In other words, this style applies to new cells.
func (o *RowOption) Formatting(styleID format.StyleRefID) RowOption {
	return func(ro *RowOptions) {
		ro.StyleID = styleID
	}
}

//Height sets the row height in point size. There is no margin padding on row height.
func (o *RowOption) Height(height float32) RowOption {
	return func(ro *RowOptions) {
		ro.Height = height
	}
}
