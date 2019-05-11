package options

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRowOptions(t *testing.T) {
	o := NewRowOptions(
		Row.OutlineLevel(5),
		Row.Hidden(true),
		Row.Phonetic(true),
		Row.Height(45.5),
		Row.Collapsed(true),
	)

	require.IsType(t, &RowOptions{}, o)
	require.Equal(t, &RowOptions{
		OutlineLevel: 5,
		Hidden:       true,
		Phonetic:     true,
		Height:       45.5,
		Collapsed:    true,
	}, o)

	o = NewRowOptions(
		Row.OutlineLevel(10),
		Row.Hidden(false),
		Row.Phonetic(true),
	)
	require.Equal(t, &RowOptions{
		Hidden:   false,
		Phonetic: true,
	}, o)
}
