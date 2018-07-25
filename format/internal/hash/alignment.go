package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//Alignment returns string with all values of alignment
func Alignment(alignment *ml.CellAlignment) Key {
	if alignment == nil {
		alignment = &ml.CellAlignment{}
	}

	return Key(strings.Join([]string{
		strconv.FormatInt(int64(alignment.Horizontal), 10),
		strconv.FormatInt(int64(alignment.Vertical), 10),
		strconv.FormatInt(int64(alignment.TextRotation), 10),
		strconv.FormatBool(alignment.WrapText),
		strconv.FormatInt(int64(alignment.Indent), 10),
		strconv.FormatInt(int64(alignment.RelativeIndent), 10),
		strconv.FormatBool(alignment.JustifyLastLine),
		strconv.FormatBool(alignment.ShrinkToFit),
		strconv.FormatInt(int64(alignment.ReadingOrder), 10),
	}, ":"))
}
