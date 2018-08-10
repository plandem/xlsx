package helpers_test

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/helpers"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
	"encoding/xml"
)

func TestStyleFormat_Settings(t *testing.T) {
	style := format.New()

	//empty
	font, fill, alignment, number, protection, border, namedInfo := helpers.FromStyleFormat(style)
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)
	require.Nil(t, namedInfo)

	//full featured
	style.Set(
		format.Alignment.VAlign(format.VAlignBottom),
		format.Alignment.HAlign(format.HAlignFill),
		format.Alignment.TextRotation(90),
		format.Alignment.WrapText,
		format.Alignment.Indent(10),
		format.Alignment.RelativeIndent(5),
		format.Alignment.JustifyLastLine,
		format.Alignment.ShrinkToFit,
		format.Alignment.ReadingOrder(4),
		format.Border.Type(format.BorderStyleDashDot),
		format.Border.Color("#FF00FF"),
		format.Border.Diagonal.Type(format.BorderStyleDashDot),
		format.Border.Diagonal.Color("#FF00FF"),
		format.Border.Vertical.Type(format.BorderStyleDashDot),
		format.Border.Vertical.Color("#FF00FF"),
		format.Border.Horizontal.Type(format.BorderStyleDashDot),
		format.Border.Horizontal.Color("#FF00FF"),
		format.Fill.Color("#FF00FF"),
		format.Fill.Background("#00FF00"),
		format.Fill.Type(format.PatternTypeDarkDown),
		format.Font.Name("Calibri"),
		format.Font.Size(10),
		format.Font.Bold,
		format.Font.Italic,
		format.Font.Strikeout,
		format.Font.Shadow,
		format.Font.Condense,
		format.Font.Extend,
		format.Font.Family(format.FontFamilyDecorative),
		format.Font.Color("#FF00FF"),
		format.Font.Underline(format.UnderlineTypeSingle),
		format.Font.VAlign(format.FontVAlignBaseline),
		format.Font.Scheme(format.FontSchemeMinor),
		format.NumberFormatID(8),
		format.Protection.Hidden,
		format.Protection.Locked,
	)

	font, fill, alignment, number, protection, border, namedInfo = helpers.FromStyleFormat(style)
	require.NotNil(t, font)
	require.NotNil(t, fill)
	require.NotNil(t, alignment)
	require.NotNil(t, number)
	require.NotNil(t, protection)
	require.NotNil(t, border)

	require.Equal(t, &ml.CellAlignment{
		Vertical:        format.VAlignBottom,
		Horizontal:      format.HAlignFill,
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
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Top: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Bottom: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Right: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Diagonal: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Vertical: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Horizontal: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
	}, border)

	require.Equal(t, &ml.Fill{
		Pattern: &ml.PatternFill{
			Color:      color.New("FFFF00FF"),
			Background: color.New("FF00FF00"),
			Type:       format.PatternTypeDarkDown,
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
		Family:    format.FontFamilyDecorative,
		Underline: format.UnderlineTypeSingle,
		VAlign:    format.FontVAlignBaseline,
		Scheme:    format.FontSchemeMinor,
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
	style := format.New(
		format.Alignment.VAlign(format.VAlignBottom),
		format.Alignment.HAlign(format.HAlignFill),
		format.Alignment.TextRotation(90),
		format.Alignment.WrapText,
		format.Alignment.Indent(10),
		format.Alignment.RelativeIndent(5),
		format.Alignment.JustifyLastLine,
		format.Alignment.ShrinkToFit,
		format.Alignment.ReadingOrder(4),
	)

	font, fill, alignment, number, protection, border, namedInfo := helpers.FromStyleFormat(style)
	require.Nil(t, font)
	require.Nil(t, fill)
	require.NotNil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.Nil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.CellAlignment{
		Vertical:        format.VAlignBottom,
		Horizontal:      format.HAlignFill,
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
	style := format.New(
		format.Border.Type(format.BorderStyleDashDot),
		format.Border.Color("#FF00FF"),
		format.Border.Diagonal.Type(format.BorderStyleDashDot),
		format.Border.Diagonal.Color("#FF00FF"),
		format.Border.Vertical.Type(format.BorderStyleDashDot),
		format.Border.Vertical.Color("#FF00FF"),
		format.Border.Horizontal.Type(format.BorderStyleDashDot),
		format.Border.Horizontal.Color("#FF00FF"),
	)

	font, fill, alignment, number, protection, border, namedInfo := helpers.FromStyleFormat(style)
	require.Nil(t, font)
	require.Nil(t, fill)
	require.Nil(t, alignment)
	require.Nil(t, number)
	require.Nil(t, protection)
	require.NotNil(t, border)
	require.Nil(t, namedInfo)

	require.Equal(t, &ml.Border{
		Left: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Top: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Bottom: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Right: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Diagonal: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Vertical: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
		Horizontal: &ml.BorderSegment{
			Type:  format.BorderStyleDashDot,
			Color: color.New("#FF00FF"),
		},
	}, border)
}

func TestStyleFormat_Settings_Fill(t *testing.T) {
	//pattern fill settings present
	style := format.New(
		format.Fill.Color("#FF00FF"),
		format.Fill.Background("#00FF00"),
		format.Fill.Type(format.PatternTypeDarkDown),
	)
	font, fill, alignment, number, protection, border, namedInfo := helpers.FromStyleFormat(style)
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
			Type:       format.PatternTypeDarkDown,
		},
	}, fill)

	//gradient fill settings present
	style.Set(
		format.Fill.Gradient.Degree(90),
		format.Fill.Gradient.Type(format.GradientTypePath),
		format.Fill.Gradient.Left(1),
		format.Fill.Gradient.Right(2),
		format.Fill.Gradient.Top(3),
		format.Fill.Gradient.Bottom(4),
		format.Fill.Gradient.Stop(100, "#FF00FF"),
		format.Fill.Gradient.Stop(200, "#00FF00"),
	)
	font, fill, alignment, number, protection, border, namedInfo = helpers.FromStyleFormat(style)
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
			Type:   format.GradientTypePath,
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
	style := format.New(
		format.Font.Name("Calibri"),
		format.Font.Size(10),
		format.Font.Bold,
		format.Font.Italic,
		format.Font.Strikeout,
		format.Font.Shadow,
		format.Font.Condense,
		format.Font.Extend,
		format.Font.Family(format.FontFamilyDecorative),
		format.Font.Color("#FF00FF"),
		format.Font.Underline(format.UnderlineTypeSingle),
		format.Font.VAlign(format.FontVAlignBaseline),
		format.Font.Scheme(format.FontSchemeMinor),
	)

	font, fill, alignment, number, protection, border, namedInfo := helpers.FromStyleFormat(style)
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
		Family:    format.FontFamilyDecorative,
		Underline: format.UnderlineTypeSingle,
		VAlign:    format.FontVAlignBaseline,
		Scheme:    format.FontSchemeMinor,
	}, font)
}

func TestStyleFormat_Settings_Number(t *testing.T) {
	style := format.New(
		format.NumberFormatID(8),
	)
	font, fill, alignment, number, protection, border, namedInfo := helpers.FromStyleFormat(style)
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
	style := format.New(
		format.Protection.Hidden,
		format.Protection.Locked,
	)
	font, fill, alignment, number, protection, border, namedInfo := helpers.FromStyleFormat(style)
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
	font, _,_,_,_,_,_ := helpers.FromStyleFormat(format.New(
		format.Font.Size(0),
		format.Font.Family(0),
		format.Font.Charset(0),
	))
	require.Nil(t, font)

	//simple version
	font, _,_,_,_,_,_ = helpers.FromStyleFormat(format.New(
		format.Font.Name("Calibri"),
	))
	encoded, _ := xml.Marshal(font)
	require.Equal(t, `<Font><name val="Calibri"></name></Font>`, string(encoded))

	//full version
	font, _,_,_,_,_,_ = helpers.FromStyleFormat(format.New(
		format.Font.Name("Calibri"),
		format.Font.Size(10),
		format.Font.Bold,
		format.Font.Italic,
		format.Font.Strikeout,
		format.Font.Shadow,
		format.Font.Condense,
		format.Font.Extend,
		format.Font.Family(format.FontFamilyDecorative),
		format.Font.Color("#FF00FF"),
		format.Font.Underline(format.UnderlineTypeSingle),
		format.Font.VAlign(format.FontVAlignBaseline),
		format.Font.Scheme(format.FontSchemeMinor),
	))

	encoded, _ = xml.Marshal(font)
	require.Equal(t, `<Font><name val="Calibri"></name><family val="5"></family><b val="true"></b><i val="true"></i><strike val="true"></strike><shadow val="true"></shadow><condense val="true"></condense><extend val="true"></extend><color indexed="6"></color><sz val="10"></sz><u val="single"></u><vertAlign val="baseline"></vertAlign><scheme val="minor"></scheme></Font>`, string(encoded))
}

func TestFillMarshal(t *testing.T) {
	//0 must be omitted
	 _, fill,_,_,_,_,_ := helpers.FromStyleFormat(format.New())
	require.Nil(t, fill)

	//pattern version
	_, fill,_,_,_,_,_ = helpers.FromStyleFormat(format.New(
		format.Fill.Color("#FF00FF"),
		format.Fill.Background("#00FF00"),
		format.Fill.Type(format.PatternTypeDarkDown),
	))
	encoded, _ := xml.Marshal(fill)
	require.Equal(t, `<Fill><patternFill patternType="darkDown"><fgColor indexed="6"></fgColor><bgColor indexed="3"></bgColor></patternFill></Fill>`, string(encoded))

	//gradient version
	_, fill,_,_,_,_,_ = helpers.FromStyleFormat(format.New(
		format.Fill.Gradient.Degree(90),
		format.Fill.Gradient.Type(format.GradientTypePath),
		format.Fill.Gradient.Left(1),
		format.Fill.Gradient.Right(2),
		format.Fill.Gradient.Top(3),
		format.Fill.Gradient.Bottom(4),
		format.Fill.Gradient.Stop(100, "#FF00FF"),
		format.Fill.Gradient.Stop(200, "#00FF00"),
	))
	encoded, _ = xml.Marshal(fill)
	require.Equal(t, `<Fill><gradientFill degree="90" left="1" right="2" top="3" bottom="4" type="path"><stop position="100"><color indexed="6"></color></stop><stop position="200"><color indexed="3"></color></stop></gradientFill></Fill>`, string(encoded))
}

func TestBorderMarshal(t *testing.T) {
	//0 must be omitted
	_, _,_,_,_,border,_ := helpers.FromStyleFormat(format.New())
	require.Nil(t, border)

	//simple version
	_, _,_,_,_,border,_ = helpers.FromStyleFormat(format.New(
		format.Border.Outline,
	))
	encoded, _ := xml.Marshal(border)
	require.Equal(t, `<Border outline="true"></Border>`, string(encoded))

	//full version
	_, _,_,_,_,border,_ = helpers.FromStyleFormat(format.New(
		format.Border.Outline,
		format.Border.DiagonalUp,
		format.Border.DiagonalDown,
		format.Border.Left.Type(format.BorderStyleDashDot),
		format.Border.Left.Color("#FF00FF"),
		format.Border.Right.Type(format.BorderStyleDashDot),
		format.Border.Right.Color("#FF00FF"),
		format.Border.Top.Type(format.BorderStyleDashDot),
		format.Border.Top.Color("#FF00FF"),
		format.Border.Bottom.Type(format.BorderStyleDashDot),
		format.Border.Bottom.Color("#FF00FF"),
		format.Border.Diagonal.Type(format.BorderStyleDashDot),
		format.Border.Diagonal.Color("#FF00FF"),
		format.Border.Vertical.Type(format.BorderStyleDashDot),
		format.Border.Vertical.Color("#FF00FF"),
		format.Border.Horizontal.Type(format.BorderStyleDashDot),
		format.Border.Horizontal.Color("#FF00FF"),
	))
	encoded, _ = xml.Marshal(border)
	require.Equal(t, `<Border diagonalUp="true" diagonalDown="true" outline="true"><left style="dashDot"><color indexed="6"></color></left><right style="dashDot"><color indexed="6"></color></right><top style="dashDot"><color indexed="6"></color></top><bottom style="dashDot"><color indexed="6"></color></bottom><diagonal style="dashDot"><color indexed="6"></color></diagonal><vertical style="dashDot"><color indexed="6"></color></vertical><horizontal style="dashDot"><color indexed="6"></color></horizontal></Border>`, string(encoded))
}
