package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for VAlignType
const (
	_ primitives.VAlignType = iota
	VAlignTop
	VAlignCenter
	VAlignBottom
	VAlignJustify
	VAlignDistributed
)

func init() {
	primitives.FromVAlignType = map[primitives.VAlignType]string{
		VAlignTop:         "top",
		VAlignCenter:      "center",
		VAlignBottom:      "bottom",
		VAlignJustify:     "justify",
		VAlignDistributed: "distributed",
	}

	primitives.ToVAlignType = make(map[string]primitives.VAlignType, len(primitives.FromVAlignType))
	for k, v := range primitives.FromVAlignType {
		primitives.ToVAlignType[v] = k
	}
}
