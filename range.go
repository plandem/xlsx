package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/types"
)

//Range is a object that provides some functionality for cells inside of range. E.g.: A1:D12
type Range struct {
	//we don't want to pollute Range with bound's public properties
	bounds types.Bounds
	sheet  Sheet
}

//newRangeFromRef create and returns Range for requested ref
func newRangeFromRef(sheet Sheet, ref types.Ref) *Range {
	return &Range{
		ref.ToBounds(),
		sheet,
	}
}

//newRange create and returns Range for requested 0-based indexes
func newRange(sheet Sheet, fromCol, toCol, fromRow, toRow int) *Range {
	return &Range{
		types.BoundsFromIndexes(fromCol, fromRow, toCol, toRow),
		sheet,
	}
}

//Bounds returns bounds of range
func (r *Range) Bounds() types.Bounds {
	return r.bounds
}

//Reset resets each cell data into zero state
func (r *Range) Reset() {
	r.Walk(func(idx, cIdx, rIdx int, c *Cell) { c.Reset() })
}

//Clear clears each cell value in range
func (r *Range) Clear() {
	r.Walk(func(idx, cIdx, rIdx int, c *Cell) { c.Clear() })
}

//Cells returns iterator for all cells in range
func (r *Range) Cells() RangeIterator {
	return newRangeIterator(r)
}

//Values returns values for all cells in range
func (r *Range) Values() []string {
	width, height := r.bounds.Dimension()
	values := make([]string, 0, width*height)

	r.Walk(func(idx, cIdx, rIdx int, c *Cell) {
		values = append(values, c.Value())
	})

	return values
}

//Walk calls callback cb for each Cell in range
func (r *Range) Walk(cb func(idx, cIdx, rIdx int, c *Cell)) {
	for idx, cells := 0, r.Cells(); cells.HasNext(); idx++ {
		iCol, iRow, cell := cells.Next()
		cb(idx, iCol, iRow, cell)
	}
}

//SetFormatting sets style format to all cells in range
func (r *Range) SetFormatting(styleID format.StyleID) {
	r.Walk(func(idx, cIdx, rIdx int, c *Cell) {
		c.SetFormatting(styleID)
	})
}

//CopyToRef copies range cells into another range starting with ref.
//N.B.: Merged cells are not supported
func (r *Range) CopyToRef(ref types.Ref) {
	target := ref.ToBounds()
	r.CopyTo(target.ToCol, target.ToRow)
}

//CopyTo copies range cells into another range starting indexes cIdx and rIdx
//N.B.: Merged cells are not supported
func (r *Range) CopyTo(cIdx, rIdx int) {
	//result is unpredictable in stream mode
	if mode := r.sheet.mode(); (mode & sheetModeStream) != 0 {
		panic(errorNotSupportedStream)
	}

	//ignore self-copying
	if cIdx != r.bounds.FromCol || rIdx != r.bounds.FromRow {
		cOffset, rOffset := cIdx-r.bounds.FromCol, rIdx-r.bounds.FromRow

		r.Walk(func(idx, cIdxSource, rIdxSource int, source *Cell) {
			//process only non empty cells
			if !isCellEmpty(source.ml) {
				//ignore target cells with negative indexes
				cIdxTarget, rIdxTarget := cIdxSource+cOffset, rIdxSource+rOffset
				if cIdxTarget >= 0 && rIdxTarget >= 0 {
					target := r.sheet.Cell(cIdxTarget, rIdxTarget)

					//copy data
					*target.ml = *source.ml

					//refresh ref
					target.ml.Ref = types.CellRefFromIndexes(cIdxTarget, rIdxTarget)
				}
			}
		})
	}
}
