package primitives_test

import (
	"encoding/xml"
	"fmt"

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
		"today":     primitives.TimePeriodToday,
		"yesterday": primitives.TimePeriodYesterday,
		"tomorrow":  primitives.TimePeriodTomorrow,
		"last7Days": primitives.TimePeriodLast7Days,
		"thisMonth": primitives.TimePeriodThisMonth,
		"lastMonth": primitives.TimePeriodLastMonth,
		"nextMonth": primitives.TimePeriodNextMonth,
		"thisWeek":  primitives.TimePeriodThisWeek,
		"lastWeek":  primitives.TimePeriodLastWeek,
		"nextWeek":  primitives.TimePeriodNextWeek,
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
