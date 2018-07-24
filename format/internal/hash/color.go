package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//Color returns string with all values of color
func Color(color *ml.Color) string {
	if color == nil {
		color = &ml.Color{}
	}

	result := []string{
		strconv.FormatBool(color.Auto),
		color.RGB,
		strconv.FormatFloat(color.Tint, 'f', -1, 64),
	}

	if color.Indexed != nil {
		result = append(result, strconv.FormatInt(int64(*color.Indexed), 10))
	} else {
		result = append(result, "")
	}

	if color.Theme != nil {
		result = append(result, strconv.FormatInt(int64(*color.Theme), 10))
	} else {
		result = append(result, "")
	}

	return strings.Join(result, ":")
}
