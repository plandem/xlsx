// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml_test

import (
	"github.com/plandem/ooxml/index"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignment_Hash(t *testing.T) {
	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.CellAlignment{}, 1))
	require.NotNil(t, idx.Add((*ml.CellAlignment)(nil), 1))

	require.Nil(t, idx.Add(&ml.CellAlignment{Vertical: styles.VAlignBottom}, 1))
	require.Nil(t, idx.Add(&ml.CellAlignment{Horizontal: styles.HAlignDistributed}, 1))
	require.Nil(t, idx.Add(&ml.CellAlignment{TextRotation: 90}, 1))
	require.Nil(t, idx.Add(&ml.CellAlignment{WrapText: true}, 1))
	require.Nil(t, idx.Add(&ml.CellAlignment{Indent: 10}, 1))
	require.Nil(t, idx.Add(&ml.CellAlignment{RelativeIndent: 12}, 1))
	require.Nil(t, idx.Add(&ml.CellAlignment{JustifyLastLine: true}, 1))
	require.Nil(t, idx.Add(&ml.CellAlignment{ShrinkToFit: true}, 1))
	require.Nil(t, idx.Add(&ml.CellAlignment{ReadingOrder: 13}, 1))
	require.Nil(t, idx.Add(&ml.CellAlignment{
		Horizontal:      styles.HAlignDistributed,
		Vertical:        styles.VAlignBottom,
		ReadingOrder:    13,
		TextRotation:    90,
		ShrinkToFit:     true,
		JustifyLastLine: true,
		WrapText:        true,
		Indent:          10,
		RelativeIndent:  12,
	}, 1))
}

func TestBorder_Hash(t *testing.T) {
	b := &ml.Border{}
	_ = b.Hash()
	//b should not be mutated
	require.Equal(t, &ml.Border{}, b)

	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.Border{}, 1))
	require.NotNil(t, idx.Add((*ml.Border)(nil), 1))

	require.NotNil(t, idx.Add(&ml.Border{Left: &ml.BorderSegment{}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Left: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Left: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}, 1))

	require.NotNil(t, idx.Add(&ml.Border{Right: &ml.BorderSegment{}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Right: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Right: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}, 1))

	require.NotNil(t, idx.Add(&ml.Border{Top: &ml.BorderSegment{}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Top: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Top: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}, 1))

	require.NotNil(t, idx.Add(&ml.Border{Bottom: &ml.BorderSegment{}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Bottom: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Bottom: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}, 1))

	require.NotNil(t, idx.Add(&ml.Border{Diagonal: &ml.BorderSegment{}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Diagonal: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Diagonal: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}, 1))

	require.NotNil(t, idx.Add(&ml.Border{Vertical: &ml.BorderSegment{}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Vertical: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Vertical: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}, 1))

	require.NotNil(t, idx.Add(&ml.Border{Horizontal: &ml.BorderSegment{}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Horizontal: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}, 1))
	require.Nil(t, idx.Add(&ml.Border{Horizontal: &ml.BorderSegment{Type: styles.BorderStyleMediumDashDot}}, 1))

	require.Nil(t, idx.Add(&ml.Border{DiagonalUp: true}, 1))
	require.Nil(t, idx.Add(&ml.Border{DiagonalDown: true}, 1))
	require.Nil(t, idx.Add(&ml.Border{Outline: true}, 1))

	require.Nil(t, idx.Add(&ml.Border{
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
	}, 1))
}

func TestColor_Hash(t *testing.T) {
	index1 := 1
	index2 := 2

	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.Color{}, 1))
	require.NotNil(t, idx.Add((*ml.Color)(nil), 1))

	require.Nil(t, idx.Add(&ml.Color{Auto: true}, 1))
	require.Nil(t, idx.Add(&ml.Color{RGB: "112233"}, 1))
	require.Nil(t, idx.Add(&ml.Color{Tint: 1}, 1))
	require.Nil(t, idx.Add(&ml.Color{Indexed: &index1}, 1))
	require.Nil(t, idx.Add(&ml.Color{Theme: &index2}, 1))
	require.Nil(t, idx.Add(&ml.Color{
		Indexed: &index1,
		RGB:     "112233",
		Auto:    true,
		Theme:   &index2,
		Tint:    1,
	}, 1))
}

func TestFill_Hash(t *testing.T) {
	f := &ml.Fill{}
	_ = f.Hash()
	//b should not be mutated
	require.Equal(t, &ml.Fill{}, f)

	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.Fill{}, 1))
	require.NotNil(t, idx.Add((*ml.Fill)(nil), 1))
	require.NotNil(t, idx.Add(&ml.Fill{Pattern: &ml.PatternFill{}}, 1))
	require.NotNil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Pattern: &ml.PatternFill{Color: &ml.Color{RGB: "112233"}}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Pattern: &ml.PatternFill{Background: &ml.Color{RGB: "112233"}}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Pattern: &ml.PatternFill{Type: styles.PatternTypeDarkTrellis}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{Degree: 90}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{Left: 1.1}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{Right: 1.1}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{Top: 1.1}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{Bottom: 1.1}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{Type: styles.GradientTypePath}}, 1))
	require.NotNil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{Type: styles.GradientTypeLinear}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{Stop: []*ml.GradientStop{{Color: &ml.Color{RGB: "112233"}, Position: 1.1}}}}, 1))
	require.Nil(t, idx.Add(&ml.Fill{Gradient: &ml.GradientFill{Stop: []*ml.GradientStop{
		{Color: &ml.Color{RGB: "112233"}, Position: 1.1},
		{Color: &ml.Color{RGB: "AABBCC"}, Position: 2.2},
	}}}, 1))

	require.Nil(t, idx.Add(&ml.Fill{
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
	}, 1))
}

func TestFont_Hash(t *testing.T) {
	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.Font{}, 1))
	require.NotNil(t, idx.Add((*ml.Font)(nil), 1))

	require.Nil(t, idx.Add(&ml.Font{Name: "calibri"}, 1))
	require.Nil(t, idx.Add(&ml.Font{Charset: 1}, 1))
	require.Nil(t, idx.Add(&ml.Font{Family: styles.FontFamilyRoman}, 1))
	require.NotNil(t, idx.Add(&ml.Font{Bold: false}, 1))
	require.Nil(t, idx.Add(&ml.Font{Bold: true}, 1))
	require.Nil(t, idx.Add(&ml.Font{Italic: true}, 1))
	require.Nil(t, idx.Add(&ml.Font{Strike: true}, 1))
	require.Nil(t, idx.Add(&ml.Font{Shadow: true}, 1))
	require.Nil(t, idx.Add(&ml.Font{Condense: true}, 1))
	require.Nil(t, idx.Add(&ml.Font{Extend: true}, 1))
	require.Nil(t, idx.Add(&ml.Font{Color: &ml.Color{RGB: "112233"}}, 1))
	require.Nil(t, idx.Add(&ml.Font{Size: 2.2}, 1))
	require.Nil(t, idx.Add(&ml.Font{Underline: styles.UnderlineTypeDoubleAccounting}, 1))
	require.Nil(t, idx.Add(&ml.Font{VAlign: styles.FontVAlignSubscript}, 1))
	require.Nil(t, idx.Add(&ml.Font{Scheme: styles.FontSchemeMajor}, 1))
	require.Nil(t, idx.Add(&ml.Font{
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
	}, 1))
}

func TestNumber_Hash(t *testing.T) {
	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.NumberFormat{}, 1))
	require.NotNil(t, idx.Add((*ml.NumberFormat)(nil), 1))
	require.Nil(t, idx.Add(&ml.NumberFormat{ID: 1}, 1))
	require.Nil(t, idx.Add(&ml.NumberFormat{Code: "aaa"}, 1))
	require.Nil(t, idx.Add(&ml.NumberFormat{ID: 1, Code: "aaa"}, 1))
}

func TestProtection_Hash(t *testing.T) {
	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.CellProtection{}, 1))
	require.NotNil(t, idx.Add((*ml.CellProtection)(nil), 1))
	require.Nil(t, idx.Add(&ml.CellProtection{Locked: true}, 1))
	require.Nil(t, idx.Add(&ml.CellProtection{Hidden: true}, 1))
	require.Nil(t, idx.Add(&ml.CellProtection{Locked: true, Hidden: true}, 1))
}

func TestDirectStyle_Hash(t *testing.T) {
	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.DirectStyle{}, 1))
	require.NotNil(t, idx.Add((*ml.DirectStyle)(nil), 1))

	require.Nil(t, idx.Add(&ml.DirectStyle{
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
	}, 1))
}

func TestNamedStyle_Hash(t *testing.T) {
	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.NamedStyle{}, 1))
	require.NotNil(t, idx.Add((*ml.NamedStyle)(nil), 1))

	require.Nil(t, idx.Add(&ml.NamedStyle{
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
	}, 1))
}

func TestDiffStyle_Hash(t *testing.T) {
	idx := index.Index{}
	require.Nil(t, idx.Add(&ml.DiffStyle{}, 1))
	require.NotNil(t, idx.Add((*ml.DiffStyle)(nil), 1))

	require.Nil(t, idx.Add(&ml.DiffStyle{
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
		}}, 1))
}
