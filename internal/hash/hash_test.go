package hash_test

import (
	"fmt"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/hash"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignment(t *testing.T) {
	require.Equal(t, hash.Key("0:0:0:false:0:0:false:false:0"), hash.Alignment(nil))
	require.Equal(t, hash.Key("0:0:0:false:0:0:false:false:0"), hash.Alignment(&ml.CellAlignment{}))
	require.Equal(t, hash.Key("0:3:0:false:0:0:false:false:0"), hash.Alignment(&ml.CellAlignment{Vertical: styles.VAlignBottom}))
	require.Equal(t, hash.Key("8:0:0:false:0:0:false:false:0"), hash.Alignment(&ml.CellAlignment{Horizontal: styles.HAlignDistributed}))
	require.Equal(t, hash.Key("0:0:90:false:0:0:false:false:0"), hash.Alignment(&ml.CellAlignment{TextRotation: 90}))
	require.Equal(t, hash.Key("0:0:0:true:0:0:false:false:0"), hash.Alignment(&ml.CellAlignment{WrapText: true}))
	require.Equal(t, hash.Key("0:0:0:false:10:0:false:false:0"), hash.Alignment(&ml.CellAlignment{Indent: 10}))
	require.Equal(t, hash.Key("0:0:0:false:0:12:false:false:0"), hash.Alignment(&ml.CellAlignment{RelativeIndent: 12}))
	require.Equal(t, hash.Key("0:0:0:false:0:0:true:false:0"), hash.Alignment(&ml.CellAlignment{JustifyLastLine: true}))
	require.Equal(t, hash.Key("0:0:0:false:0:0:false:true:0"), hash.Alignment(&ml.CellAlignment{ShrinkToFit: true}))
	require.Equal(t, hash.Key("0:0:0:false:0:0:false:false:13"), hash.Alignment(&ml.CellAlignment{ReadingOrder: 13}))
	require.Equal(t, hash.Key("8:3:90:true:10:12:true:true:13"), hash.Alignment(&ml.CellAlignment{
		Horizontal:      styles.HAlignDistributed,
		Vertical:        styles.VAlignBottom,
		ReadingOrder:    13,
		TextRotation:    90,
		ShrinkToFit:     true,
		JustifyLastLine: true,
		WrapText:        true,
		Indent:          10,
		RelativeIndent:  12,
	}))
}

func TestBorder(t *testing.T) {
	b := ml.Border{}
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(nil))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&b))
	require.Equal(t, &ml.Border{}, &b)

	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Left: &ml.BorderSegment{}}))
	require.Equal(t, hash.Key("false:112233:0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Left: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, hash.Key("false::0:::mediumDashDot:false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Left: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}))

	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Right: &ml.BorderSegment{}}))
	require.Equal(t, hash.Key("false::0::::false:112233:0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Right: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, hash.Key("false::0::::false::0:::mediumDashDot:false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Right: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}))

	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Top: &ml.BorderSegment{}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false:112233:0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Top: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0:::mediumDashDot:false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Top: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}))

	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Bottom: &ml.BorderSegment{}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false:112233:0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Bottom: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0:::mediumDashDot:false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Bottom: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}))

	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Diagonal: &ml.BorderSegment{}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false:112233:0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Diagonal: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0:::mediumDashDot:false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Diagonal: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}))

	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Vertical: &ml.BorderSegment{}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false:112233:0::::false::0::::false:false:false"), hash.Border(&ml.Border{Vertical: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0:::mediumDashDot:false::0::::false:false:false"), hash.Border(&ml.Border{Vertical: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}))

	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false"), hash.Border(&ml.Border{Horizontal: &ml.BorderSegment{}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:112233:0::::false:false:false"), hash.Border(&ml.Border{Horizontal: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0:::mediumDashDot:false:false:false"), hash.Border(&ml.Border{Horizontal: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}))

	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::true:false:false"), hash.Border(&ml.Border{DiagonalUp: true}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:true:false"), hash.Border(&ml.Border{DiagonalDown: true}))
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:true"), hash.Border(&ml.Border{Outline: true}))

	require.Equal(t, hash.Key("false:777777:0:::slantDashDot:false:666666:0:::thin:false:555555:0:::thick:false:444444:0:::hair:false:333333:0:::dotted:false:222222:0:::dashDot:false:111111:0:::medium:true:true:true"), hash.Border(&ml.Border{
		Outline:      true,
		DiagonalDown: true,
		DiagonalUp:   true,
		Horizontal:   &ml.BorderSegment{Color: &ml.Color{RGB: "111111"}, Type: styles.BorderStyleMedium},
		Vertical:     &ml.BorderSegment{Color: &ml.Color{RGB: "222222"}, Type: styles.BorderStyleDashDot},
		Diagonal:     &ml.BorderSegment{Color: &ml.Color{RGB: "333333"}, Type: styles.BorderStyleDotted},
		Bottom:       &ml.BorderSegment{Color: &ml.Color{RGB: "444444"}, Type: styles.BorderStyleHair},
		Top:          &ml.BorderSegment{Color: &ml.Color{RGB: "555555"}, Type: styles.BorderStyleThick},
		Right:        &ml.BorderSegment{Color: &ml.Color{RGB: "666666"}, Type: styles.BorderStyleThin},
		Left:         &ml.BorderSegment{Color: &ml.Color{RGB: "777777"}, Type: styles.BorderStyleSlantDashDot},
	}))
}

func TestColor(t *testing.T) {
	index1 := 1
	index2 := 2

	require.Equal(t, hash.Key("false::0::"), hash.Color(nil))
	require.Equal(t, hash.Key("false::0::"), hash.Color(&ml.Color{}))
	require.Equal(t, hash.Key("true::0::"), hash.Color(&ml.Color{Auto: true}))
	require.Equal(t, hash.Key("false:112233:0::"), hash.Color(&ml.Color{RGB: "112233"}))
	require.Equal(t, hash.Key("false::1::"), hash.Color(&ml.Color{Tint: 1}))
	require.Equal(t, hash.Key("false::0:1:"), hash.Color(&ml.Color{Indexed: &index1}))
	require.Equal(t, hash.Key("false::0::2"), hash.Color(&ml.Color{Theme: &index2}))
	require.Equal(t, hash.Key("true:112233:1:1:2"), hash.Color(&ml.Color{
		Indexed: &index1,
		RGB:     "112233",
		Auto:    true,
		Theme:   &index2,
		Tint:    1,
	}))
}

func TestFill(t *testing.T) {
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:0:0:0"), hash.Fill(nil))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:0:0:0"), hash.Fill(&ml.Fill{}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:0:0:0"), hash.Fill(&ml.Fill{Pattern: &ml.PatternFill{}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:0:0:0"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{}}))
	require.Equal(t, hash.Key("0:false:112233:0:::false::0:::0:0:0:0:0:0"), hash.Fill(&ml.Fill{Pattern: &ml.PatternFill{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, hash.Key("0:false::0:::false:112233:0:::0:0:0:0:0:0"), hash.Fill(&ml.Fill{Pattern: &ml.PatternFill{Background: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, hash.Key("11:false::0:::false::0:::0:0:0:0:0:0"), hash.Fill(&ml.Fill{Pattern: &ml.PatternFill{Type: styles.PatternTypeDarkTrellis}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:90:0:0:0:0"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{Degree: 90}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:1.1:0:0:0"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{Left: 1.1}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:1.1:0:0"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{Right: 1.1}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:0:1.1:0"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{Top: 1.1}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:0:0:1.1"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{Bottom: 1.1}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::1:0:0:0:0:0"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{Type: styles.GradientTypePath}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:0:0:0"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{Type: styles.GradientTypeLinear}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:0:0:0:1.1:false:112233:0::"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{Stop: []*ml.GradientStop{{Color: &ml.Color{RGB: "112233"}, Position: 1.1}}}}))
	require.Equal(t, hash.Key("0:false::0:::false::0:::0:0:0:0:0:0:1.1:false:112233:0:::2.2:false:AABBCC:0::"), hash.Fill(&ml.Fill{Gradient: &ml.GradientFill{Stop: []*ml.GradientStop{
		{Color: &ml.Color{RGB: "112233"}, Position: 1.1},
		{Color: &ml.Color{RGB: "AABBCC"}, Position: 2.2},
	}}}))

	require.Equal(t, hash.Key("11:false:112233:0:::false:112233:0:::1:90:1.1:2.2:3.3:4.4:1.1:false:112233:0:::2.2:false:AABBCC:0::"), hash.Fill(&ml.Fill{
		Gradient: &ml.GradientFill{
			Type:   styles.GradientTypePath,
			Degree: 90,
			Left:   1.1,
			Right:  2.2,
			Top:    3.3,
			Bottom: 4.4,
			Stop: []*ml.GradientStop{
				{Color: &ml.Color{RGB: "112233"}, Position: 1.1},
				{Color: &ml.Color{RGB: "AABBCC"}, Position: 2.2},
			},
		},
		Pattern: &ml.PatternFill{
			Type:       styles.PatternTypeDarkTrellis,
			Background: &ml.Color{RGB: "112233"},
			Color:      &ml.Color{RGB: "112233"},
		},
	}))
}

func TestFont(t *testing.T) {
	require.Equal(t, hash.Key(":0:0:false:false:false:false:false:false:false::0:::0:::"), hash.Font(nil))
	require.Equal(t, hash.Key(":0:0:false:false:false:false:false:false:false::0:::0:::"), hash.Font(&ml.Font{}))
	require.Equal(t, hash.Key("calibri:0:0:false:false:false:false:false:false:false::0:::0:::"), hash.Font(&ml.Font{Name: "calibri"}))
	require.Equal(t, hash.Key(":1:0:false:false:false:false:false:false:false::0:::0:::"), hash.Font(&ml.Font{Charset: 1}))
	require.Equal(t, hash.Key(":0:1:false:false:false:false:false:false:false::0:::0:::"), hash.Font(&ml.Font{Family: styles.FontFamilyRoman}))
	require.Equal(t, hash.Key(":0:0:false:false:false:false:false:false:false::0:::0:::"), hash.Font(&ml.Font{Bold: false}))
	require.Equal(t, hash.Key(":0:0:true:false:false:false:false:false:false::0:::0:::"), hash.Font(&ml.Font{Bold: true}))
	require.Equal(t, hash.Key(":0:0:false:true:false:false:false:false:false::0:::0:::"), hash.Font(&ml.Font{Italic: true}))
	require.Equal(t, hash.Key(":0:0:false:false:true:false:false:false:false::0:::0:::"), hash.Font(&ml.Font{Strike: true}))
	require.Equal(t, hash.Key(":0:0:false:false:false:true:false:false:false::0:::0:::"), hash.Font(&ml.Font{Shadow: true}))
	require.Equal(t, hash.Key(":0:0:false:false:false:false:true:false:false::0:::0:::"), hash.Font(&ml.Font{Condense: true}))
	require.Equal(t, hash.Key(":0:0:false:false:false:false:false:true:false::0:::0:::"), hash.Font(&ml.Font{Extend: true}))
	require.Equal(t, hash.Key(":0:0:false:false:false:false:false:false:false:112233:0:::0:::"), hash.Font(&ml.Font{Color: &ml.Color{RGB: "112233"}}))
	require.Equal(t, hash.Key(":0:0:false:false:false:false:false:false:false::0:::2.2:::"), hash.Font(&ml.Font{Size: 2.2}))
	require.Equal(t, hash.Key(":0:0:false:false:false:false:false:false:false::0:::0:doubleAccounting::"), hash.Font(&ml.Font{Underline: styles.UnderlineTypeDoubleAccounting}))
	require.Equal(t, hash.Key(":0:0:false:false:false:false:false:false:false::0:::0::subscript:"), hash.Font(&ml.Font{VAlign: styles.FontVAlignSubscript}))
	require.Equal(t, hash.Key(":0:0:false:false:false:false:false:false:false::0:::0:::major"), hash.Font(&ml.Font{Scheme: styles.FontSchemeMajor}))
	require.Equal(t, hash.Key("calibri:1:1:true:true:true:true:true:true:false:112233:0:::2.2:doubleAccounting:subscript:major"), hash.Font(&ml.Font{
		Scheme:    styles.FontSchemeMajor,
		VAlign:    styles.FontVAlignSubscript,
		Underline: styles.UnderlineTypeDoubleAccounting,
		Size:      2.2,
		Color:     &ml.Color{RGB: "112233"},
		Extend:    true,
		Condense:  true,
		Shadow:    true,
		Strike:    true,
		Italic:    true,
		Bold:      true,
		Family:    styles.FontFamilyRoman,
		Charset:   1,
		Name:      "calibri",
	}))
}

func TestNumber(t *testing.T) {
	require.Equal(t, hash.Key("0:"), hash.NumberFormat(nil))
	require.Equal(t, hash.Key("0:"), hash.NumberFormat(&ml.NumberFormat{}))
	require.Equal(t, hash.Key("1:"), hash.NumberFormat(&ml.NumberFormat{ID: 1}))
	require.Equal(t, hash.Key("0:aaa"), hash.NumberFormat(&ml.NumberFormat{Code: "aaa"}))
	require.Equal(t, hash.Key("1:aaa"), hash.NumberFormat(&ml.NumberFormat{ID: 1, Code: "aaa"}))
}

func TestProtection(t *testing.T) {
	require.Equal(t, hash.Key("false:false"), hash.Protection(nil))
	require.Equal(t, hash.Key("false:false"), hash.Protection(&ml.CellProtection{}))
	require.Equal(t, hash.Key("true:false"), hash.Protection(&ml.CellProtection{Locked: true}))
	require.Equal(t, hash.Key("false:true"), hash.Protection(&ml.CellProtection{Hidden: true}))
	require.Equal(t, hash.Key("true:true"), hash.Protection(&ml.CellProtection{Locked: true, Hidden: true}))
}

func TestDirectStyle(t *testing.T) {
	require.Equal(t, hash.Key("0:0:0:0:false:false:false:false:false:false:false:false:0:0:0:false:0:0:false:false:0:false:false::0"), hash.DirectStyle(nil))
	require.Equal(t, hash.Key("-1:-2:-3:-4:true:true:true:true:true:true:true:true:8:3:90:true:10:12:true:true:13:true:true::-10"), hash.DirectStyle(&ml.DirectStyle{
		Style: ml.Style{
			NumFmtId:          -1,
			FontId:            -2,
			FillId:            -3,
			BorderId:          -4,
			QuotePrefix:       true,
			PivotButton:       true,
			ApplyNumberFormat: true,
			ApplyFont:         true,
			ApplyFill:         true,
			ApplyBorder:       true,
			ApplyAlignment:    true,
			ApplyProtection:   true,
			Alignment: &ml.CellAlignment{
				Horizontal:      styles.HAlignDistributed,
				Vertical:        styles.VAlignBottom,
				ReadingOrder:    13,
				TextRotation:    90,
				ShrinkToFit:     true,
				JustifyLastLine: true,
				WrapText:        true,
				Indent:          10,
				RelativeIndent:  12,
			},
			Protection: &ml.CellProtection{Locked: true, Hidden: true},
		},
		XfId: -10,
	}))
}

func TestNamedStyle(t *testing.T) {
	require.Equal(t, hash.Key("0:0:0:0:false:false:false:false:false:false:false:false:0:0:0:false:0:0:false:false:0:false:false:"), hash.NamedStyle(nil))
	require.Equal(t, hash.Key("-1:-2:-3:-4:true:true:true:true:true:true:true:true:8:3:90:true:10:12:true:true:13:true:true:"), hash.NamedStyle(&ml.NamedStyle{
		NumFmtId:          -1,
		FontId:            -2,
		FillId:            -3,
		BorderId:          -4,
		QuotePrefix:       true,
		PivotButton:       true,
		ApplyNumberFormat: true,
		ApplyFont:         true,
		ApplyFill:         true,
		ApplyBorder:       true,
		ApplyAlignment:    true,
		ApplyProtection:   true,
		Alignment: &ml.CellAlignment{
			Horizontal:      styles.HAlignDistributed,
			Vertical:        styles.VAlignBottom,
			ReadingOrder:    13,
			TextRotation:    90,
			ShrinkToFit:     true,
			JustifyLastLine: true,
			WrapText:        true,
			Indent:          10,
			RelativeIndent:  12,
		},
		Protection: &ml.CellProtection{Locked: true, Hidden: true},
	}))
}

func TestDiffStyle(t *testing.T) {
	require.Equal(t, hash.Key("false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false:0:false::0:::false::0:::0:0:0:0:0:0::0:0:false:false:false:false:false:false:false::0:::0::::0::0:0:0:false:0:0:false:false:0:false:false:"), hash.DiffStyle(nil))
	require.Equal(t, hash.Key("false:777777:0:::slantDashDot:false:666666:0:::thin:false:555555:0:::thick:false:444444:0:::hair:false:333333:0:::dotted:false:222222:0:::dashDot:false:111111:0:::medium:true:true:true:11:false:112233:0:::false:112233:0:::1:90:1.1:2.2:3.3:4.4:1.1:false:112233:0:::2.2:false:AABBCC:0:::calibri:1:1:true:true:true:true:true:true:false:112233:0:::2.2:doubleAccounting:subscript:major:1:aaa:8:3:90:true:10:12:true:true:13:true:true:"), hash.DiffStyle(&ml.DiffStyle{
		Font: &ml.Font{
			Scheme:    styles.FontSchemeMajor,
			VAlign:    styles.FontVAlignSubscript,
			Underline: styles.UnderlineTypeDoubleAccounting,
			Size:      2.2,
			Color:     &ml.Color{RGB: "112233"},
			Extend:    true,
			Condense:  true,
			Shadow:    true,
			Strike:    true,
			Italic:    true,
			Bold:      true,
			Family:    styles.FontFamilyRoman,
			Charset:   1,
			Name:      "calibri",
		},
		Fill: &ml.Fill{
			Gradient: &ml.GradientFill{
				Type:   styles.GradientTypePath,
				Degree: 90,
				Left:   1.1,
				Right:  2.2,
				Top:    3.3,
				Bottom: 4.4,
				Stop: []*ml.GradientStop{
					{Color: &ml.Color{RGB: "112233"}, Position: 1.1},
					{Color: &ml.Color{RGB: "AABBCC"}, Position: 2.2},
				},
			},
			Pattern: &ml.PatternFill{
				Type:       styles.PatternTypeDarkTrellis,
				Background: &ml.Color{RGB: "112233"},
				Color:      &ml.Color{RGB: "112233"},
			},
		},
		Alignment: &ml.CellAlignment{
			Horizontal:      styles.HAlignDistributed,
			Vertical:        styles.VAlignBottom,
			ReadingOrder:    13,
			TextRotation:    90,
			ShrinkToFit:     true,
			JustifyLastLine: true,
			WrapText:        true,
			Indent:          10,
			RelativeIndent:  12,
		},
		NumberFormat: &ml.NumberFormat{ID: 1, Code: "aaa"},
		Protection:   &ml.CellProtection{Locked: true, Hidden: true},
		Border: &ml.Border{
			Outline:      true,
			DiagonalDown: true,
			DiagonalUp:   true,
			Horizontal:   &ml.BorderSegment{Color: &ml.Color{RGB: "111111"}, Type: styles.BorderStyleMedium},
			Vertical:     &ml.BorderSegment{Color: &ml.Color{RGB: "222222"}, Type: styles.BorderStyleDashDot},
			Diagonal:     &ml.BorderSegment{Color: &ml.Color{RGB: "333333"}, Type: styles.BorderStyleDotted},
			Bottom:       &ml.BorderSegment{Color: &ml.Color{RGB: "444444"}, Type: styles.BorderStyleHair},
			Top:          &ml.BorderSegment{Color: &ml.Color{RGB: "555555"}, Type: styles.BorderStyleThick},
			Right:        &ml.BorderSegment{Color: &ml.Color{RGB: "666666"}, Type: styles.BorderStyleThin},
			Left:         &ml.BorderSegment{Color: &ml.Color{RGB: "777777"}, Type: styles.BorderStyleSlantDashDot},
		}},
	))
}

func TestVmlShape(t *testing.T) {

	shape := &vml.Shape{}
	shape.ID = fmt.Sprintf("_x0000_s%d", 1025)
	shape.Type = "#_x0000_t202"
	shape.FillColor = "#ffeeee"
	shape.InsetMode = vml.InsetModeAuto

	shape.ClientData = &vml.ClientData{
		Column: 1,
		Row:    2,
	}

	require.Equal(t, hash.Key("#_x0000_t202:1:2"), hash.Vml(shape))
}
