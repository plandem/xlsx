package xlsx

import (
	"github.com/plandem/xlsx/row_option"
)

//NewRowOptions create and returns option set for row
func NewRowOptions(options ...rowOption.RowOption) *rowOption.RowOptions {
	s := &rowOption.RowOptions{}
	s.Set(options...)
	return s
}
