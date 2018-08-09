package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for BorderStyleType
const (
	_ primitives.BorderStyleType = iota
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
	primitives.FromBorderStyleType = map[primitives.BorderStyleType]string{
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

	primitives.ToBorderStyleType = make(map[string]primitives.BorderStyleType, len(primitives.FromBorderStyleType))
	for k, v := range primitives.FromBorderStyleType {
		primitives.ToBorderStyleType[v] = k
	}
}
