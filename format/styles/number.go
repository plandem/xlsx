package styles

import (
	"github.com/plandem/xlsx/internal/number_format"
)

//NumberFormat is option to update Info with provided custom format of number, but respecting built-in number formats
func NumberFormat(format string) func(*Info) {
	return func(s *Info) {
		*s.styleInfo.NumberFormat = numberFormat.New(-1, format)
	}
}

//NumberFormatID is option to update Info with provided id of already existing or built-in number format
func NumberFormatID(id int) func(*Info) {
	return func(s *Info) {
		*s.styleInfo.NumberFormat = numberFormat.New(id, "")
	}
}
