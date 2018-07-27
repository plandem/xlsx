package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
)

//List of all possible values for VAlignType
const (
	_ styles.VAlignType = iota
	VAlignTop
	VAlignCenter
	VAlignBottom
	VAlignJustify
	VAlignDistributed
)

func init() {
	styles.FromVAlignType = map[styles.VAlignType]string{
		VAlignTop:         "top",
		VAlignCenter:      "center",
		VAlignBottom:      "bottom",
		VAlignJustify:     "justify",
		VAlignDistributed: "distributed",
	}

	styles.ToVAlignType = make(map[string]styles.VAlignType, len(styles.FromVAlignType))
	for k, v := range styles.FromVAlignType {
		styles.ToVAlignType[v] = k
	}
}
