package conditional

import (
	//"github.com/plandem/xlsx/format/styles"
	//"github.com/plandem/xlsx/internal/ml"
	//"github.com/plandem/xlsx/internal/ml/primitives"
	"testing"
)

func TestConditionalFormat_Set(t *testing.T) {
	//z := New(Type.DataBar.Z())
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

//func TestConditionalFormat_Validate(t *testing.T) {
//	require.NotNil(t, New().Validate())
//
//	require.NotNil(t, New(Refs("A10:B20")).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Priority(-1),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeCellIs),
//		),
//	).Validate())
//
//	require.Nil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeCellIs),
//			rule1.Priority(1),
//			rule1.Operator(primitives.OperatorLessThanOrEqual),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeTop10),
//			rule1.Priority(1),
//			//Rule.Rank(),
//		),
//	).Validate())
//
//	require.Nil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeTop10),
//			rule1.Priority(1),
//			rule1.Rank(1),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeContainsText),
//			rule1.Priority(1),
//		),
//	).Validate())
//
//	require.Nil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeContainsText),
//			rule1.Priority(1),
//			rule1.Text("abc"),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeTimePeriod),
//			rule1.Priority(1),
//		),
//	).Validate())
//
//	require.Nil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeTimePeriod),
//			rule1.Priority(1),
//			rule1.TimePeriod(primitives.TimePeriodLastMonth),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeAboveAverage),
//			rule1.Priority(1),
//			rule1.ColorScale(),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeAboveAverage),
//			rule1.Priority(1),
//			rule1.ColorScale(
//				rule1.Value(primitives.ConditionValueTypePercent, "10", false),
//				"#112233",
//			),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeAboveAverage),
//			rule1.Priority(1),
//			rule1.ColorScale(
//				rule1.Value(primitives.ConditionValueTypePercent, "10", false),
//				rule1.Value(primitives.ConditionValueTypePercent, "50", false),
//				"#112233",
//			),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeAboveAverage),
//			rule1.Priority(1),
//			rule1.ColorScale(
//				rule1.Value(primitives.ConditionValueTypePercent, "50", false),
//				"#112233",
//				"#334455",
//			),
//		),
//	).Validate())
//
//	require.Nil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeAboveAverage),
//			rule1.Priority(1),
//			rule1.ColorScale(
//				rule1.Value(primitives.ConditionValueTypePercent, "10", false),
//				rule1.Value(primitives.ConditionValueTypePercent, "50", false),
//				"#112233",
//				"#334455",
//			),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeAboveAverage),
//			rule1.Priority(1),
//			rule1.IconSet(primitives.IconSetType3Arrows, true, true, true),
//		),
//	).Validate())
//
//	require.NotNil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeAboveAverage),
//			rule1.Priority(1),
//			rule1.IconSet(primitives.IconSetType3Arrows, true, true, true,
//				rule1.Value(primitives.ConditionValueTypePercent, "10", false),
//			),
//		),
//	).Validate())
//
//	require.Nil(t, New(
//		Refs("A10:B20"),
//		AddRule(
//			rule1.Type(primitives.ConditionTypeAboveAverage),
//			rule1.Priority(1),
//			rule1.IconSet(primitives.IconSetType3Arrows, true, true, true,
//				rule1.Value(primitives.ConditionValueTypePercent, "10", false),
//				rule1.Value(primitives.ConditionValueTypePercent, "50", false),
//			),
//		),
//	).Validate())
//}
