package xlsx

import (
	"github.com/plandem/xlsx/column_option"
)

//NewColumnOptions create and returns option set for column
func NewColumnOptions(options ...columnOption.ColumnOption) *columnOption.ColumnOptions {
	s := &columnOption.ColumnOptions{}
	s.Set(options...)
	return s
}
