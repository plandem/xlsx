package format

import (
	"github.com/plandem/xlsx/internal/ml"
)

//List of all possible values for HAlignType
const (
	_ ml.HAlignType = iota
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
	ml.FromHAlignType = map[ml.HAlignType]string{
		HAlignGeneral:          "general",
		HAlignLeft:             "left",
		HAlignCenter:           "center",
		HAlignRight:            "right",
		HAlignFill:             "fill",
		HAlignJustify:          "justify",
		HAlignCenterContinuous: "centerContinuous",
		HAlignDistributed:      "distributed",
	}

	ml.ToHAlignType = make(map[string]ml.HAlignType, len(ml.FromHAlignType))
	for k, v := range ml.FromHAlignType {
		ml.ToHAlignType[v] = k
	}
}
