package xlsx

import (
	"github.com/plandem/xlsx/internal/ml"
)

//TODO:
// 1) pack same info into one object
// 2) remove empty objects
type columns struct {
	sheet *sheetInfo
}

//newColumns creates an object that implements columns functionality (don't miss with col)
func newColumns(sheet *sheetInfo) *columns {
	//attach columns object if required
	if sheet.ml.Cols == nil {
		var cols []*ml.Col
		sheet.ml.Cols = &cols
	}

	return &columns{sheet: sheet}
}

func (cols *columns) Resolve(index int) *ml.Col {
	var data *ml.Col

	//Cols has 1-based index, but we are using 0-based to unify all indexes at library
	index++

	for _, c := range *cols.sheet.ml.Cols {
		if c.Min == c.Max && c.Min == index {
			data = c
			break
		}
	}

	if data == nil {
		data = &ml.Col{
			Min: index,
			Max: index,
		}

		*cols.sheet.ml.Cols = append(*cols.sheet.ml.Cols, data)
	}

	return data
}

func (cols *columns) Delete(index int) {
	//Cols has 1-based index, but we are using 0-based to unify all indexes at library
	index++

	for idx, c := range *cols.sheet.ml.Cols {
		if c.Min == c.Max && c.Min == index {
			*cols.sheet.ml.Cols = append((*cols.sheet.ml.Cols)[:idx], (*cols.sheet.ml.Cols)[idx+1:]...)
		}
	}
}

func (cols *columns) BeforeMarshalXML() *[]*ml.Col {
	//columns must have at least one object
	if cols.sheet.ml.Cols != nil && len(*cols.sheet.ml.Cols) == 0 {
		return nil
	}

	return nil
}
