package rule

import (
	//"github.com/plandem/xlsx/internal/ml"
	//"github.com/plandem/xlsx/internal/ml/primitives"
	//"github.com/stretchr/testify/require"
	"testing"
)

func TestIconSet(t *testing.T) {
	//	conditions := New(
	//		AddRule(
	//			Rule.IconSet.Type(IconSet4Arrows),
	//			Rule.IconSet.ReverseIcons,
	//			Rule.IconSet.IconsOnly,
	//		),
	//	)
	//
	//	require.Equal(t, &Info{
	//		info: &ml.ConditionalFormatting{},
	//		rules: []*ruleInfo{
	//			{
	//				initialized: true,
	//				rule: &ml.ConditionalRule{
	//					Type: primitives.ConditionTypeDataBar,
	//					DataBar: &ml.DataBar{
	//						Values: []*ml.ConditionValue{
	//							{
	//								Type:  primitives.ConditionValueTypeMin,
	//								Value: "1",
	//							},
	//							{
	//								Type:  primitives.ConditionValueTypeMax,
	//								Value: "50",
	//							},
	//						},
	//						Color: &ml.Color{
	//							RGB: "FF638EC6",
	//						},
	//						ShowValue: true,
	//					},
	//				},
	//			},
	//		},
	//	}, conditions)
	//
	//	conditions = New(
	//		AddRule(
	//			Rule.DataBar.Min("1"),
	//			Rule.DataBar.Max("50"),
	//			Rule.DataBar.Color("#110000"),
	//			Rule.DataBar.BarOnly,
	//			Rule.DataBar.MinLength(10),
	//			Rule.DataBar.MaxLength(20),
	//		),
	//	)
	//
	//	require.Equal(t, &Info{
	//		info: &ml.ConditionalFormatting{},
	//		rules: []*ruleInfo{
	//			{
	//				initialized: true,
	//				rule: &ml.ConditionalRule{
	//					Type: primitives.ConditionTypeDataBar,
	//					DataBar: &ml.DataBar{
	//						Values: []*ml.ConditionValue{
	//							{
	//								Type:  primitives.ConditionValueTypeMin,
	//								Value: "1",
	//							},
	//							{
	//								Type:  primitives.ConditionValueTypeMax,
	//								Value: "50",
	//							},
	//						},
	//						ShowValue: false,
	//						MinLength: 10,
	//						MaxLength: 20,
	//						Color: &ml.Color{
	//							RGB: "FF110000",
	//						},
	//					},
	//				},
	//			},
	//		},
	//	}, conditions)
}
