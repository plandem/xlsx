package hash

import (
	"crypto/md5"
	"fmt"
	"github.com/plandem/xlsx/internal/ml"
	"io"
	"strings"
)

//Style returns md5 hash string for provided information
func Style(font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, number *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border) string {
	h := md5.New()

	io.WriteString(h, strings.Join([]string{
		Font(font),
		Fill(fill),
		Alignment(alignment),
		NumberFormat(number),
		Protection(protection),
		Border(border),
	}, ":"))

	return fmt.Sprintf("%x", h.Sum(nil))
}
