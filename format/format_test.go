package format

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStyleFormat_Settings(t *testing.T) {
	style := New()

	//empty
	font, fill, alignment, number, protection, border := style.Settings()
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)

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

	font, fill, alignment, number, protection, border = style.Settings()
	require.NotNil(t, font)
	require.NotNil(t, fill)
	require.NotNil(t, alignment)
	require.NotNil(t, number)
	require.NotNil(t, protection)
	require.NotNil(t, border)

	require.NotEqual(t, font, style.font)
	require.NotEqual(t, fill, style.fill)
	require.NotEqual(t, alignment, style.alignment)
	require.NotEqual(t, number, style.numFormat)
	require.NotEqual(t, protection, style.protection)
	require.NotEqual(t, border, style.border)

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

	font, fill, alignment, number, protection, border := style.Settings()
	require.Nil(t, font)
	require.Nil(t, fill)
	require.NotNil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)

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

	font, fill, alignment, number, protection, border := style.Settings()
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.NotNil(t, border)

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
	font, fill, alignment, number, protection, border := style.Settings()
	require.Nil(t, font)
	require.NotNil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)

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
	font, fill, alignment, number, protection, border = style.Settings()
	require.Nil(t, font)
	require.NotNil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)

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

	font, fill, alignment, number, protection, border := style.Settings()
	require.NotNil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)

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
	font, fill, alignment, number, protection, border := style.Settings()
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.NotNil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)

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
	font, fill, alignment, number, protection, border := style.Settings()
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.NotNil(t, protection)
	require.Nil(t, border)

	require.Equal(t, &ml.CellProtection{
		Locked: true,
		Hidden: true,
	}, protection)
}
