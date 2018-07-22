package format

import (
	"github.com/plandem/xlsx/internal/ml"
)

//List of all possible values for VAlignType
const (
	_ ml.VAlignType = iota
	VAlignTop
	VAlignCenter
	VAlignBottom
	VAlignJustify
	VAlignDistributed
)

func init() {
	ml.FromVAlignType = map[ml.VAlignType]string{
		VAlignTop:         "top",
		VAlignCenter:      "center",
		VAlignBottom:      "bottom",
		VAlignJustify:     "justify",
		VAlignDistributed: "distributed",
	}

	ml.ToVAlignType = make(map[string]ml.VAlignType, len(ml.FromVAlignType))
	for k, v := range ml.FromVAlignType {
		ml.ToVAlignType[v] = k
	}
}