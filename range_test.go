package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRange(t *testing.T) {
	xl, err := Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()
	sheet := xl.Sheet(0)
	r := sheet.Range("D10:E10")
	require.Equal(t, []string{"1", "2"}, r.Values())
	require.Equal(t, format.StyleID(0), sheet.CellByRef("D10").ml.Style)
	require.Equal(t, format.StyleID(0), sheet.CellByRef("E10").ml.Style)

	//test styles
	style := format.New(
		format.Font.Name("Calibri"),
		format.Font.Size(12),
	)

	styleRef := xl.AddFormatting(style)
	r.SetFormatting(styleRef)

	require.Equal(t, format.StyleID(styleRef), sheet.CellByRef("D10").ml.Style)
	require.Equal(t, format.StyleID(styleRef), sheet.CellByRef("E10").ml.Style)

	r.Clear()
	require.Equal(t, []string{"", ""}, r.Values())
	require.Equal(t, format.StyleID(styleRef), sheet.CellByRef("D10").ml.Style)
	require.Equal(t, format.StyleID(styleRef), sheet.CellByRef("E10").ml.Style)

	r.Reset()
	require.Equal(t, []string{"", ""}, r.Values())
	require.Equal(t, format.StyleID(0), sheet.CellByRef("D10").ml.Style)
	require.Equal(t, format.StyleID(0), sheet.CellByRef("E10").ml.Style)
}
