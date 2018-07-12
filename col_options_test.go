package xlsx

import (
	"github.com/plandem/xlsx/column_option"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColumnOptions(t *testing.T) {
	options := NewColumnOptions(
		columnOption.OutlineLevel(5),
		columnOption.Hidden(true),
		columnOption.Phonetic(true),
		columnOption.Formatting(12345),
		columnOption.Width(45.5),
	)

	require.IsType(t, &columnOption.ColumnOptions{}, options)
	require.Equal(t, &columnOption.ColumnOptions{
		OutlineLevel: 5,
		Hidden:       true,
		Phonetic:     true,
		StyleID:      12345,
		Width:        45.5,
	}, options)

	options.Set(columnOption.Width(50))
	require.Equal(t, &columnOption.ColumnOptions{
		OutlineLevel: 5,
		Hidden:       true,
		Phonetic:     true,
		StyleID:      12345,
		Width:        50,
	}, options)
}
