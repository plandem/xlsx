package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for ConditionType
const (
	_ primitives.ConditionType = iota
	ConditionExpression
	ConditionCellIs
	ConditionColorScale
	ConditionDataBar
	ConditionIconSet
	ConditionTop10
	ConditionUniqueValues
	ConditionDuplicateValues
	ConditionContainsText
	ConditionNotContainsText
	ConditionBeginsWith
	ConditionEndsWith
	ConditionContainsBlanks
	ConditionNotContainsBlanks
	ConditionContainsErrors
	ConditionNotContainsErrors
	ConditionTimePeriod
	ConditionAboveAverage
)

func init() {
	primitives.FromConditionType = map[primitives.ConditionType]string{
		ConditionExpression:        "expression",
		ConditionCellIs:            "cellIs",
		ConditionColorScale:        "colorScale",
		ConditionDataBar:           "dataBar",
		ConditionIconSet:           "iconSet",
		ConditionTop10:             "top10",
		ConditionUniqueValues:      "uniqueValues",
		ConditionDuplicateValues:   "duplicateValues",
		ConditionContainsText:      "containsText",
		ConditionNotContainsText:   "notContainsText",
		ConditionBeginsWith:        "beginsWith",
		ConditionEndsWith:          "endsWith",
		ConditionContainsBlanks:    "containsBlanks",
		ConditionNotContainsBlanks: "notContainsBlanks",
		ConditionContainsErrors:    "containsErrors",
		ConditionNotContainsErrors: "notContainsErrors",
		ConditionTimePeriod:        "timePeriod",
		ConditionAboveAverage:      "aboveAverage",
	}

	primitives.ToConditionType = make(map[string]primitives.ConditionType, len(primitives.FromConditionType))
	for k, v := range primitives.FromConditionType {
		primitives.ToConditionType[v] = k
	}
}
