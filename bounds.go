package xlsx

import "github.com/plandem/xlsx/types"

type bounds struct {
	fromCol int
	fromRow int
	toCol   int
	toRow   int
}

//newBoundsFromRef returns a bounds information for provided r Ref
func newBoundsFromRef(r types.Ref) *bounds {
	fromCellRef, toCellRef := r.ToCellRefs()
	fromCol, fromRow := fromCellRef.ToIndexes()
	toCol, toRow := toCellRef.ToIndexes()

	return newBounds(fromCol, toCol, fromRow, toRow)
}

//newBounds returns a bounds information for provided 0-based indexes
func newBounds(fromCol, toCol, fromRow, toRow int) *bounds {
	//rebound cols if required
	if fromCol > toCol {
		toCol, fromCol = fromCol, toCol
	}

	//rebound rows if required
	if fromRow > toRow {
		toRow, fromRow = fromRow, toRow
	}

	return &bounds{
		fromCol,
		fromRow,
		toCol,
		toRow,
	}
}

//ContainsRef checks if celRef is inside of range
func (r *bounds) ContainsRef(celRef types.CellRef) bool {
	return r.Contains(celRef.ToIndexes())
}

//Contains checks if indexes cIdx and rIdx are inside of range
func (r *bounds) Contains(cIdx, rIdx int) bool {
	return (cIdx >= r.fromCol && cIdx <= r.toCol) && (rIdx >= r.fromRow && rIdx <= r.toRow)
}

//Dimension returns total number of cols and rows in range
func (r *bounds) Dimension() (width int, height int) {
	width = r.toCol - r.fromCol
	height = r.toRow - r.fromRow
	return
}
