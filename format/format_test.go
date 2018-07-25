package format

import (
	"github.com/stretchr/testify/require"
	"testing"
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
		NumberFormatID(0),
		NumberFormat("#.### usd"),
		Fill.Pattern.Type(PatternTypeDarkDown),
		Fill.Pattern.Color("#FFFFFF"),
		Fill.Pattern.Background("#FF0000"),
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

	//font, fill, alignment, number, protection, border := style.Pack()
	//require.Equal(t, )
}

