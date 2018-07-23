package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
)

//List of all possible values for BorderStyleType
const (
	_ styles.BorderStyleType = iota
	BorderStyleNone
	BorderStyleThin
	BorderStyleMedium
	BorderStyleDashed
	BorderStyleDotted
	BorderStyleThick
	BorderStyleDouble
	BorderStyleHair
	BorderStyleMediumDashed
	BorderStyleDashDot
	BorderStyleMediumDashDot
	BorderStyleDashDotDot
	BorderStyleMediumDashDotDot
	BorderStyleSlantDashDot
)

func init() {
	styles.FromBorderStyleType = map[styles.BorderStyleType]string{
		BorderStyleNone:             "none",
		BorderStyleThin:             "thin",
		BorderStyleMedium:           "medium",
		BorderStyleDashed:           "dashed",
		BorderStyleDotted:           "dotted",
		BorderStyleThick:            "thick",
		BorderStyleDouble:           "double",
		BorderStyleHair:             "hair",
		BorderStyleMediumDashed:     "mediumDashed",
		BorderStyleDashDot:          "dashDot",
		BorderStyleMediumDashDot:    "mediumDashDot",
		BorderStyleDashDotDot:       "dashDotDot",
		BorderStyleMediumDashDotDot: "mediumDashDotDot",
		BorderStyleSlantDashDot:     "slantDashDot",
	}

	styles.ToBorderStyleType = make(map[string]styles.BorderStyleType, len(styles.FromBorderStyleType))
	for k, v := range styles.FromBorderStyleType {
		styles.ToBorderStyleType[v] = k
	}
}