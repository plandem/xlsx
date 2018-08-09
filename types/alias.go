package types

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//Bounds is alias of original primitives.Bounds type make it public.
type Bounds = primitives.Bounds

//Ref is alias of original primitives.Ref type make it public.
type Ref = primitives.Ref

//CellRef is alias of original primitives.CellRef type make it public.
type CellRef = primitives.CellRef

//Text is alias of original primitives.Text type make it public.
type Text = primitives.Text

//RefFromCellRefs is alias of original primitives.RefFromCellRefs to make it public
func RefFromCellRefs(from CellRef, to CellRef) Ref {
	return primitives.RefFromCellRefs(from, to)
}

//CellRefFromIndexes is alias of original primitives.CellRefFromIndexes to make it public
func CellRefFromIndexes(colIndex, rowIndex int) CellRef {
	return primitives.CellRefFromIndexes(colIndex, rowIndex)
}

//BoundsFromIndexes is alias of original primitives.BoundsFromIndexes to make it public
func BoundsFromIndexes(fromCol, fromRow, toCol, toRow int) Bounds {
	return primitives.BoundsFromIndexes(fromCol, fromRow, toCol, toRow)
}
