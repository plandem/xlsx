package format

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditionalRule_Set(t *testing.T) {
	rule := newConditionalRule(
		Condition.AboveAverage,
		Condition.StopIfTrue,
		Condition.Percent,
		Condition.Bottom,
		Condition.EqualAverage,
		Condition.Priority(10),
		Condition.Style(NewStyles(
			Font.Bold,
			Font.Color("#112233"),
		)),
		Condition.Type(ConditionTypeCellIs),
		Condition.Operator(ConditionOperatorBetween),
		Condition.Text("some text"),
		Condition.TimePeriod(TimePeriodLastMonth),
		Condition.Rank(10),
		Condition.Formula("formula"),
		Condition.ColorScale(
			ConditionValue(ConditionValueTypePercent, "10", false),
			ConditionValue(ConditionValueTypePercent, "50", false),
			ConditionValue(ConditionValueTypePercent, "90", true),
			"#112233",
			"#223344",
			"#334455",
		),
		Condition.DataBar(
			ConditionValue(ConditionValueTypeMin, "10", false), 10,
			ConditionValue(ConditionValueTypeMax, "90", true), 20,
			"#112233", true,
		),
		Condition.IconSet(IconSetType3Arrows, true, true, true,
			ConditionValue(ConditionValueTypePercent, "10", false),
			ConditionValue(ConditionValueTypePercent, "50", false),
			ConditionValue(ConditionValueTypePercent, "90", true),
		),
	)

	require.Equal(t, &conditionalRule{
		rule: &ml.ConditionalRule{
			Formula:      "formula",
			ColorScale: &ml.ColorScale{
				Values: []*ml.ConditionValue{
					{
						Type: ConditionValueTypePercent,
						Value: "10",
					},
					{
						Type: ConditionValueTypePercent,
						Value: "50",
						GreaterOrEqual: false,
					},
					{
						Type: ConditionValueTypePercent,
						Value: "90",
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
						Type: ConditionValueTypeMin,
						Value: "10",
					},
					{
						Type: ConditionValueTypeMax,
						Value: "90",
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
				Type: IconSetType3Arrows,
				Percent: true,
				ShowValue: true,
				Reverse: true,
				Values: []*ml.ConditionValue{
					{
						Type: ConditionValueTypePercent,
						Value: "10",
					},
					{
						Type: ConditionValueTypePercent,
						Value: "50",
						GreaterOrEqual: false,
					},
					{
						Type: ConditionValueTypePercent,
						Value: "90",
						GreaterOrEqual: true,
					},
				},
			},
			Type:         ConditionTypeCellIs,
			Priority:     10,
			StopIfTrue:   true,
			AboveAverage: true,
			Percent:      true,
			Bottom:       true,
			Operator:     ConditionOperatorBetween,
			Text:         "some text",
			TimePeriod:   TimePeriodLastMonth,
			Rank:         10,
			EqualAverage: true,
		},
		style: &StyleFormat{
			&ml.DiffStyle{
				Font: &ml.Font{
					Bold:  true,
					Color: &ml.Color{RGB: "FF112233"},
				},
				NumberFormat: &ml.NumberFormat{},
				Fill: &ml.Fill{
					Pattern:  &ml.PatternFill{},
					Gradient: &ml.GradientFill{},
				},
				Border: &ml.Border{
					Left:       &ml.BorderSegment{},
					Right:      &ml.BorderSegment{},
					Top:        &ml.BorderSegment{},
					Bottom:     &ml.BorderSegment{},
					Diagonal:   &ml.BorderSegment{},
					Vertical:   &ml.BorderSegment{},
					Horizontal: &ml.BorderSegment{},
				},
				Alignment:  &ml.CellAlignment{},
				Protection: &ml.CellProtection{},
			},
			&ml.NamedStyleInfo{},
		},
	}, rule)
}
