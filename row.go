package xlsx

import (
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
)

//Row is a higher level object that wraps ml.Row with functionality
type Row struct {
	ml    *ml.Row
	sheet *Sheet
}

//walk walks for row range and call callback for each cell
func (r *Row) walk(cb rangeWalkerFn) {
	for _, c := range r.ml.Cells {
		if c != nil {
			cb(c)
		}
	}
}

//Reset resets each cell data into zero state
func (r *Row) Reset() {
	r.walk(rangeResetFn)
}

//Clear clears each cell value
func (r *Row) Clear() {
	r.walk(rangeClearFn)
}

//Cells returns all cells in row
func (r *Row) Cells() []*Cell {
	r.sheet.resolveDimension(false)
	_, toRef := r.sheet.ml.Dimension.Ref.ToCellRefs()
	maxWidth, _ := toRef.ToIndexes()

	cells := make([]*Cell, maxWidth + 1)
	for cid := 0; cid <= maxWidth; cid++ {
		cells[cid] = r.sheet.Cell(cid, r.ml.Ref-1)
	}

	return cells
}

//Values returns values for all cells in row
func (r *Row) Values() []string {
	cells := r.Cells()
	values := make([]string, len(cells))
	for i, cell := range cells {
		values[i] = cell.Value()
	}

	return values
}

//Set sets options for row
func (r *Row) Set(options *internal.RowOptions) {
	if options.Height > 0 {
		r.ml.Height = options.Height
		r.ml.CustomHeight = true
	}

	if options.StyleID > 0 {
		r.ml.CustomFormat = true
		r.ml.Style = options.StyleID
	}

	r.ml.OutlineLevel = options.OutlineLevel
	r.ml.Hidden = options.Hidden
	r.ml.Collapsed = options.Collapsed
	r.ml.Phonetic = options.Phonetic
}
