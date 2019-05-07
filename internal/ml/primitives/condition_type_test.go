package primitives_test

import (
	"encoding/xml"
	"fmt"
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
		"expression":        primitives.ConditionTypeExpression,
		"cellIs":            primitives.ConditionTypeCellIs,
		"colorScale":        primitives.ConditionTypeColorScale,
		"dataBar":           primitives.ConditionTypeDataBar,
		"iconSet":           primitives.ConditionTypeIconSet,
		"top10":             primitives.ConditionTypeTop10,
		"uniqueValues":      primitives.ConditionTypeUniqueValues,
		"duplicateValues":   primitives.ConditionTypeDuplicateValues,
		"containsText":      primitives.ConditionTypeContainsText,
		"notContainsText":   primitives.ConditionTypeNotContainsText,
		"beginsWith":        primitives.ConditionTypeBeginsWith,
		"endsWith":          primitives.ConditionTypeEndsWith,
		"containsBlanks":    primitives.ConditionTypeContainsBlanks,
		"notContainsBlanks": primitives.ConditionTypeNotContainsBlanks,
		"containsErrors":    primitives.ConditionTypeContainsErrors,
		"notContainsErrors": primitives.ConditionTypeNotContainsErrors,
		"timePeriod":        primitives.ConditionTypeTimePeriod,
		"aboveAverage":      primitives.ConditionTypeAboveAverage,
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
