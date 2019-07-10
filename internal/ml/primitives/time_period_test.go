// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

	list := map[primitives.TimePeriodType]string{
		primitives.TimePeriodType(0):   "",
		primitives.TimePeriodToday:     primitives.TimePeriodToday.String(),
		primitives.TimePeriodYesterday: primitives.TimePeriodYesterday.String(),
		primitives.TimePeriodTomorrow:  primitives.TimePeriodTomorrow.String(),
		primitives.TimePeriodLast7Days: primitives.TimePeriodLast7Days.String(),
		primitives.TimePeriodThisMonth: primitives.TimePeriodThisMonth.String(),
		primitives.TimePeriodLastMonth: primitives.TimePeriodLastMonth.String(),
		primitives.TimePeriodNextMonth: primitives.TimePeriodNextMonth.String(),
		primitives.TimePeriodThisWeek:  primitives.TimePeriodThisWeek.String(),
		primitives.TimePeriodLastWeek:  primitives.TimePeriodLastWeek.String(),
		primitives.TimePeriodNextWeek:  primitives.TimePeriodNextWeek.String(),
	}

	for v, s := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if v == 0 {
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
