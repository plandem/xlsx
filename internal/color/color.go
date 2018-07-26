package color

import (
	"github.com/plandem/xlsx/internal/ml"
	"strings"
)

//New create and return ml.Color type for provided value, respecting built-in indexed colors
func New(color string) *ml.Color {
	color = Normalize(color)

	//check if it's indexed color
	for i, c := range indexed {
		if color == c {
			return &ml.Color{Indexed: &i}
		}
	}

	return &ml.Color{RGB: color}
}

//Normalize check if color in #RGB format and convert it into ARGB format
func Normalize(color string) string {
	//normalize color
	if len(color) > 1 {
		if color[0] == '#' {
			if len(color) == 7 {
				color = "FF" + color[1:]
			} else {
				color = color[1:]
			}
		}
	}

	return strings.ToUpper(color)
}
