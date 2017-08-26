package xlsx

import "github.com/plandem/xlsx/types"

type rangeInfo struct {
	fromCol int
	fromRow int
	toCol   int
	toRow   int
}

func newRangeInfo(r types.Ref) *rangeInfo {
	fromCellRef, toCellRef := r.ReboundIfRequired().ToCellRefs()
	fromCol, fromRow := fromCellRef.ToIndexes()
	toCol, toRow := toCellRef.ToIndexes()

	return &rangeInfo{
		fromCol,
		fromRow,
		toCol,
		toRow,
	}
}

//Contains checks if celRef is inside of range
func (r *rangeInfo) ContainsRef(celRef types.CellRef) bool {
	return r.Contains(celRef.ToIndexes())
}

//ContainsIndexes checks if is inside of range
func (r *rangeInfo) Contains(cellCol, cellRow int) bool {
	return (cellCol >= r.fromCol && cellCol <= r.toCol) && (cellRow >= r.fromRow && cellRow <= r.toRow)
}
