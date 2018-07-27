package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/options"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCol(t *testing.T) {
	xl, err := Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()
	sheet := xl.Sheet(0)
	c := sheet.Col(5)

	o := options.NewColumnOptions(
		options.Column.Width(0),
		options.Column.OutlineLevel(10),
		options.Column.Hidden(true),
		options.Column.Phonetic(true),
		options.Column.Collapsed(true),
	)

	c.Set(o)

	require.Equal(t, c.ml.OutlineLevel, o.OutlineLevel)
	require.Equal(t, c.ml.Hidden, o.Hidden)
	require.Equal(t, c.ml.Phonetic, o.Phonetic)
	require.Equal(t, c.ml.Collapsed, o.Collapsed)
	require.Equal(t, c.ml.Width, float32(0.0))
	require.Equal(t, c.ml.CustomWidth, false)
	require.Equal(t, c.ml.Style, ml.StyleRefID(0))

	o = options.NewColumnOptions(
		options.Column.Width(100.0),
	)

	c.Set(o)
	require.Equal(t, c.ml.OutlineLevel, o.OutlineLevel)
	require.Equal(t, c.ml.Hidden, o.Hidden)
	require.Equal(t, c.ml.Phonetic, o.Phonetic)
	require.Equal(t, c.ml.Collapsed, o.Collapsed)
	require.Equal(t, c.ml.Width, float32(100.0))
	require.Equal(t, c.ml.CustomWidth, true)
	require.Equal(t, c.ml.Style, ml.StyleRefID(0))

	style := format.New(
		format.Font.Name("Calibri"),
		format.Font.Size(12),
	)

	styleRef := xl.AddFormatting(style)
	c.SetFormatting(styleRef)

	require.Equal(t, c.ml.Style, ml.StyleRefID(styleRef))
}
