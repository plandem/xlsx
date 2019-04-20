package xlsx

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
)

type mergedCells struct {
	sheet *sheetInfo
}

//newMergedCells creates an object that implements merged cells functionality
func newMergedCells(sheet *sheetInfo) *mergedCells {
	return &mergedCells{sheet: sheet}
}

func (m *mergedCells) initIfRequired() {
	//attach merged cells object if required
	if m.sheet.ml.MergeCells == nil {
		var mergedCells []*ml.MergeCell
		m.sheet.ml.MergeCells = &mergedCells
	}
}

//Resolve check if requested cIdx and rIdx related to merged range and if so, then translate indexes to valid values
func (m *mergedCells) Resolve(cIdx, rIdx int) (int, int, bool) {
	m.initIfRequired()

	merged := false
	mergedCells := *m.sheet.ml.MergeCells
	for _, mc := range mergedCells {
		if merged = mc.Bounds.Contains(cIdx, rIdx); merged {
			cIdx, rIdx = mc.Bounds.FromCol, mc.Bounds.FromRow
			break
		}
	}

	return cIdx, rIdx, merged
}

//Merge adds a merged cells info for bounds
func (m *mergedCells) Add(bounds types.Bounds) error {
	m.initIfRequired()

	//let's check existing merged cells for overlapping
	mergedCells := *m.sheet.ml.MergeCells
	for _, mc := range mergedCells {
		if mc.Bounds.Overlaps(bounds) {
			return errors.New(fmt.Sprintf("intersection of different merged ranges is not allowed, %s intersects with %s", mc.Bounds, bounds))
		}
	}

	//looks like there are no any merged cells in that area, so let's add it
	mergedCells = append(mergedCells, &ml.MergeCell{
		Bounds: bounds,
	})

	m.sheet.ml.MergeCells = &mergedCells
	return nil
}

//Remove removes merged cells info for bounds
func (m *mergedCells) Remove(bounds types.Bounds) {
	m.initIfRequired()

	mergedCells := *m.sheet.ml.MergeCells
	if len(mergedCells) > 0 {
		newMergedCells := make([]*ml.MergeCell, 0, len(mergedCells))

		for _, mc := range mergedCells {
			if !mc.Bounds.Overlaps(bounds) {
				//copy only non overlapping bounds
				newMergedCells = append(newMergedCells, mc)
			}
		}

		m.sheet.ml.MergeCells = &newMergedCells
	}
}

func (m *mergedCells) pack() *[]*ml.MergeCell {
	//merged cells must have at least one object
	if m.sheet.ml.MergeCells != nil && len(*m.sheet.ml.MergeCells) == 0 {
		m.sheet.ml.MergeCells = nil
	}

	return m.sheet.ml.MergeCells
}