package xlsx

import (
	"github.com/plandem/ooxml"
)

type hyperlinkManager struct {
	sheet *sheetInfo
	file  *ooxml.PackageFile
}

//newHyperlinkManager creates an object that implements hyperlink cells functionality
func newHyperlinkManager(sheet *sheetInfo) *hyperlinkManager {
	return &hyperlinkManager{
		sheet: sheet,
	}
}

//Resolve check if requested cIdx and rIdx related to merged range and if so, then translate indexes to valid values
//func (m *hyperlinkManager) Resolve(cIdx, rIdx int) (int, int) {
//	//if m.sheet.ml.MergeCells != nil {
//	//	mergedCells := *m.sheet.ml.MergeCells
//	//	if len(mergedCells) > 0 {
//	//		for _, mc := range mergedCells {
//	//			if mc.Bounds.Contains(cIdx, rIdx) {
//	//				cIdx, rIdx = mc.Bounds.FromCol, mc.Bounds.FromRow
//	//				break
//	//			}
//	//		}
//	//	}
//	//}
//	//
//	//return cIdx, rIdx
//}
//
////Resolve check if requested cIdx and rIdx related to merged range and if so, then translate indexes to valid values
//func (m *mergedCellManager) Resolve(cIdx, rIdx int) (int, int) {
//	if m.sheet.ml.MergeCells != nil {
//		mergedCells := *m.sheet.ml.MergeCells
//		if len(mergedCells) > 0 {
//			for _, mc := range mergedCells {
//				if mc.Bounds.Contains(cIdx, rIdx) {
//					cIdx, rIdx = mc.Bounds.FromCol, mc.Bounds.FromRow
//					break
//				}
//			}
//		}
//	}
//
//	return cIdx, rIdx
//}
