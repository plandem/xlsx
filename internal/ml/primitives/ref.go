package primitives

import (
	"strings"
)

//Ref is a type to encode XSD ST_Ref, a reference that identifies a cell or a range of cells. E.g.: N28 or B5:N10
type Ref string

//ToCellRefs returns from/to CellRef of Ref
func (r Ref) ToCellRefs() (from CellRef, to CellRef) {
	cellRefs := strings.Split(string(r), ":")

	if len(cellRefs) == 1 {
		from = CellRef("A1")
		to = CellRef(cellRefs[0])
	} else {
		from = CellRef(cellRefs[0])
		to = CellRef(cellRefs[1])
	}

	return
}

//ToBounds returns related bounds of Ref
func (r Ref) ToBounds() Bounds {
	from, to := r.ToCellRefs()
	fromCol, fromRow := from.ToIndexes()
	toCol, toRow := to.ToIndexes()

	return BoundsFromIndexes(fromCol, fromRow, toCol, toRow)
}

//RefFromCellRefs returns Ref for from/to CellRefs
func RefFromCellRefs(from CellRef, to CellRef) Ref {
	if from == to {
		return Ref(from)
	}

	return Ref(string(from) + ":" + string(to))
}
