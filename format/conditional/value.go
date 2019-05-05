package conditional

import (
	"github.com/plandem/xlsx/internal/ml"
)

type value struct {
	value ml.ConditionValue
}

//ConditionValue returns a conditional value
func Value(t ValueType, v string, gte bool) *value {
	return &value{
		ml.ConditionValue{
			Type:           t,
			Value:          v,
			GreaterOrEqual: gte,
		},
	}
}
