package format

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProtection(t *testing.T) {
	style := New(
		Protection.Hidden,
		Protection.Locked,
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, &StyleFormat{
		key: "0b65d5dc27be42d7f4826817e6650162",
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
		Protection: ml.CellProtection{
			Locked: true,
			Hidden: true,
		},
	}, style)

}
