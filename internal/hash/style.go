package hash

import (
	"crypto/md5"
	"fmt"
	"github.com/plandem/xlsx/internal/ml"
	"io"
	"strings"
)

type Key string

func (k Key) Hash() string {
	h := md5.New()
	io.WriteString(h, string(k))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//Style returns md5 hash string for provided information
func Style(font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, number *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border) string {
	k := Key(strings.Join([]string{
		string(Font(font)),
		string(Fill(fill)),
		string(Alignment(alignment)),
		string(NumberFormat(number)),
		string(Protection(protection)),
		string(Border(border)),
	}, ":"))

	return k.Hash()
}
