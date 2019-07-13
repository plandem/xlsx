// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"encoding/xml"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func createStylesAndFill(callback func(*Info)) *Info {
	f := New()
	callback(f)
	return f
}

func TestStyleFormat_Settings(t *testing.T) {
	style := New()

	//empty
	font, fill, alignment, number, protection, border, namedInfo := from(style)
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)
	require.Nil(t, namedInfo)

	//full featured
	style.Set(
		Alignment.VAlign(VAlignBottom),
		Alignment.HAlign(HAlignFill),
		Alignment.TextRotation(90),
		Alignment.WrapText,
		Alignment.Indent(10),
		Alignment.RelativeIndent(5),
		Alignment.JustifyLastLine,
		Alignment.ShrinkToFit,
		Alignment.ReadingOrder(4),
		Border.Type(BorderStyleDashDot),
		Border.Color("#FF00FF"),
		Border.Diagonal.Type(BorderStyleDashDot),
		Border.Diagonal.Color("#FF00FF"),
		Border.Vertical.Type(BorderStyleDashDot),
		Border.Vertical.Color("#FF00FF"),
		Border.Horizontal.Type(BorderStyleDashDot),
		Border.Horizontal.Color("#FF00FF"),
		Fill.Color("#FF00FF"),
		Fill.Background("#00FF00"),
		Fill.Type(PatternTypeDarkDown),
		Font.Name("Calibri"),
		Font.Size(10),
		Font.Bold,
		Font.Italic,
		Font.Strikeout,
		Font.Shadow,
		Font.Condense,
		Font.Extend,
		Font.Family(FontFamilyDecorative),
		Font.Color("#FF00FF"),
		Font.Underline(UnderlineTypeSingle),
		Font.VAlign(FontVAlignBaseline),
		Font.Scheme(FontSchemeMinor),
		NumberFormatID(8),
		Protection.Hidden,
		Protection.Locked,
	)

	font, fill, alignment, number, protection, border, namedInfo = from(style)
	require.NotNil(t, font)
	require.NotNil(t, fill)
	require.NotNil(t, alignment)
	require.NotNil(t, number)
	require.NotNil(t, protection)
	require.NotNil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.CellAlignment{
		Vertical:        VAlignBottom,
		Horizontal:      HAlignFill,
		TextRotation:    90,
		WrapText:        true,
		Indent:          10,
		RelativeIndent:  5,
		JustifyLastLine: true,
		ShrinkToFit:     true,
		ReadingOrder:    4,
	}, alignment)

	require.Equal(t, &ml.Border{
		Left: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Top: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Bottom: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Right: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Diagonal: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Vertical: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Horizontal: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
	}, border)

	require.Equal(t, &ml.Fill{
		Pattern: &ml.PatternFill{
			Color:      color.New("FFFF00FF"),
			Background: color.New("FF00FF00"),
			Type:       PatternTypeDarkDown,
		},
	}, fill)

	require.Equal(t, &ml.Font{
		Name:      "Calibri",
		Bold:      true,
		Italic:    true,
		Strike:    true,
		Shadow:    true,
		Condense:  true,
		Extend:    true,
		Size:      10.0,
		Color:     color.New("FFFF00FF"),
		Family:    FontFamilyDecorative,
		Underline: UnderlineTypeSingle,
		VAlign:    FontVAlignBaseline,
		Scheme:    FontSchemeMinor,
	}, font)

	require.Equal(t, &ml.NumberFormat{
		ID:   8,
		Code: "($#,##0.00_);[RED]($#,##0.00_)",
	}, number)

	require.Equal(t, &ml.CellProtection{
		Locked: true,
		Hidden: true,
	}, protection)
}

func TestStyleFormat_Settings_Alignment(t *testing.T) {
	style := New(
		Alignment.VAlign(VAlignBottom),
		Alignment.HAlign(HAlignFill),
		Alignment.TextRotation(90),
		Alignment.WrapText,
		Alignment.Indent(10),
		Alignment.RelativeIndent(5),
		Alignment.JustifyLastLine,
		Alignment.ShrinkToFit,
		Alignment.ReadingOrder(4),
	)

	font, fill, alignment, number, protection, border, namedInfo := from(style)
	require.Nil(t, font)
	require.Nil(t, fill)
	require.NotNil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.CellAlignment{
		Vertical:        VAlignBottom,
		Horizontal:      HAlignFill,
		TextRotation:    90,
		WrapText:        true,
		Indent:          10,
		RelativeIndent:  5,
		JustifyLastLine: true,
		ShrinkToFit:     true,
		ReadingOrder:    4,
	}, alignment)
}

func TestStyleFormat_Settings_Border(t *testing.T) {
	style := New(
		Border.Type(BorderStyleDashDot),
		Border.Color("#FF00FF"),
		Border.Diagonal.Type(BorderStyleDashDot),
		Border.Diagonal.Color("#FF00FF"),
		Border.Vertical.Type(BorderStyleDashDot),
		Border.Vertical.Color("#FF00FF"),
		Border.Horizontal.Type(BorderStyleDashDot),
		Border.Horizontal.Color("#FF00FF"),
	)

	font, fill, alignment, number, protection, border, namedInfo := from(style)
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.NotNil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.Border{
		Left: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Top: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Bottom: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Right: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Diagonal: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Vertical: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Horizontal: &ml.BorderSegment{
			Type:  BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
	}, border)
}

func TestStyleFormat_Settings_Fill(t *testing.T) {
	//pattern fill settings present
	style := New(
		Fill.Color("#FF00FF"),
		Fill.Background("#00FF00"),
		Fill.Type(PatternTypeDarkDown),
	)
	font, fill, alignment, number, protection, border, namedInfo := from(style)
	require.Nil(t, font)
	require.NotNil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.Fill{
		Pattern: &ml.PatternFill{
			Color:      color.New("FFFF00FF"),
			Background: color.New("FF00FF00"),
			Type:       PatternTypeDarkDown,
		},
	}, fill)

	//gradient fill settings present
	style.Set(
		Fill.Gradient.Degree(90),
		Fill.Gradient.Type(GradientTypePath),
		Fill.Gradient.Left(1),
		Fill.Gradient.Right(2),
		Fill.Gradient.Top(3),
		Fill.Gradient.Bottom(4),
		Fill.Gradient.Stop(100, "#FF00FF"),
		Fill.Gradient.Stop(200, "#00FF00"),
	)
	font, fill, alignment, number, protection, border, namedInfo = from(style)
	require.Nil(t, font)
	require.NotNil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.Fill{
		Gradient: &ml.GradientFill{
			Degree: 90,
			Type:   GradientTypePath,
			Left:   1,
			Right:  2,
			Top:    3,
			Bottom: 4,
			Stop: []*ml.GradientStop{
				{Position: 100, Color: color.New("FFFF00FF")},
				{Position: 200, Color: color.New("FF00FF00")},
			},
		},
	}, fill)

}

func TestStyleFormat_Settings_Font(t *testing.T) {
	style := New(
		Font.Name("Calibri"),
		Font.Size(10),
		Font.Bold,
		Font.Italic,
		Font.Strikeout,
		Font.Shadow,
		Font.Condense,
		Font.Extend,
		Font.Family(FontFamilyDecorative),
		Font.Color("#FF00FF"),
		Font.Underline(UnderlineTypeSingle),
		Font.VAlign(FontVAlignBaseline),
		Font.Scheme(FontSchemeMinor),
	)

	font, fill, alignment, number, protection, border, namedInfo := from(style)
	require.NotNil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.Font{
		Name:      "Calibri",
		Bold:      true,
		Italic:    true,
		Strike:    true,
		Shadow:    true,
		Condense:  true,
		Extend:    true,
		Size:      10.0,
		Color:     color.New("FFFF00FF"),
		Family:    FontFamilyDecorative,
		Underline: UnderlineTypeSingle,
		VAlign:    FontVAlignBaseline,
		Scheme:    FontSchemeMinor,
	}, font)
}

func TestStyleFormat_Settings_Number(t *testing.T) {
	style := New(
		NumberFormatID(8),
	)
	font, fill, alignment, number, protection, border, namedInfo := from(style)
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.NotNil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.NumberFormat{
		ID:   8,
		Code: "($#,##0.00_);[RED]($#,##0.00_)",
	}, number)
}

func TestStyleFormat_Settings_Protection(t *testing.T) {
	style := New(
		Protection.Hidden,
		Protection.Locked,
	)
	font, fill, alignment, number, protection, border, namedInfo := from(style)
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.NotNil(t, protection)
	require.Nil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.CellProtection{
		Locked: true,
		Hidden: true,
	}, protection)
}

func TestFontMarshal(t *testing.T) {
	//0 must be omitted
	font, _, _, _, _, _, _ := from(New(
		Font.Size(0),
		Font.Family(0),
		Font.Charset(0),
	))
	require.Nil(t, font)

	//simple version
	font, _, _, _, _, _, _ = from(New(
		Font.Name("Calibri"),
	))
	encoded, _ := xml.Marshal(font)
	require.Equal(t, `<Font><name val="Calibri"></name></Font>`, string(encoded))

	//full version
	font, _, _, _, _, _, _ = from(New(
		Font.Name("Calibri"),
		Font.Size(10),
		Font.Bold,
		Font.Italic,
		Font.Strikeout,
		Font.Shadow,
		Font.Condense,
		Font.Extend,
		Font.Family(FontFamilyDecorative),
		Font.Color("#FF00FF"),
		Font.Underline(UnderlineTypeSingle),
		Font.VAlign(FontVAlignBaseline),
		Font.Scheme(FontSchemeMinor),
	))

	encoded, _ = xml.Marshal(font)
	require.Equal(t, `<Font><name val="Calibri"></name><family val="5"></family><b val="true"></b><i val="true"></i><strike val="true"></strike><shadow val="true"></shadow><condense val="true"></condense><extend val="true"></extend><color indexed="6"></color><sz val="10"></sz><u val="single"></u><vertAlign val="baseline"></vertAlign><scheme val="minor"></scheme></Font>`, string(encoded))
}

func TestFillMarshal(t *testing.T) {
	//0 must be omitted
	_, fill, _, _, _, _, _ := from(New())
	require.Nil(t, fill)

	//pattern version
	_, fill, _, _, _, _, _ = from(New(
		Fill.Color("#FF00FF"),
		Fill.Background("#00FF00"),
		Fill.Type(PatternTypeDarkDown),
	))
	encoded, _ := xml.Marshal(fill)
	require.Equal(t, `<Fill><patternFill patternType="darkDown"><fgColor indexed="6"></fgColor><bgColor indexed="3"></bgColor></patternFill></Fill>`, string(encoded))

	//gradient version
	_, fill, _, _, _, _, _ = from(New(
		Fill.Gradient.Degree(90),
		Fill.Gradient.Type(GradientTypePath),
		Fill.Gradient.Left(1),
		Fill.Gradient.Right(2),
		Fill.Gradient.Top(3),
		Fill.Gradient.Bottom(4),
		Fill.Gradient.Stop(100, "#FF00FF"),
		Fill.Gradient.Stop(200, "#00FF00"),
	))
	encoded, _ = xml.Marshal(fill)
	require.Equal(t, `<Fill><gradientFill degree="90" left="1" right="2" top="3" bottom="4" type="path"><stop position="100"><color indexed="6"></color></stop><stop position="200"><color indexed="3"></color></stop></gradientFill></Fill>`, string(encoded))
}

func TestBorderMarshal(t *testing.T) {
	//0 must be omitted
	_, _, _, _, _, border, _ := from(New())
	require.Nil(t, border)

	//simple version
	_, _, _, _, _, border, _ = from(New(
		Border.Outline,
	))
	encoded, _ := xml.Marshal(border)
	require.Equal(t, `<Border outline="true"></Border>`, string(encoded))

	//full version
	_, _, _, _, _, border, _ = from(New(
		Border.Outline,
		Border.DiagonalUp,
		Border.DiagonalDown,
		Border.Left.Type(BorderStyleDashDot),
		Border.Left.Color("#FF00FF"),
		Border.Right.Type(BorderStyleDashDot),
		Border.Right.Color("#FF00FF"),
		Border.Top.Type(BorderStyleDashDot),
		Border.Top.Color("#FF00FF"),
		Border.Bottom.Type(BorderStyleDashDot),
		Border.Bottom.Color("#FF00FF"),
		Border.Diagonal.Type(BorderStyleDashDot),
		Border.Diagonal.Color("#FF00FF"),
		Border.Vertical.Type(BorderStyleDashDot),
		Border.Vertical.Color("#FF00FF"),
		Border.Horizontal.Type(BorderStyleDashDot),
		Border.Horizontal.Color("#FF00FF"),
	))
	encoded, _ = xml.Marshal(border)
	require.Equal(t, `<Border diagonalUp="true" diagonalDown="true" outline="true"><left style="dashDot"><color indexed="6"></color></left><right style="dashDot"><color indexed="6"></color></right><top style="dashDot"><color indexed="6"></color></top><bottom style="dashDot"><color indexed="6"></color></bottom><diagonal style="dashDot"><color indexed="6"></color></diagonal><vertical style="dashDot"><color indexed="6"></color></vertical><horizontal style="dashDot"><color indexed="6"></color></horizontal></Border>`, string(encoded))
}
