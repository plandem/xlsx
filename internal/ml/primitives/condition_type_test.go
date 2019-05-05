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
		"expression":        format.ConditionTypeExpression,
		"cellIs":            format.ConditionTypeCellIs,
		"colorScale":        format.ConditionTypeColorScale,
		"dataBar":           format.ConditionTypeDataBar,
		"iconSet":           format.ConditionTypeIconSet,
		"top10":             format.ConditionTypeTop10,
		"uniqueValues":      format.ConditionTypeUniqueValues,
		"duplicateValues":   format.ConditionTypeDuplicateValues,
		"containsText":      format.ConditionTypeContainsText,
		"notContainsText":   format.ConditionTypeNotContainsText,
		"beginsWith":        format.ConditionTypeBeginsWith,
		"endsWith":          format.ConditionTypeEndsWith,
		"containsBlanks":    format.ConditionTypeContainsBlanks,
		"notContainsBlanks": format.ConditionTypeNotContainsBlanks,
		"containsErrors":    format.ConditionTypeContainsErrors,
		"notContainsErrors": format.ConditionTypeNotContainsErrors,
		"timePeriod":        format.ConditionTypeTimePeriod,
		"aboveAverage":      format.ConditionTypeAboveAverage,
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
