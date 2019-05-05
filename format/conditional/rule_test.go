package conditional

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditionalRule_Set(t *testing.T) {
	s := styles.New(
		styles.Font.Bold,
		styles.Font.Color("#112233"),
	)

	rule := newRule(
		Rule.AboveAverage,
		Rule.StopIfTrue,
		Rule.Percent,
		Rule.Bottom,
		Rule.EqualAverage,
		Rule.Priority(10),
		Rule.Style(s),
		Rule.Type(TypeCellIs),
		Rule.Operator(OperatorBetween),
		Rule.Text("some text"),
		Rule.TimePeriod(TimePeriodLastMonth),
		Rule.Rank(10),
		Rule.Formula("formula"),
		Rule.ColorScale(
			Value(ValueTypePercent, "10", false),
			Value(ValueTypePercent, "50", false),
			Value(ValueTypePercent, "90", true),
			"#112233",
			"#223344",
			"#334455",
		),
		Rule.DataBar(
			Value(ValueTypeMin, "10", false), 10,
			Value(ValueTypeMax, "90", true), 20,
			"#112233", true,
		),
		Rule.IconSet(IconSetType3Arrows, true, true, true,
			Value(ValueTypePercent, "10", false),
			Value(ValueTypePercent, "50", false),
			Value(ValueTypePercent, "90", true),
		),
	)

	require.Equal(t, &ruleInfo{
		rule: &ml.ConditionalRule{
			Formula: "formula",
			ColorScale: &ml.ColorScale{
				Values: []*ml.ConditionValue{
					{
						Type:  ValueTypePercent,
						Value: "10",
					},
					{
						Type:           ValueTypePercent,
						Value:          "50",
						GreaterOrEqual: false,
					},
					{
						Type:           ValueTypePercent,
						Value:          "90",
						GreaterOrEqual: true,
					},
				},
				Colors: []*ml.Color{
					{
						RGB: "FF112233",
					},
					{
						RGB: "FF223344",
					},
					{
						RGB: "FF334455",
					},
				},
			},
			DataBar: &ml.DataBar{
				Values: []*ml.ConditionValue{
					{
						Type:  ValueTypeMin,
						Value: "10",
					},
					{
						Type:           ValueTypeMax,
						Value:          "90",
						GreaterOrEqual: true,
					},
				},
				MinLength: 10,
				MaxLength: 20,
				ShowValue: true,
				Color: &ml.Color{
					RGB: "FF112233",
				},
			},
			IconSet: &ml.IconSet{
				Type:      IconSetType3Arrows,
				Percent:   true,
				ShowValue: true,
				Reverse:   true,
				Values: []*ml.ConditionValue{
					{
						Type:  ValueTypePercent,
						Value: "10",
					},
					{
						Type:           ValueTypePercent,
						Value:          "50",
						GreaterOrEqual: false,
					},
					{
						Type:           ValueTypePercent,
						Value:          "90",
						GreaterOrEqual: true,
					},
				},
			},
			Type:         TypeCellIs,
			Priority:     10,
			StopIfTrue:   true,
			AboveAverage: true,
			Percent:      true,
			Bottom:       true,
			Operator:     OperatorBetween,
			Text:         "some text",
			TimePeriod:   TimePeriodLastMonth,
			Rank:         10,
			EqualAverage: true,
		},
		style: s,
	}, rule)
}
