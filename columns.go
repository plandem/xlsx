// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/xlsx/internal/ml"
)

type columns struct {
	sheet *sheetInfo
}

//newColumns creates an object that implements columns functionality (don't miss with Col)
func newColumns(sheet *sheetInfo) *columns {
	return &columns{sheet: sheet}
}

func (cols *columns) Resolve(index int) *ml.Col {
	var data *ml.Col

	//Cols has 1-based index, but we are using 0-based to unify all indexes at library
	index++

	for _, c := range cols.sheet.ml.Cols.Items {
		//existing non-grouped column?
		if c.Min == c.Max && c.Min == index {
			data = c
			break
		}
	}

	if data == nil {
		for _, c := range cols.sheet.ml.Cols.Items {
			//mark grouped column as updated and create non-grouped column with same settings
			if index >= c.Min && index <= c.Max {
				data = &ml.Col{}
				*data = *c
				data.Min = index
				data.Max = index
				break
			}
		}

		//if there was no any grouped column, then create a new one non-grouped
		if data == nil {
			data = &ml.Col{
				Min: index,
				Max: index,
			}
		}

		cols.sheet.ml.Cols.Items = append(cols.sheet.ml.Cols.Items, data)
	}

	return data
}

func (cols *columns) Delete(index int) {
	//Cols has 1-based index, but we are using 0-based to unify all indexes at library
	index++

	//N.B.: we can have few columns with same index - grouped and non-grouped, so both should be processed
	for idx, c := range cols.sheet.ml.Cols.Items {
		if c.Min == c.Max && c.Min == index {
			//non-grouped column should be deleted
			cols.sheet.ml.Cols.Items = append(cols.sheet.ml.Cols.Items[:idx], cols.sheet.ml.Cols.Items[idx+1:]...)
		} else if index >= c.Min && index <= c.Max {
			//grouped column should be only updated
			c.Max--
		}
	}
}
