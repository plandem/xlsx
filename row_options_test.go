package xlsx

import (
	"github.com/plandem/xlsx/row_option"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRowOptions(t *testing.T) {
	options := NewRowOptions(
		rowOption.OutlineLevel(5),
		rowOption.Hidden(true),
		rowOption.Phonetic(true),
		rowOption.Formatting(12345),
		rowOption.Height(45.5),
	)

	require.IsType(t, &rowOption.Options{}, options)
	require.Equal(t, &rowOption.Options{
		OutlineLevel: 5,
		Hidden:       true,
		Phonetic:     true,
		StyleID:      12345,
		Height:       45.5,
	}, options)

	options.Set(rowOption.Height(50))
	require.Equal(t, &rowOption.Options{
		OutlineLevel: 5,
		Hidden:       true,
		Phonetic:     true,
		StyleID:      12345,
		Height:       50,
	}, options)
}
