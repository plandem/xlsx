package format

import (
	"encoding/xml"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFill(t *testing.T) {
	//pattern only
	style := New(
		Fill.Pattern.Type(PatternTypeDarkDown),
		Fill.Pattern.Color("#FFFFFF"),
		Fill.Pattern.Background("#FF0000"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, &StyleFormat{
		fill: ml.Fill{
			Pattern: &ml.PatternFill{
				Color:      color.New("FFFFFFFF"),
				Background: color.New("FFFF0000"),
				Type:       PatternTypeDarkDown,
			},
		},
	}, style)

	//gradient only
	style = New(
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
	require.Equal(t, &StyleFormat{
		fill: ml.Fill{
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
		},
	}, style)

	//pattern overriden by gradient
	style = New(
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
	require.Equal(t, &StyleFormat{
		fill: ml.Fill{
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
		},
	}, style)

	//gradient overriden by pattern
	style = New(
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
	require.Equal(t, &StyleFormat{
		fill: ml.Fill{
			Pattern: &ml.PatternFill{
				Color:      color.New("FFFFFFFF"),
				Background: color.New("FFFF0000"),
				Type:       PatternTypeDarkDown,
			},
		},
	}, style)
}

func TestFillMarshal(t *testing.T) {
	//0 must be omitted
	style := New()
	encoded, err := xml.Marshal(&style.fill)
	require.Empty(t, err)
	require.Equal(t, `<Fill></Fill>`, string(encoded))

	//pattern version
	style = New(
		Fill.Color("#FF00FF"),
		Fill.Background("#00FF00"),
		Fill.Type(PatternTypeDarkDown),
	)
	encoded, _ = xml.Marshal(&style.fill)
	require.Equal(t, `<Fill><patternFill patternType="darkDown"><fgColor indexed="6"></fgColor><bgColor indexed="3"></bgColor></patternFill></Fill>`, string(encoded))

	//gradient version
	style = New(
		Fill.Gradient.Degree(90),
		Fill.Gradient.Type(GradientTypePath),
		Fill.Gradient.Left(1),
		Fill.Gradient.Right(2),
		Fill.Gradient.Top(3),
		Fill.Gradient.Bottom(4),
		Fill.Gradient.Stop(100, "#FF00FF"),
		Fill.Gradient.Stop(200, "#00FF00"),
	)
	encoded, _ = xml.Marshal(&style.fill)
	require.Equal(t, `<Fill><gradientFill degree="90" left="1" right="2" top="3" bottom="4" type="path"><stop position="100"><color indexed="6"></color></stop><stop position="200"><color indexed="3"></color></stop></gradientFill></Fill>`, string(encoded))
}
