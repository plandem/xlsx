package rowOption

import (
	"github.com/plandem/xlsx/format"
)

type RowOption func(co *RowOptions)

//RowOptions is a helper type to simplify process of settings options for row
type RowOptions struct {
	OutlineLevel uint8
	Collapsed    bool
	Phonetic     bool
	Hidden       bool
	StyleID      format.StyleRefID
	Height       float32
}

//Set sets new options for option set
func (ro *RowOptions) Set(options ...RowOption) {
	for _, o := range options {
		o(ro)
	}
}
