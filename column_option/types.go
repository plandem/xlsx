package columnOption

import (
	"github.com/plandem/xlsx/format"
)

type ColumnOption func(co *ColumnOptions)

//ColumnOptions is a helper type to simplify process of settings options for column
type ColumnOptions struct {
	OutlineLevel uint8
	Collapsed    bool
	Phonetic     bool
	Hidden       bool
	StyleID      format.StyleRefID
	Width        float32
}

//Set sets new options for option set
func (co *ColumnOptions) Set(options ...ColumnOption) {
	for _, o := range options {
		o(co)
	}
}
