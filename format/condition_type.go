package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for ConditionType
const (
	_ primitives.ConditionType = iota
	ConditionTypeExpression
	ConditionTypeCellIs
	ConditionTypeColorScale
	ConditionTypeDataBar
	ConditionTypeIconSet
	ConditionTypeTop10
	ConditionTypeUniqueValues
	ConditionTypeDuplicateValues
	ConditionTypeContainsText
	ConditionTypeNotContainsText
	ConditionTypeBeginsWith
	ConditionTypeEndsWith
	ConditionTypeContainsBlanks
	ConditionTypeNotContainsBlanks
	ConditionTypeContainsErrors
	ConditionTypeNotContainsErrors
	ConditionTypeTimePeriod
	ConditionTypeAboveAverage
)

func init() {
	primitives.FromConditionType = map[primitives.ConditionType]string{
		ConditionTypeExpression:        "expression",
		ConditionTypeCellIs:            "cellIs",
		ConditionTypeColorScale:        "colorScale",
		ConditionTypeDataBar:           "dataBar",
		ConditionTypeIconSet:           "iconSet",
		ConditionTypeTop10:             "top10",
		ConditionTypeUniqueValues:      "uniqueValues",
		ConditionTypeDuplicateValues:   "duplicateValues",
		ConditionTypeContainsText:      "containsText",
		ConditionTypeNotContainsText:   "notContainsText",
		ConditionTypeBeginsWith:        "beginsWith",
		ConditionTypeEndsWith:          "endsWith",
		ConditionTypeContainsBlanks:    "containsBlanks",
		ConditionTypeNotContainsBlanks: "notContainsBlanks",
		ConditionTypeContainsErrors:    "containsErrors",
		ConditionTypeNotContainsErrors: "notContainsErrors",
		ConditionTypeTimePeriod:        "timePeriod",
		ConditionTypeAboveAverage:      "aboveAverage",
	}

	primitives.ToConditionType = make(map[string]primitives.ConditionType, len(primitives.FromConditionType))
	for k, v := range primitives.FromConditionType {
		primitives.ToConditionType[v] = k
	}
}
