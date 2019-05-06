package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/conditional/rule"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTimePeriod(t *testing.T) {
	type Entity struct {
		Attribute primitives.TimePeriodType `xml:"attribute,attr"`
	}

	list := map[string]primitives.TimePeriodType{
		"":          primitives.TimePeriodType(0),
		"today":     rule.TimePeriodToday,
		"yesterday": rule.TimePeriodYesterday,
		"tomorrow":  rule.TimePeriodTomorrow,
		"last7Days": rule.TimePeriodLast7Days,
		"thisMonth": rule.TimePeriodThisMonth,
		"lastMonth": rule.TimePeriodLastMonth,
		"nextMonth": rule.TimePeriodNextMonth,
		"thisWeek":  rule.TimePeriodThisWeek,
		"lastWeek":  rule.TimePeriodLastWeek,
		"nextWeek":  rule.TimePeriodNextWeek,
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
