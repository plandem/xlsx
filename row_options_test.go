package xlsx

import (
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

	require.IsType(t, &rowOptions{}, options)
	require.Equal(t, &rowOptions{
		outlineLevel: 5,
		hidden:       true,
		phonetic:     true,
		styleID:      12345,
		height:       45.5,
	}, options)

	options.Set(RowOption.Height(50))
	require.Equal(t, &rowOptions{
		outlineLevel: 5,
		hidden:       true,
		phonetic:     true,
		styleID:      12345,
		height:       50,
	}, options)
}
