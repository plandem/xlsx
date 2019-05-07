package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for ConditionOperatorType
const (
	_ primitives.ConditionOperatorType = iota
	ConditionOperatorLessThan
	ConditionOperatorLessThanOrEqual
	ConditionOperatorEqual
	ConditionOperatorNotEqual
	ConditionOperatorGreaterThanOrEqual
	ConditionOperatorGreaterThan
	ConditionOperatorBetween
	ConditionOperatorNotBetween
	ConditionOperatorContainsText
	ConditionOperatorNotContains
	ConditionOperatorBeginsWith
	ConditionOperatorEndsWith
)

func init() {
	primitives.FromConditionOperatorType = map[primitives.ConditionOperatorType]string{
		ConditionOperatorLessThan:           "lessThan",
		ConditionOperatorLessThanOrEqual:    "lessThanOrEqual",
		ConditionOperatorEqual:              "equal",
		ConditionOperatorNotEqual:           "notEqual",
		ConditionOperatorGreaterThanOrEqual: "greaterThanOrEqual",
		ConditionOperatorGreaterThan:        "greaterThan",
		ConditionOperatorBetween:            "between",
		ConditionOperatorNotBetween:         "notBetween",
		ConditionOperatorContainsText:       "containsText",
		ConditionOperatorNotContains:        "notContains",
		ConditionOperatorBeginsWith:         "beginsWith",
		ConditionOperatorEndsWith:           "endsWith",
	}

	primitives.ToConditionOperatorType = make(map[string]primitives.ConditionOperatorType, len(primitives.FromConditionOperatorType))
	for k, v := range primitives.FromConditionOperatorType {
		primitives.ToConditionOperatorType[v] = k
	}
}
