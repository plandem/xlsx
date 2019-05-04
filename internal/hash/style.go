package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//NamedStyle returns string with all values of ml.NamedStyle
func NamedStyle(style *ml.NamedStyle) Key {
	if style == nil {
		style = &ml.NamedStyle{}
	}

	return Key(strings.Join([]string{
		strconv.FormatInt(int64(style.NumFmtId), 10),
		strconv.FormatInt(int64(style.FontId), 10),
		strconv.FormatInt(int64(style.FillId), 10),
		strconv.FormatInt(int64(style.BorderId), 10),
		strconv.FormatBool(style.QuotePrefix),
		strconv.FormatBool(style.PivotButton),
		strconv.FormatBool(style.ApplyNumberFormat),
		strconv.FormatBool(style.ApplyFont),
		strconv.FormatBool(style.ApplyFill),
		strconv.FormatBool(style.ApplyBorder),
		strconv.FormatBool(style.ApplyAlignment),
		strconv.FormatBool(style.ApplyProtection),
		string(Alignment(style.Alignment)),
		string(Protection(style.Protection)),
		string(Reserved(style.ExtLst)),
	}, ":"))
}

//DirectStyle returns string with all values of ml.DirectStyle
func DirectStyle(style *ml.DirectStyle) Key {
	if style == nil {
		style = &ml.DirectStyle{}
	}

	s := ml.NamedStyle(style.Style)
	return Key(string(NamedStyle(&s)) + ":" + strconv.FormatInt(int64(style.XfId), 10))
}

//DiffStyle returns string with all values of ml.DiffStyle
func DiffStyle(style *ml.DiffStyle) Key {
	if style == nil {
		style = &ml.DiffStyle{}
	}

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
