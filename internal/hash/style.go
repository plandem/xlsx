package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//Style returns string with all values of ml.Style
func Style(style *ml.Style) Key {
	return Key(strings.Join([]string{
		strconv.FormatInt(int64(style.NumFmtId), 10),
		strconv.FormatInt(int64(style.FontId), 10),
		strconv.FormatInt(int64(style.FillId), 10),
		strconv.FormatInt(int64(style.BorderId), 10),
		strconv.FormatInt(int64(style.XfId), 10),
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
