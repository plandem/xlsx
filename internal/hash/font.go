package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//Font return string with all values of font
func Font(font *ml.Font) Key {
	if font == nil {
		font = &ml.Font{}
	}

	return Key(strings.Join([]string{
		string(font.Name),
		strconv.FormatInt(int64(font.Charset), 10),
		strconv.FormatInt(int64(font.Family), 10),
		strconv.FormatBool(bool(font.Bold)),
		strconv.FormatBool(bool(font.Italic)),
		strconv.FormatBool(bool(font.Strike)),
		strconv.FormatBool(bool(font.Shadow)),
		strconv.FormatBool(bool(font.Condense)),
		strconv.FormatBool(bool(font.Extend)),
		string(Color(font.Color)),
		strconv.FormatFloat(float64(font.Size), 'f', -1, 64),
		string(font.Underline),
		string(font.VAlign),
		string(font.Scheme),
	}, ":"))
}
