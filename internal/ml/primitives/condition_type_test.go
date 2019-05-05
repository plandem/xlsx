package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
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
		"expression":        format.ConditionExpression,
		"cellIs":            format.ConditionCellIs,
		"colorScale":        format.ConditionColorScale,
		"dataBar":           format.ConditionDataBar,
		"iconSet":           format.ConditionIconSet,
		"top10":             format.ConditionTop10,
		"uniqueValues":      format.ConditionUniqueValues,
		"duplicateValues":   format.ConditionDuplicateValues,
		"containsText":      format.ConditionContainsText,
		"notContainsText":   format.ConditionNotContainsText,
		"beginsWith":        format.ConditionBeginsWith,
		"endsWith":          format.ConditionEndsWith,
		"containsBlanks":    format.ConditionContainsBlanks,
		"notContainsBlanks": format.ConditionNotContainsBlanks,
		"containsErrors":    format.ConditionContainsErrors,
		"notContainsErrors": format.ConditionNotContainsErrors,
		"timePeriod":        format.ConditionTimePeriod,
		"aboveAverage":      format.ConditionAboveAverage,
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
