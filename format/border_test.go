package format

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBorder(t *testing.T) {
	style := NewStyles(
		Border.Type(BorderStyleDashDot),
		Border.Color("#FF00FF"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, createAndFill(func(f *StyleFormat) {
		f.styleInfo.Border = &ml.Border{
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
			Diagonal:   &ml.BorderSegment{},
			Vertical:   &ml.BorderSegment{},
			Horizontal: &ml.BorderSegment{},
		}
	}), style)

	style = NewStyles(
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
	)

	require.Equal(t, createAndFill(func(f *StyleFormat) {
		f.styleInfo.Border = &ml.Border{
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
		}
	}), style)
}
