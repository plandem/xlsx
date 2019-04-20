package xlsx

import (
	"github.com/plandem/xlsx/internal/ml"
	"sort"
)

type columns struct {
	sheet *sheetInfo
}

//newColumns creates an object that implements columns functionality (don't miss with Col)
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
		//existing non-grouped column?
		if c.Min == c.Max && c.Min == index {
			data = c
			break
		}
	}

	if data == nil {
		for _, c := range *cols.sheet.ml.Cols {
			//mark grouped column as updated and create non-grouped column with same settings
			if index >= c.Min && index <= c.Max {
				data = &ml.Col{}
				*data = *c
				data.Min = index
				data.Max = index
				data.Updated = false
				c.Updated = true
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

		*cols.sheet.ml.Cols = append(*cols.sheet.ml.Cols, data)
	}

	return data
}

func (cols *columns) Delete(index int) {
	//Cols has 1-based index, but we are using 0-based to unify all indexes at library
	index++

	//N.B.: we can have few columns with same index - grouped and non-grouped, so both should be processed
	for idx, c := range *cols.sheet.ml.Cols {
		if c.Min == c.Max && c.Min == index {
			//non-grouped column should be deleted
			*cols.sheet.ml.Cols = append((*cols.sheet.ml.Cols)[:idx], (*cols.sheet.ml.Cols)[idx+1:]...)
		} else if index >= c.Min && index <= c.Max {
			//grouped column should be only updated
			c.Max--
		}
	}
}

func (cols *columns) pack() *[]*ml.Col {
	//moving grouped column ahead
	packed := *cols.sheet.ml.Cols
	sort.Slice(packed, func(i, j int) bool { return packed[i].Min != packed[i].Max })

	//unpack columns
	unpacked := make(map[int]*ml.Col, len(packed))
	unpackedKeys := make([]int, 0, len(packed))

	for _, c := range packed {
		//add columns with settings only
		if *c != (ml.Col{Min: c.Min, Max: c.Max}) {
			if c.Min == c.Max {
				//we need unique indexes, but packed has mix of grouped/non-grouped columns with intersection of indexes
				if _, ok := unpacked[c.Min]; !ok {
					unpackedKeys = append(unpackedKeys, c.Min)
				}

				unpacked[c.Min] = c
			} else {
				for i := c.Min; i <= c.Max; i++ {
					unpacked[i] = c
					unpackedKeys = append(unpackedKeys, i)
				}
			}
		}
	}

	//pack columns
	var prevCol *ml.Col
	packed = make([]*ml.Col, 0, len(packed))
	sort.Ints(unpackedKeys)

	//cases:
	//same data: 1-10, 2, 3, 11 => 1-11
	//diff data: 1-10, 2, 3, 11 => 1, 2, 3, 4-10, 11
	for _, idx := range unpackedKeys {
		col := unpacked[idx]

		if prevCol == nil {
			prevCol = &ml.Col{}
			*prevCol = *col
			prevCol.Min = idx
			prevCol.Max = idx
			prevCol.Updated = false
			packed = append(packed, prevCol)
		} else {
			colA := *prevCol
			colB := *col
			colA.Min = 0
			colA.Max = 0
			colB.Min = 0
			colB.Max = 0
			colA.Updated = false
			colB.Updated = false

			if colA == colB && idx == prevCol.Max+1 {
				prevCol.Max++
			} else {
				prevCol = &ml.Col{}
				*prevCol = *col
				prevCol.Min = idx
				prevCol.Max = idx
				prevCol.Updated = false
				packed = append(packed, prevCol)
			}
		}
	}

	cols.sheet.ml.Cols = &packed

	//columns must have at least one object
	if len(packed) == 0 {
		return nil
	}

	return cols.sheet.ml.Cols
}
