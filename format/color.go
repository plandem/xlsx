package format

import (
	"strings"
	//"github.com/plandem/xlsx/internal/ml"
)

func newColor(color string) string {
	return color
}
//func newColor(color string) *ml.Color {
	//color = normalizeColor(color)
	//
	////check if it's indexed color
	//for i, c := range indexedColors {
	//	if color == c {
	//		return &ml.Color{ Indexed: &i }
	//	}
	//}
	//
	//return &ml.Color{ RGB: color }
//}

//normalizeColor check if color in #RGB format and convert it into ARGB format
func normalizeColor(color string) string {
	//normalize color
	if len(color) > 1 {
		if color[0] == '#' {
			color = "FF" + color[1:]
		}
	}

	return strings.ToUpper(color)
}
