package xlsx

import (
	"github.com/plandem/xlsx/internal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColumnOptions(t *testing.T) {
	options := NewColumnOptions(
		ColumnOption.OutlineLevel(5),
		ColumnOption.Hidden(true),
		ColumnOption.Phonetic(true),
		ColumnOption.Formatting(12345),
		ColumnOption.Width(45.5),
	)

	require.IsType(t, &internal.ColumnOptions{}, options)
	require.Equal(t, &internal.ColumnOptions{
		OutlineLevel: 5,
		Hidden:       true,
		Phonetic:     true,
		StyleID:      12345,
		Width:        45.5,
	}, options)
}
