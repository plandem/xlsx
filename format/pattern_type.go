package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
)

//List of all possible values for PatternType
const (
	_ styles.PatternType = iota
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
	styles.FromPatternType = map[styles.PatternType]string{
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

	styles.ToPatternType = make(map[string]styles.PatternType, len(styles.FromPatternType))
	for k, v := range styles.FromPatternType {
		styles.ToPatternType[v] = k
	}
}
