package rule

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDataBar(t *testing.T) {
	r := New(
		DataBar.Default,
	)

	require.Equal(t, &Info{
		initialized: true,
		validator: DataBar,
		rule: &ml.ConditionalRule{
			Type: primitives.ConditionTypeDataBar,
			DataBar: &ml.DataBar{
				Values: []*ml.ConditionValue{
					{
						Type:  ValueTypeLowest,
						//Value: "1",
					},
					{
						Type:  ValueTypeHighest,
						//Value: "50",
					},
				},
				MinLength: 10,
				MaxLength: 90,
				Color: &ml.Color{
					RGB: "FF638EC6",
				},
				ShowValue: true,
			},
		},
	}, r)

	r = New(
		DataBar.Min("1", ValueTypeLowest),
		DataBar.Max("50", ValueTypeHighest),
		DataBar.Color("#110000"),
		DataBar.BarOnly,
	)

	require.Equal(t, &Info{
		initialized: true,
		validator: DataBar,
		rule: &ml.ConditionalRule{
			Type: primitives.ConditionTypeDataBar,
			DataBar: &ml.DataBar{
				Values: []*ml.ConditionValue{
					{
						Type:  ValueTypeLowest,
						Value: "1",
					},
					{
						Type:  ValueTypeHighest,
						Value: "50",
					},
				},
				ShowValue: false,
				MinLength: 10,
				MaxLength: 90,
				Color: &ml.Color{
					RGB: "FF110000",
				},
			},
		},
	}, r)
}
