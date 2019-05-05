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
	)

	require.Equal(t, &conditionalRule{
		rule: &ml.ConditionalRule{
			Formula:      "formula",
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
