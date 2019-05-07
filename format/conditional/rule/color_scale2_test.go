package rule

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColorScale2(t *testing.T) {
	r := New(
		ColorScale2.Min("1", "#110000"),
		ColorScale2.Max("10", "#001100"),
	)

	require.Equal(t, &Info{
		initialized: true,
		validator: ColorScale2,
		rule: &ml.ConditionalRule{
			Type: primitives.ConditionTypeColorScale,
			ColorScale: &ml.ColorScale{
				Values: []*ml.ConditionValue{
					{
						Type:  ValueTypeLowest,
						Value: "1",
					},
					{
						Type:  ValueTypeHighest,
						Value: "10",
					},
				},
				Colors: []*ml.Color{
					{
						RGB: "FF110000",
					},
					{
						RGB: "FF001100",
					},
				},
			},
		},
	}, r)
}
