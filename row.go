// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types/options/row"
)

//Row is a higher level object that wraps ml.Row with functionality. Inherits functionality of Range
type Row struct {
	ml *ml.Row
	*Range
}

//Cell returns cell of row at col with colIndex
func (r *Row) Cell(colIndex int) *Cell {
	return r.sheet.Cell(colIndex, r.bounds.FromRow)
}

//SetOptions sets options for row
func (r *Row) SetOptions(o *options.Info) {
	if o.Height > 0 {
		r.ml.Height = o.Height
		r.ml.CustomHeight = true
	}

	if o.Format != nil {
		r.SetStyles(o.Format)
	}

	r.ml.OutlineLevel = o.OutlineLevel
	r.ml.Hidden = o.Hidden
	r.ml.Collapsed = o.Collapsed
	r.ml.Phonetic = o.Phonetic
}

//Styles returns DirectStyleID of default format for row
func (r *Row) Styles() styles.DirectStyleID {
	return r.ml.Style
}

//SetStyles sets default style for the row. Affects cells not yet allocated in the row. In other words, this style applies to new cells.
func (r *Row) SetStyles(s interface{}) {
	r.ml.CustomFormat = true
	r.ml.Style = r.sheet.info().resolveStyleID(s)
}

//CopyTo copies row cells into another row with rIdx index.
//N.B.: Merged cells are not supported
func (r *Row) CopyTo(rIdx int, withOptions bool) {
	if withOptions {
		//TODO: copy row options
		panic(errorNotSupported)
	}

	//copy cell data
	r.Range.CopyTo(r.Range.bounds.FromCol, rIdx)
}
