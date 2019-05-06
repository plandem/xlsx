package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/conditional/rule"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditionType(t *testing.T) {
	type Entity struct {
		Attribute primitives.ConditionType `xml:"attribute,attr"`
	}

	list := map[string]primitives.ConditionType{
		"":                  primitives.ConditionType(0),
		"expression":        rule.TypeExpression,
		"cellIs":            rule.TypeCellIs,
		"colorScale":        rule.TypeColorScale,
		"dataBar":           rule.TypeDataBar,
		"iconSet":           rule.TypeIconSet,
		"top10":             rule.TypeTop10,
		"uniqueValues":      rule.TypeUniqueValues,
		"duplicateValues":   rule.TypeDuplicateValues,
		"containsText":      rule.TypeContainsText,
		"notContainsText":   rule.TypeNotContainsText,
		"beginsWith":        rule.TypeBeginsWith,
		"endsWith":          rule.TypeEndsWith,
		"containsBlanks":    rule.TypeContainsBlanks,
		"notContainsBlanks": rule.TypeNotContainsBlanks,
		"containsErrors":    rule.TypeContainsErrors,
		"notContainsErrors": rule.TypeNotContainsErrors,
		"timePeriod":        rule.TypeTimePeriod,
		"aboveAverage":      rule.TypeAboveAverage,
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
