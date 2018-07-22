package format

import (
	"github.com/plandem/xlsx/internal/ml"
)

//List of all possible values for BorderStyleType
const (
	_ ml.BorderStyleType = iota
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
	ml.FromBorderStyleType = map[ml.BorderStyleType]string{
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

	ml.ToBorderStyleType = make(map[string]ml.BorderStyleType, len(ml.FromBorderStyleType))
	for k, v := range ml.FromBorderStyleType {
		ml.ToBorderStyleType[v] = k
	}
}