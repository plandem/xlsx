package format

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditionalFormat_Set(t *testing.T) {
	conditions := NewConditions(
		Conditions.Pivot,
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Priority(10),
			Condition.Type(ConditionTypeCellIs),
		),
		Conditions.Rule(
			Condition.Priority(90),
			Condition.Type(ConditionTypeAboveAverage),
			Condition.Style(NewStyles(
				Font.Bold,
				Font.Color("#112233"),
			)),
		),
	)

	require.Equal(t, &ConditionalFormat{
		info: &ml.ConditionalFormatting{
			Pivot:  true,
			Bounds: primitives.BoundsListFromRefs("A10:B20"),
		},
		rules: []*conditionalRule{
			{
				rule: &ml.ConditionalRule{
					Type:     ConditionTypeCellIs,
					Priority: 10,
				},
			},
			{
				rule: &ml.ConditionalRule{
					Type:     ConditionTypeAboveAverage,
					Priority: 90,
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
			},
		},
	}, conditions)
}

func TestConditionalFormat_Validate(t *testing.T) {
	require.NotNil(t, NewConditions().Validate())

	require.NotNil(t, NewConditions(Conditions.Refs("A10:B20")).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Priority(-1),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeCellIs),
		),
	).Validate())

	require.Nil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeCellIs),
			Condition.Priority(1),
			Condition.Operator(ConditionOperatorLessThanOrEqual),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeTop10),
			Condition.Priority(1),
			//Condition.Rank(),
		),
	).Validate())

	require.Nil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeTop10),
			Condition.Priority(1),
			Condition.Rank(1),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeContainsText),
			Condition.Priority(1),
		),
	).Validate())

	require.Nil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeContainsText),
			Condition.Priority(1),
			Condition.Text("abc"),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeTimePeriod),
			Condition.Priority(1),
		),
	).Validate())

	require.Nil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeTimePeriod),
			Condition.Priority(1),
			Condition.TimePeriod(TimePeriodLastMonth),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeAboveAverage),
			Condition.Priority(1),
			Condition.ColorScale(),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeAboveAverage),
			Condition.Priority(1),
			Condition.ColorScale(
				ConditionValue(ConditionValueTypePercent, "10", false),
				"#112233",
			),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeAboveAverage),
			Condition.Priority(1),
			Condition.ColorScale(
				ConditionValue(ConditionValueTypePercent, "10", false),
				ConditionValue(ConditionValueTypePercent, "50", false),
				"#112233",
			),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeAboveAverage),
			Condition.Priority(1),
			Condition.ColorScale(
				ConditionValue(ConditionValueTypePercent, "50", false),
				"#112233",
				"#334455",
			),
		),
	).Validate())

	require.Nil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeAboveAverage),
			Condition.Priority(1),
			Condition.ColorScale(
				ConditionValue(ConditionValueTypePercent, "10", false),
				ConditionValue(ConditionValueTypePercent, "50", false),
				"#112233",
				"#334455",
			),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeAboveAverage),
			Condition.Priority(1),
			Condition.IconSet(IconSetType3Arrows, true, true, true),
		),
	).Validate())

	require.NotNil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeAboveAverage),
			Condition.Priority(1),
			Condition.IconSet(IconSetType3Arrows, true, true, true,
				ConditionValue(ConditionValueTypePercent, "10", false),
			),
		),
	).Validate())

	require.Nil(t, NewConditions(
		Conditions.Refs("A10:B20"),
		Conditions.Rule(
			Condition.Type(ConditionTypeAboveAverage),
			Condition.Priority(1),
			Condition.IconSet(IconSetType3Arrows, true, true, true,
				ConditionValue(ConditionValueTypePercent, "10", false),
				ConditionValue(ConditionValueTypePercent, "50", false),
			),
		),
	).Validate())
}
