package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
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
		"lessThan":           format.ConditionOperatorLessThan,
		"lessThanOrEqual":    format.ConditionOperatorLessThanOrEqual,
		"equal":              format.ConditionOperatorEqual,
		"notEqual":           format.ConditionOperatorNotEqual,
		"greaterThanOrEqual": format.ConditionOperatorGreaterThanOrEqual,
		"greaterThan":        format.ConditionOperatorGreaterThan,
		"between":            format.ConditionOperatorBetween,
		"notBetween":         format.ConditionOperatorNotBetween,
		"containsText":       format.ConditionOperatorContainsText,
		"notContains":        format.ConditionOperatorNotContains,
		"beginsWith":         format.ConditionOperatorBeginsWith,
		"endsWith":           format.ConditionOperatorEndsWith,
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
