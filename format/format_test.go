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
		Font.Strikeout,
		Alignment.VAlign(VAlignBottom),
		Alignment.HAlign(HAlignFill),
		Protection.Hidden,
		Protection.Locked,
		NumberFormat(10, "#.### usd"),
		Fill.Type(PatternTypeDarkDown),
		Fill.Color("#FFFFFF"),
		Fill.Background("#FF0000"),
		Border.Type(BorderStyleDashDot),
		Border.Left.Color("#FF00FF"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, &StyleFormat{
		key: "742e4d1372ee38e0d92d86accded1ce1",
		Font: font{
			Name:   "Calibri",
			Bold:   true,
			Strike: true,
			Size:   10.0,
		},
		Alignment: alignment{
			Vertical:   VAlignBottom,
			Horizontal: HAlignFill,
		},
		Protection: protection{
			Locked: true,
			Hidden: true,
		},
		NumFormat: numberFormat{
			10,
			"#.### usd",
		},
		Fill: fill{
			Color:      ColorToARGB("FFFFFFFF"),
			Background: ColorToARGB("FFFF0000"),
			Type:       PatternTypeDarkDown,
		},
		Border: border{
			Left: borderSegment{
				Type:  BorderStyleDashDot,
				Color: ColorToARGB("#FF00FF"),
			},
			Top: borderSegment{
				Type: BorderStyleDashDot,
			},
			Bottom: borderSegment{
				Type: BorderStyleDashDot,
			},
			Right: borderSegment{
				Type: BorderStyleDashDot,
			},
		},
	}, style)

}
