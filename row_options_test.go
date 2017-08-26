package xlsx

import (
	"github.com/plandem/xlsx/internal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRowOptions(t *testing.T) {
	options := NewRowOptions(
		RowOption.OutlineLevel(5),
		RowOption.Hidden(true),
		RowOption.Phonetic(true),
		RowOption.Formatting(12345),
		RowOption.Height(45.5),
	)

	require.IsType(t, &internal.RowOptions{}, options)
	require.Equal(t, &internal.RowOptions{
		OutlineLevel: 5,
		Hidden:       true,
		Phonetic:     true,
		StyleID:      12345,
		Height:       45.5,
	}, options)
}
