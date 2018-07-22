package format

import (
	"github.com/plandem/xlsx/internal/ml"
)

//List of all possible values for PatternType
const (
	_ ml.PatternType = iota
	PatternTypeNone
	PatternTypeSolid
	PatternTypeMediumGray
	PatternTypeDarkGray
	PatternTypeLightGray
	PatternTypeDarkHorizontal
	PatternTypeDarkVertical
	PatternTypeDarkDown
	PatternTypeDarkUp
	PatternTypeDarkGrid
	PatternTypeDarkTrellis
	PatternTypeLightHorizontal
	PatternTypeLightVertical
	PatternTypeLightDown
	PatternTypeLightUp
	PatternTypeLightGrid
	PatternTypeLightTrellis
	PatternTypeGray125
	PatternTypeGray0625
)

func init() {
	ml.FromPatternType = map[ml.PatternType]string{
		PatternTypeNone:            "none",
		PatternTypeSolid:           "solid",
		PatternTypeMediumGray:      "mediumGray",
		PatternTypeDarkGray:        "darkGray",
		PatternTypeLightGray:       "lightGray",
		PatternTypeDarkHorizontal:  "darkHorizontal",
		PatternTypeDarkVertical:    "darkVertical",
		PatternTypeDarkDown:        "darkDown",
		PatternTypeDarkUp:          "darkUp",
		PatternTypeDarkGrid:        "darkGrid",
		PatternTypeDarkTrellis:     "darkTrellis",
		PatternTypeLightHorizontal: "lightHorizontal",
		PatternTypeLightVertical:   "lightVertical",
		PatternTypeLightDown:       "lightDown",
		PatternTypeLightUp:         "lightUp",
		PatternTypeLightGrid:       "lightGrid",
		PatternTypeLightTrellis:    "lightTrellis",
		PatternTypeGray125:         "gray125",
		PatternTypeGray0625:        "gray0625",
	}

	ml.ToPatternType = make(map[string]ml.PatternType, len(ml.FromPatternType))
	for k, v := range ml.FromPatternType {
		ml.ToPatternType[v] = k
	}
}