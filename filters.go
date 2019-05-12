package xlsx

import (
	"github.com/plandem/xlsx/types"
)

//TODO: add custom filters
type filters struct {
	sheet *sheetInfo
}

//newFilters creates an object that implements filters functionality
func newFilters(sheet *sheetInfo) *filters {
	return &filters{sheet: sheet}
}

//Add adds a new filter info for column with 0-based colIndex
func (f *filters) Add(colIndex int, settings []interface{}) error {
	b := f.sheet.ml.AutoFilter.Bounds

	//resolve bounds for auto filter
	if b.IsEmpty() {
		b = types.BoundsFromIndexes(colIndex,0,colIndex,0)
	} else {
		if colIndex < b.FromCol {
			b.FromCol = colIndex
		} else if colIndex > b.ToCol {
			b.ToCol = colIndex
		}
	}

	f.sheet.ml.AutoFilter.Bounds = b
	return nil
}

//Remove removes filter info for column with 0-based colIndex
func (f *filters) Remove(colIndex int) {
	b := f.sheet.ml.AutoFilter.Bounds

	if b.IsEmpty() {
		return
	}

	if colIndex == b.FromCol {
		if b.FromCol == b.ToCol {
			//remove the only column with filter
			b = types.Bounds{}
		} else {
			//move head of filters
			b.FromCol++
		}
	} else if colIndex == b.ToCol {
		//move tail of filters
		b.ToCol--
	}

	f.sheet.ml.AutoFilter.Bounds = b
}
