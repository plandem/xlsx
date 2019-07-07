package xlsx

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
)

//TODO: add support for custom filters and sorting
type filters struct {
	sheet *sheetInfo
}

//newFilters creates an object that implements filters functionality
func newFilters(sheet *sheetInfo) *filters {
	return &filters{sheet: sheet}
}

func (f *filters) initIfRequired() {
	//attach custom filters if required
	if f.sheet.ml.AutoFilter.FilterColumn == nil {
		var filters []*ml.FilterColumn
		f.sheet.ml.AutoFilter.FilterColumn = &filters
	}
}

//AddAuto adds auto filter for range
func (f *filters) AddAuto(ref types.Ref, settings []interface{}) {
	b := ref.ToBounds()

	//adjust auto filter range to include custom filters if required
	for _, fInfo := range *f.sheet.ml.AutoFilter.FilterColumn {
		if fInfo.ColId < b.FromCol {
			b.FromCol = fInfo.ColId
		} else if fInfo.ColId > b.ToCol {
			b.ToCol = fInfo.ColId
		}
	}

	//TODO: add support for 'sort' settings
	f.sheet.ml.AutoFilter.Bounds = b
}

//Add adds a new custom filter info for column with 0-based colIndex
func (f *filters) Add(colIndex int, settings []interface{}) error {
	f.initIfRequired()
	b := f.sheet.ml.AutoFilter.Bounds

	//resolve bounds for auto filter
	if b.IsEmpty() {
		b = types.BoundsFromIndexes(colIndex, 0, colIndex, 0)
	} else {
		if colIndex < b.FromCol {
			b.FromCol = colIndex
		} else if colIndex > b.ToCol {
			b.ToCol = colIndex
		}
	}

	//TODO: add custom filter
	f.sheet.ml.AutoFilter.Bounds = b
	return nil
}

//Remove removes filter info for column with 0-based colIndex
func (f *filters) Remove(colIndex int) {
	f.initIfRequired()
	b := f.sheet.ml.AutoFilter.Bounds

	if b.IsEmpty() {
		return
	}

	//adjust auto filter range
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

	//remove custom filter if required
	filters := *f.sheet.ml.AutoFilter.FilterColumn
	for i, fInfo := range filters {
		if fInfo.ColId == colIndex {
			filters = append(filters[:i], filters[i+1:]...)
		}
	}

	f.sheet.ml.AutoFilter.Bounds = b
	f.sheet.ml.AutoFilter.FilterColumn = &filters
}

func (f *filters) pack() {
	//custom filters must have at least one object
	if f.sheet.ml.AutoFilter.FilterColumn != nil && len(*f.sheet.ml.AutoFilter.FilterColumn) == 0 {
		f.sheet.ml.AutoFilter.FilterColumn = nil
	}
}
