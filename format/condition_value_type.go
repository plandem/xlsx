package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for ConditionValueType
const (
	_ primitives.ConditionValueType = iota
	ConditionValueTypeNum
	ConditionValueTypePercent
	ConditionValueTypeMax
	ConditionValueTypeMin
	ConditionValueTypeFormula
	ConditionValueTypePercentile
)

func init() {
	primitives.FromConditionValueType = map[primitives.ConditionValueType]string{
		ConditionValueTypeNum:        "num",
		ConditionValueTypePercent:    "percent",
		ConditionValueTypeMax:        "max",
		ConditionValueTypeMin:        "min",
		ConditionValueTypeFormula:    "formula",
		ConditionValueTypePercentile: "percentile",
	}

	primitives.ToConditionValueType = make(map[string]primitives.ConditionValueType, len(primitives.FromConditionValueType))
	for k, v := range primitives.FromConditionValueType {
		primitives.ToConditionValueType[v] = k
	}
}
