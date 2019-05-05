package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/conditional"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditionOperator(t *testing.T) {
	type Entity struct {
		Attribute primitives.ConditionOperatorType `xml:"attribute,attr"`
	}

	list := map[string]primitives.ConditionOperatorType{
		"":                   primitives.ConditionOperatorType(0),
		"lessThan":           conditional.OperatorLessThan,
		"lessThanOrEqual":    conditional.OperatorLessThanOrEqual,
		"equal":              conditional.OperatorEqual,
		"notEqual":           conditional.OperatorNotEqual,
		"greaterThanOrEqual": conditional.OperatorGreaterThanOrEqual,
		"greaterThan":        conditional.OperatorGreaterThan,
		"between":            conditional.OperatorBetween,
		"notBetween":         conditional.OperatorNotBetween,
		"containsText":       conditional.OperatorContainsText,
		"notContains":        conditional.OperatorNotContains,
		"beginsWith":         conditional.OperatorBeginsWith,
		"endsWith":           conditional.OperatorEndsWith,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if s == "" {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
			require.Equal(tt, s, decoded.Attribute.String())
		})
	}
}
