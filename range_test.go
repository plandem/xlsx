package xlsx

import (
	"github.com/plandem/xlsx/format/styles"
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
	r := sheet.RangeByRef("D10:E10")
	require.Equal(t, []string{"1", "2"}, r.Values())
	require.Equal(t, styles.DirectStyleID(0), sheet.CellByRef("D10").ml.Style)
	require.Equal(t, styles.DirectStyleID(0), sheet.CellByRef("E10").ml.Style)

	r = sheet.Range(3, 9, 4, 9)
	require.Equal(t, []string{"1", "2"}, r.Values())
	require.Equal(t, styles.DirectStyleID(0), sheet.CellByRef("D10").ml.Style)
	require.Equal(t, styles.DirectStyleID(0), sheet.CellByRef("E10").ml.Style)

	//test styles
	style := styles.New(
		styles.Font.Name("Calibri"),
		styles.Font.Size(12),
	)

	styleRef := xl.AddStyles(style)
	r.SetStyles(styleRef)

	require.Equal(t, styles.DirectStyleID(styleRef), sheet.CellByRef("D10").ml.Style)
	require.Equal(t, styles.DirectStyleID(styleRef), sheet.CellByRef("E10").ml.Style)

	r.Clear()
	require.Equal(t, []string{"", ""}, r.Values())
	require.Equal(t, styles.DirectStyleID(styleRef), sheet.CellByRef("D10").ml.Style)
	require.Equal(t, styles.DirectStyleID(styleRef), sheet.CellByRef("E10").ml.Style)

	r.Reset()
	require.Equal(t, []string{"", ""}, r.Values())
	require.Equal(t, styles.DirectStyleID(0), sheet.CellByRef("D10").ml.Style)
	require.Equal(t, styles.DirectStyleID(0), sheet.CellByRef("E10").ml.Style)
}
