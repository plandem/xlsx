package rule

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColorScale3(t *testing.T) {
	r := New(
		ColorScale3.Min("1", "#110000"),
		ColorScale3.Mid("50", "#111100"),
		ColorScale3.Max("10", "#001100"),
	)

	require.Equal(t, &Info{
		initialized: true,
		validator:   ColorScale3,
		rule: &ml.ConditionalRule{
			Type: primitives.ConditionTypeColorScale,
			ColorScale: &ml.ColorScale{
				Values: []*ml.ConditionValue{
					{
						Type:  ValueTypeLowest,
						Value: "1",
					},
					{
						Type:  ValueTypePercentile,
						Value: "50",
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
						RGB: "FF111100",
					},
					{
						RGB: "FF001100",
					},
				},
			},
		},
	}, r)
}
