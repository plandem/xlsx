package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strings"
)

//DiffStyle returns string with all values of ml.DiffStyle
func DiffStyle(style *ml.DiffStyle) Key {
	return Key(strings.Join([]string{
		string(Border(style.Border)),
		string(Fill(style.Fill)),
		string(Font(style.Font)),
		string(NumberFormat(style.NumberFormat)),
		string(Alignment(style.Alignment)),
		string(Protection(style.Protection)),
		string(Reserved(style.ExtLst)),
	}, ":"))
}
