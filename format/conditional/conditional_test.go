package conditional

import (
	"github.com/plandem/xlsx/format/conditional/rule"
	//"github.com/plandem/xlsx/format/styles"
	//"github.com/plandem/xlsx/internal/ml"
	//"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditionalFormat_Set(t *testing.T) {
	//s := styles.New(
	//	styles.Font.Bold,
	//	styles.Font.Color("#112233"),
	//)
	//
	//conditions := New(
	//	Pivot,
	//	Refs("A10:B20"),
	//	AddRule(
	//		rule.Priority(10),
	//		rule.Type(rule.TypeCellIs),
	//	),
	//	AddRule(
	//		rule.Priority(90),
	//		rule.Type(rule.TypeAboveAverage),
	//		rule.Style(s),
	//	),
	//)

	//require.Equal(t, &Info{
	//	info: &ml.ConditionalFormatting{
	//		Pivot:  true,
	//		Bounds: primitives.BoundsListFromRefs("A10:B20"),
	//	},
	//	rules: []*rule.Info{
	//		{
	//			&ml.ConditionalRule{
	//				Type:     rule.TypeCellIs,
	//				Priority: 10,
	//			},
	//		},
	//		{
	//			rule: &ml.ConditionalRule{
	//				Type:     rule.TypeAboveAverage,
	//				Priority: 90,
	//			},
	//			style: s,
	//		},
	//	},
	//}, conditions)
}

func TestConditionalFormat_Validate(t *testing.T) {
	require.NotNil(t, New().Validate())

	require.NotNil(t, New(Refs("A10:B20")).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Priority(-1),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeCellIs),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeCellIs),
			rule.Priority(1),
			rule.Operator(rule.OperatorLessThanOrEqual),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeTop10),
			rule.Priority(1),
			//Rule.Rank(),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeTop10),
			rule.Priority(1),
			rule.Rank(1),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeContainsText),
			rule.Priority(1),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeContainsText),
			rule.Priority(1),
			rule.Text("abc"),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeTimePeriod),
			rule.Priority(1),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeTimePeriod),
			rule.Priority(1),
			rule.TimePeriod(rule.TimePeriodLastMonth),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeAboveAverage),
			rule.Priority(1),
			rule.ColorScale(),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeAboveAverage),
			rule.Priority(1),
			rule.ColorScale(
				rule.Value(rule.ValueTypePercent, "10", false),
				"#112233",
			),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeAboveAverage),
			rule.Priority(1),
			rule.ColorScale(
				rule.Value(rule.ValueTypePercent, "10", false),
				rule.Value(rule.ValueTypePercent, "50", false),
				"#112233",
			),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeAboveAverage),
			rule.Priority(1),
			rule.ColorScale(
				rule.Value(rule.ValueTypePercent, "50", false),
				"#112233",
				"#334455",
			),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeAboveAverage),
			rule.Priority(1),
			rule.ColorScale(
				rule.Value(rule.ValueTypePercent, "10", false),
				rule.Value(rule.ValueTypePercent, "50", false),
				"#112233",
				"#334455",
			),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeAboveAverage),
			rule.Priority(1),
			rule.IconSet(rule.IconSetType3Arrows, true, true, true),
		),
	).Validate())

	require.NotNil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeAboveAverage),
			rule.Priority(1),
			rule.IconSet(rule.IconSetType3Arrows, true, true, true,
				rule.Value(rule.ValueTypePercent, "10", false),
			),
		),
	).Validate())

	require.Nil(t, New(
		Refs("A10:B20"),
		AddRule(
			rule.Type(rule.TypeAboveAverage),
			rule.Priority(1),
			rule.IconSet(rule.IconSetType3Arrows, true, true, true,
				rule.Value(rule.ValueTypePercent, "10", false),
				rule.Value(rule.ValueTypePercent, "50", false),
			),
		),
	).Validate())
}
