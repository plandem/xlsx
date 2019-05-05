package conditional

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for Type
const (
	_ Type = iota
	TypeExpression
	TypeCellIs
	TypeColorScale
	TypeDataBar
	TypeIconSet
	TypeTop10
	TypeUniqueValues
	TypeDuplicateValues
	TypeContainsText
	TypeNotContainsText
	TypeBeginsWith
	TypeEndsWith
	TypeContainsBlanks
	TypeNotContainsBlanks
	TypeContainsErrors
	TypeNotContainsErrors
	TypeTimePeriod
	TypeAboveAverage
)

func init() {
	primitives.FromConditionType = map[Type]string{
		TypeExpression:        "expression",
		TypeCellIs:            "cellIs",
		TypeColorScale:        "colorScale",
		TypeDataBar:           "dataBar",
		TypeIconSet:           "iconSet",
		TypeTop10:             "top10",
		TypeUniqueValues:      "uniqueValues",
		TypeDuplicateValues:   "duplicateValues",
		TypeContainsText:      "containsText",
		TypeNotContainsText:   "notContainsText",
		TypeBeginsWith:        "beginsWith",
		TypeEndsWith:          "endsWith",
		TypeContainsBlanks:    "containsBlanks",
		TypeNotContainsBlanks: "notContainsBlanks",
		TypeContainsErrors:    "containsErrors",
		TypeNotContainsErrors: "notContainsErrors",
		TypeTimePeriod:        "timePeriod",
		TypeAboveAverage:      "aboveAverage",
	}

	primitives.ToConditionType = make(map[string]Type, len(primitives.FromConditionType))
	for k, v := range primitives.FromConditionType {
		primitives.ToConditionType[v] = k
	}
}
