package xlsx

import (
	"github.com/plandem/xlsx/row_option"
)

//NewRowOptions create and returns option set for row
func NewRowOptions(options ...rowOption.Option) *rowOption.Options {
	s := &rowOption.Options{}
	s.Set(options...)
	return s
}
