package format

import (
	"github.com/stretchr/testify/require"
	"testing"
	"github.com/plandem/xlsx/internal/ml"
)

func TestFormat(t *testing.T) {
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

		Alignment.VAlign(VAlignBottom),
		Alignment.HAlign(HAlignFill),
		Alignment.TextRotation(90),
		Alignment.WrapText,
		Alignment.Indent(10),
		Alignment.RelativeIndent(5),
		Alignment.JustifyLastLine,
		Alignment.ShrinkToFit,
		Alignment.ReadingOrder(4),
		Protection.Hidden,
		Protection.Locked,
		NumberFormat(10, "#.### usd"),
		Fill.Type(PatternTypeDarkDown),
		Fill.Color("#FFFFFF"),
		Fill.Background("#FF0000"),
		Border.Type(BorderStyleDashDot),
		Border.Color("#FF00FF"),
		Border.Left.Type(BorderStyleDashDot),
		Border.Left.Color("#FF00FF"),
		Border.Right.Type(BorderStyleDashDot),
		Border.Right.Color("#FF00FF"),
		Border.Top.Type(BorderStyleDashDot),
		Border.Top.Color("#FF00FF"),
		Border.Bottom.Type(BorderStyleDashDot),
		Border.Bottom.Color("#FF00FF"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, &StyleFormat{
		key: "b1ed234f875f321a7a8812d8b06b5b10",
		Font: ml.Font{
			Name:      "Calibri",
			Bold:      true,
			Italic:    true,
			Strike:    true,
			Shadow:    true,
			Condense:  true,
			Extend:    true,
			Size:      10.0,
			Color:     newColor("FFFF00FF"),
			Family:    FontFamilyDecorative,
			Underline: UnderlineTypeSingle,
			VAlign:    FontVAlignBaseline,
			Scheme:    FontSchemeMinor,
		},
		Alignment: ml.CellAlignment{
			Vertical:        VAlignBottom,
			Horizontal:      HAlignFill,
			TextRotation:    90,
			WrapText:        true,
			Indent:          10,
			RelativeIndent:  5,
			JustifyLastLine: true,
			ShrinkToFit:     true,
			ReadingOrder:    4,
		},
		Protection: ml.CellProtection{
			Locked: true,
			Hidden: true,
		},
		NumFormat: ml.NumberFormat{
			10,
			"#.### usd",
		},
		Fill: ml.Fill{
			Pattern: &ml.PatternFill {
				Color:      newColor("FFFFFFFF"),
				Background: newColor("FFFF0000"),
				Type: PatternTypeDarkDown,
			},
		},
		Border: ml.Border{
			Left: &ml.BorderSegment{
				Type:  BorderStyleDashDot,
				Color: newColor("#FF00FF"),
			},
			Top: &ml.BorderSegment{
				Type:  BorderStyleDashDot,
				Color: newColor("#FF00FF"),
			},
			Bottom: &ml.BorderSegment{
				Type:  BorderStyleDashDot,
				Color: newColor("#FF00FF"),
			},
			Right: &ml.BorderSegment{
				Type:  BorderStyleDashDot,
				Color: newColor("#FF00FF"),
			},
		},
	}, style)

}
