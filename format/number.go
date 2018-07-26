package format

import (
	"github.com/plandem/xlsx/internal/number_format"
)

//NumberFormat is option to update StyleFormat with provided custom format of number, but respecting built-in number formats
func NumberFormat(format string) func(*StyleFormat) {
	return func(s *StyleFormat) {
		s.numFormat = numberFormat.New(-1, format)
	}
}

//NumberFormatID is option to update StyleFormat with provided id of already existing or built-in number format
func NumberFormatID(id int) func(*StyleFormat) {
	return func(s *StyleFormat) {
		s.numFormat = numberFormat.New(id, "")
	}
}
