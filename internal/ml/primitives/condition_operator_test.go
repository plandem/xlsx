package primitives_test

import (
	"encoding/xml"
	"fmt"
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
		"lessThan":           primitives.ConditionOperatorLessThan,
		"lessThanOrEqual":    primitives.ConditionOperatorLessThanOrEqual,
		"equal":              primitives.ConditionOperatorEqual,
		"notEqual":           primitives.ConditionOperatorNotEqual,
		"greaterThanOrEqual": primitives.ConditionOperatorGreaterThanOrEqual,
		"greaterThan":        primitives.ConditionOperatorGreaterThan,
		"between":            primitives.ConditionOperatorBetween,
		"notBetween":         primitives.ConditionOperatorNotBetween,
		"containsText":       primitives.ConditionOperatorContainsText,
		"notContains":        primitives.ConditionOperatorNotContains,
		"beginsWith":         primitives.ConditionOperatorBeginsWith,
		"endsWith":           primitives.ConditionOperatorEndsWith,
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
