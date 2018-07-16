package options

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColumnOptions(t *testing.T) {
	o := NewColumnOptions(
		Column.OutlineLevel(5),
		Column.Hidden(true),
		Column.Phonetic(true),
		Column.Width(45.5),
	)

	require.IsType(t, &ColumnOptions{}, o)
	require.Equal(t, &ColumnOptions{
		OutlineLevel: 5,
		Hidden:       true,
		Phonetic:     true,
		Width:        45.5,
	}, o)

	o = NewColumnOptions(
		Column.OutlineLevel(10),
		Column.Hidden(false),
		Column.Phonetic(true),
	)
	require.Equal(t, &ColumnOptions{
		Hidden:       false,
		Phonetic:     true,
	}, o)
}
