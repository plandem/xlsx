package xlsx

import (
	"github.com/plandem/xlsx/column_option"
)

//NewColumnOptions create and returns option set for column
func NewColumnOptions(options ...columnOption.Option) *columnOption.Options {
	s := &columnOption.Options{}
	s.Set(options...)
	return s
}
