package xlsx

//TODO: implement add/remove/refresh functionality

type mergedCellManager struct {
	sheet *sheetInfo
}

//newMergedCellManager creates an object that implements merged cells functionality
func newMergedCellManager(sheet *sheetInfo) *mergedCellManager {
	return &mergedCellManager{
		sheet: sheet,
	}
}

//Resolve check if requested cIdx and rIdx related to merged range and if so, then translate indexes to valid values
func (m *mergedCellManager) Resolve(cIdx, rIdx int) (int, int) {
	if m.sheet.ml.MergeCells != nil {
		mergedCells := *m.sheet.ml.MergeCells
		if len(mergedCells) > 0 {
			for _, mc := range mergedCells {
				if mc.Bounds.Contains(cIdx, rIdx) {
					cIdx, rIdx = mc.Bounds.FromCol, mc.Bounds.FromRow
					break
				}
			}
		}
	}

	return cIdx, rIdx
}
