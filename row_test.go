package xlsx

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/types/options"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRow(t *testing.T) {
	xl, err := Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()
	sheet := xl.Sheet(0)
	r := sheet.Row(5)

	o := options.NewRowOptions(
		options.Row.Height(0),
		options.Row.OutlineLevel(10),
		options.Row.Hidden(true),
		options.Row.Phonetic(true),
		options.Row.Collapsed(true),
	)

	r.Set(o)

	require.Equal(t, r.ml.OutlineLevel, o.OutlineLevel)
	require.Equal(t, r.ml.Hidden, o.Hidden)
	require.Equal(t, r.ml.Phonetic, o.Phonetic)
	require.Equal(t, r.ml.Collapsed, o.Collapsed)
	require.Equal(t, r.ml.Height, float32(0.0))
	require.Equal(t, r.ml.CustomHeight, false)
	require.Equal(t, r.ml.CustomFormat, false)
	require.Equal(t, r.ml.Style, styles.DirectStyleID(0))

	o = options.NewRowOptions(
		options.Row.Height(100.0),
	)

	r.Set(o)
	require.Equal(t, r.ml.OutlineLevel, o.OutlineLevel)
	require.Equal(t, r.ml.Hidden, o.Hidden)
	require.Equal(t, r.ml.Phonetic, o.Phonetic)
	require.Equal(t, r.ml.Collapsed, o.Collapsed)
	require.Equal(t, r.ml.Height, float32(100.0))
	require.Equal(t, r.ml.CustomHeight, true)
	require.Equal(t, r.ml.CustomFormat, false)
	require.Equal(t, r.ml.Style, styles.DirectStyleID(0))

	style := styles.New(
		styles.Font.Name("Calibri"),
		styles.Font.Size(12),
	)

	styleRef := xl.AddStyles(style)
	r.SetStyles(styleRef)

	require.Equal(t, r.ml.CustomFormat, true)
	require.Equal(t, r.ml.Style, styles.DirectStyleID(styleRef))
}
