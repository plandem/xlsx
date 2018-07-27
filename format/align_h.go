package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
)

//List of all possible values for HAlignType
const (
	_ styles.HAlignType = iota
	HAlignGeneral
	HAlignLeft
	HAlignCenter
	HAlignRight
	HAlignFill
	HAlignJustify
	HAlignCenterContinuous
	HAlignDistributed
)

func init() {
	styles.FromHAlignType = map[styles.HAlignType]string{
		HAlignGeneral:          "general",
		HAlignLeft:             "left",
		HAlignCenter:           "center",
		HAlignRight:            "right",
		HAlignFill:             "fill",
		HAlignJustify:          "justify",
		HAlignCenterContinuous: "centerContinuous",
		HAlignDistributed:      "distributed",
	}

	styles.ToHAlignType = make(map[string]styles.HAlignType, len(styles.FromHAlignType))
	for k, v := range styles.FromHAlignType {
		styles.ToHAlignType[v] = k
	}
}
