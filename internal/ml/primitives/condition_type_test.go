package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/conditional"
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
		"expression":        conditional.TypeExpression,
		"cellIs":            conditional.TypeCellIs,
		"colorScale":        conditional.TypeColorScale,
		"dataBar":           conditional.TypeDataBar,
		"iconSet":           conditional.TypeIconSet,
		"top10":             conditional.TypeTop10,
		"uniqueValues":      conditional.TypeUniqueValues,
		"duplicateValues":   conditional.TypeDuplicateValues,
		"containsText":      conditional.TypeContainsText,
		"notContainsText":   conditional.TypeNotContainsText,
		"beginsWith":        conditional.TypeBeginsWith,
		"endsWith":          conditional.TypeEndsWith,
		"containsBlanks":    conditional.TypeContainsBlanks,
		"notContainsBlanks": conditional.TypeNotContainsBlanks,
		"containsErrors":    conditional.TypeContainsErrors,
		"notContainsErrors": conditional.TypeNotContainsErrors,
		"timePeriod":        conditional.TypeTimePeriod,
		"aboveAverage":      conditional.TypeAboveAverage,
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
