package format

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberFormat(t *testing.T) {
	style := New(
		NumberFormatID(8),
		//NumberFormat("#.### usd"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, &StyleFormat{
		key: "b828191babe328f056db7e2fd933ee1f",
		fill: ml.Fill{
			Pattern:  &ml.PatternFill{},
			Gradient: &ml.GradientFill{},
		},
		border: ml.Border{
			Left:       &ml.BorderSegment{},
			Right:      &ml.BorderSegment{},
			Top:        &ml.BorderSegment{},
			Bottom:     &ml.BorderSegment{},
			Diagonal:   &ml.BorderSegment{},
			Vertical:   &ml.BorderSegment{},
			Horizontal: &ml.BorderSegment{},
		},
		numFormat: ml.NumberFormat{
			ID:   8,
			Code: "($#,##0.00_);[RED]($#,##0.00_)",
		},
	}, style)

	style = New(
		NumberFormat("#.### usd"),
	)

	require.Equal(t, &StyleFormat{
		key: "c002d336f2642f8c69a240eb35011015",
		fill: ml.Fill{
			Pattern:  &ml.PatternFill{},
			Gradient: &ml.GradientFill{},
		},
		border: ml.Border{
			Left:       &ml.BorderSegment{},
			Right:      &ml.BorderSegment{},
			Top:        &ml.BorderSegment{},
			Bottom:     &ml.BorderSegment{},
			Diagonal:   &ml.BorderSegment{},
			Vertical:   &ml.BorderSegment{},
			Horizontal: &ml.BorderSegment{},
		},
		numFormat: ml.NumberFormat{
			ID:   -1,
			Code: "#.### usd",
		},
	}, style)
}
