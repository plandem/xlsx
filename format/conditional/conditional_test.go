package conditional

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditionalFormat_Set(t *testing.T) {
	s := styles.New(
		styles.Font.Bold,
		styles.Font.Color("#112233"),
	)

	conditions := New(
		Pivot,
		Refs("A10:B20"),
		AddRule(
			Rule.Priority(10),
			Rule.Type(TypeCellIs),
		),
		AddRule(
			Rule.Priority(90),
			Rule.Type(TypeAboveAverage),
			Rule.Style(s),
		),
	)

	require.Equal(t, &Info{
		info: &ml.ConditionalFormatting{
			Pivot:  true,
			Bounds: primitives.BoundsListFromRefs("A10:B20"),
		},
		rules: []*ruleInfo{
			{
				rule: &ml.ConditionalRule{
					Type:     TypeCellIs,
					Priority: 10,
				},
			},
			{
				rule: &ml.ConditionalRule{
					Type:     TypeAboveAverage,
					Priority: 90,
				},
				style: s,
			},
		},
	}, conditions)
}

func TestConditionalFormat_Validate(t *testing.T) {
	require.NotNil(t, New().Validate())

	require.NotNil(t, New(Refs("A10:B20")).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Priority(-1),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeCellIs),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeCellIs),
			Rule.Priority(1),
			Rule.Operator(OperatorLessThanOrEqual),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeTop10),
			Rule.Priority(1),
			//Rule.Rank(),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeTop10),
			Rule.Priority(1),
			Rule.Rank(1),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeContainsText),
			Rule.Priority(1),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeContainsText),
			Rule.Priority(1),
			Rule.Text("abc"),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeTimePeriod),
			Rule.Priority(1),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeTimePeriod),
			Rule.Priority(1),
			Rule.TimePeriod(TimePeriodLastMonth),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeAboveAverage),
			Rule.Priority(1),
			Rule.ColorScale(),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeAboveAverage),
			Rule.Priority(1),
			Rule.ColorScale(
				Value(ValueTypePercent, "10", false),
				"#112233",
			),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeAboveAverage),
			Rule.Priority(1),
			Rule.ColorScale(
				Value(ValueTypePercent, "10", false),
				Value(ValueTypePercent, "50", false),
				"#112233",
			),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeAboveAverage),
			Rule.Priority(1),
			Rule.ColorScale(
				Value(ValueTypePercent, "50", false),
				"#112233",
				"#334455",
			),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeAboveAverage),
			Rule.Priority(1),
			Rule.ColorScale(
				Value(ValueTypePercent, "10", false),
				Value(ValueTypePercent, "50", false),
				"#112233",
				"#334455",
			),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeAboveAverage),
			Rule.Priority(1),
			Rule.IconSet(IconSetType3Arrows, true, true, true),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeAboveAverage),
			Rule.Priority(1),
			Rule.IconSet(IconSetType3Arrows, true, true, true,
				Value(ValueTypePercent, "10", false),
			),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			Rule.Type(TypeAboveAverage),
			Rule.Priority(1),
			Rule.IconSet(IconSetType3Arrows, true, true, true,
				Value(ValueTypePercent, "10", false),
				Value(ValueTypePercent, "50", false),
			),
		),
	).Validate())
}
