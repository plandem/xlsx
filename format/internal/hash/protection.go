package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//Protection return string with all values of protection
func Protection(protection *ml.CellProtection) string {
	if protection == nil {
		protection = &ml.CellProtection{}
	}

	return strings.Join([]string{
		strconv.FormatBool(protection.Locked),
		strconv.FormatBool(protection.Hidden),
	}, ":")
}
