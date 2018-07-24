package format

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignment(t *testing.T) {
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

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, &StyleFormat{
		key: "31938baca968c10c008c70df782ec8a8",
		Fill: ml.Fill{
			Pattern:  &ml.PatternFill{},
			Gradient: &ml.GradientFill{},
		},
		Border: ml.Border{
			Left:       &ml.BorderSegment{},
			Right:      &ml.BorderSegment{},
			Top:        &ml.BorderSegment{},
			Bottom:     &ml.BorderSegment{},
			Diagonal:   &ml.BorderSegment{},
			Vertical:   &ml.BorderSegment{},
			Horizontal: &ml.BorderSegment{},
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
	}, style)

}
