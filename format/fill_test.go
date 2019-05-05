package format

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFill(t *testing.T) {
	//pattern only
	style := NewStyles(
		Fill.Pattern.Type(PatternTypeDarkDown),
		Fill.Pattern.Color("#FFFFFF"),
		Fill.Pattern.Background("#FF0000"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, createStylesAndFill(func(f *StyleFormat) {
		f.styleInfo.Fill.Pattern = &ml.PatternFill{
			Color:      color.New("FFFFFFFF"),
			Background: color.New("FFFF0000"),
			Type:       PatternTypeDarkDown,
		}

	}), style)

	//gradient only
	style = NewStyles(
		Fill.Gradient.Degree(90),
		Fill.Gradient.Type(GradientTypePath),
		Fill.Gradient.Left(1),
		Fill.Gradient.Right(2),
		Fill.Gradient.Top(3),
		Fill.Gradient.Bottom(4),
		Fill.Gradient.Stop(100, "#FF00FF"),
		Fill.Gradient.Stop(200, "#00FF00"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, createStylesAndFill(func(f *StyleFormat) {
		f.styleInfo.Fill.Gradient = &ml.GradientFill{
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
		}
	}), style)

	//pattern override by gradient
	style = NewStyles(
		Fill.Pattern.Type(PatternTypeDarkDown),
		Fill.Pattern.Color("#FFFFFF"),
		Fill.Pattern.Background("#FF0000"),
		Fill.Gradient.Degree(90),
		Fill.Gradient.Type(GradientTypePath),
		Fill.Gradient.Left(1),
		Fill.Gradient.Right(2),
		Fill.Gradient.Top(3),
		Fill.Gradient.Bottom(4),
		Fill.Gradient.Stop(100, "#FF00FF"),
		Fill.Gradient.Stop(200, "#00FF00"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, createStylesAndFill(func(f *StyleFormat) {
		f.styleInfo.Fill.Gradient = &ml.GradientFill{
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
		}
	}), style)

	//gradient override by pattern
	style = NewStyles(
		Fill.Gradient.Degree(90),
		Fill.Gradient.Type(GradientTypePath),
		Fill.Gradient.Left(1),
		Fill.Gradient.Right(2),
		Fill.Gradient.Top(3),
		Fill.Gradient.Bottom(4),
		Fill.Gradient.Stop(100, "#FF00FF"),
		Fill.Gradient.Stop(200, "#00FF00"),
		Fill.Pattern.Type(PatternTypeDarkDown),
		Fill.Pattern.Color("#FFFFFF"),
		Fill.Pattern.Background("#FF0000"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, createStylesAndFill(func(f *StyleFormat) {
		f.styleInfo.Fill.Pattern = &ml.PatternFill{
			Color:      color.New("FFFFFFFF"),
			Background: color.New("FFFF0000"),
			Type:       PatternTypeDarkDown,
		}
	}), style)
}
