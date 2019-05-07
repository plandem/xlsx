package format

import (
	"github.com/plandem/xlsx/internal/ml"
)

type conditionValue struct {
	value ml.ConditionValue
}

//ConditionValue returns a conditional value
func ConditionValue(t ConditionValueType, value string, gte bool) *conditionValue {
	return &conditionValue{
		ml.ConditionValue{
			Type:           t,
			Value:          value,
			GreaterOrEqual: gte,
		},
	}
}
