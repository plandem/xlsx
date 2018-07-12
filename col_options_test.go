package xlsx

import (
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

	require.IsType(t, &columnOptions{}, options)
	require.Equal(t, &columnOptions{
		outlineLevel: 5,
		hidden:       true,
		phonetic:     true,
		styleID:      12345,
		width:        45.5,
	}, options)

	options.Set(ColumnOption.Width(50))
	require.Equal(t, &columnOptions{
		outlineLevel: 5,
		hidden:       true,
		phonetic:     true,
		styleID:      12345,
		width:        50,
	}, options)
}
