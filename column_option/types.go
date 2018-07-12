package columnOption

import (
	"github.com/plandem/xlsx/format"
)

type Option func(co *Options)

//Options is a helper type to simplify process of settings options for column
type Options struct {
	OutlineLevel uint8
	Collapsed    bool
	Phonetic     bool
	Hidden       bool
	StyleID      format.StyleRefID
	Width        float32
}

//Set sets new options for option set
func (co *Options) Set(options ...Option) {
	for _, o := range options {
		o(co)
	}
}
