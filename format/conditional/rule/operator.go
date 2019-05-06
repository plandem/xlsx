package rule

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for OperatorType
const (
	_ OperatorType = iota
	OperatorLessThan
	OperatorLessThanOrEqual
	OperatorEqual
	OperatorNotEqual
	OperatorGreaterThanOrEqual
	OperatorGreaterThan
	OperatorBetween
	OperatorNotBetween
	OperatorContainsText
	OperatorNotContains
	OperatorBeginsWith
	OperatorEndsWith
)

func init() {
	primitives.FromConditionOperatorType = map[primitives.ConditionOperatorType]string{
		OperatorLessThan:           "lessThan",
		OperatorLessThanOrEqual:    "lessThanOrEqual",
		OperatorEqual:              "equal",
		OperatorNotEqual:           "notEqual",
		OperatorGreaterThanOrEqual: "greaterThanOrEqual",
		OperatorGreaterThan:        "greaterThan",
		OperatorBetween:            "between",
		OperatorNotBetween:         "notBetween",
		OperatorContainsText:       "containsText",
		OperatorNotContains:        "notContains",
		OperatorBeginsWith:         "beginsWith",
		OperatorEndsWith:           "endsWith",
	}

	primitives.ToConditionOperatorType = make(map[string]primitives.ConditionOperatorType, len(primitives.FromConditionOperatorType))
	for k, v := range primitives.FromConditionOperatorType {
		primitives.ToConditionOperatorType[v] = k
	}
}
