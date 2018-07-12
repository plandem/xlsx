package rowOption

//rowOption is a 'namespace' for all possible options for row
//
// Possible options are:
// OutlineLevel
// Collapsed
// Phonetic
// Hidden
// Formatting
// Height

import (
	"github.com/plandem/xlsx/format"
)

//OutlineLevel set outlining level of the row, when outlining is on.
func OutlineLevel(level uint8) RowOption {
	return func(ro *RowOptions) {
		if level < 8 {
			ro.OutlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected row is in the collapsed state.
func Collapsed(collapsed bool) RowOption {
	return func(ro *RowOptions) {
		ro.Collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected row of the worksheet.
func Phonetic(phonetic bool) RowOption {
	return func(ro *RowOptions) {
		ro.Phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected row are hidden on this worksheet.
func Hidden(hidden bool) RowOption {
	return func(ro *RowOptions) {
		ro.Hidden = hidden
	}
}

//Formatting sets default style for the affected row. Affects cells not yet allocated in the row. In other words, this style applies to new cells.
func Formatting(styleID format.StyleRefID) RowOption {
	return func(ro *RowOptions) {
		ro.StyleID = styleID
	}
}

//Height sets the row height in point size. There is no margin padding on row height.
func Height(height float32) RowOption {
	return func(ro *RowOptions) {
		ro.Height = height
	}
}
