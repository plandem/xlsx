package xlsx_test

import (
	"fmt"
	"github.com/plandem/xlsx"
	"strings"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/plandem/xlsx/options"
	"github.com/plandem/xlsx/types"
)

//output content via normal mode
func ExampleSheetReadStream_ReadWrite() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	sheet := xl.Sheet(0)

	defer xl.Close()
	defer sheet.Close()
	fmt.Println(sheet.Dimension())

	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}

/*
Output:
14 28
,,,,,,,,,,,,,
    with leading space,,merged rows,,merged cols,merged cols,merged cols,,,,,,,
,,merged rows,,,,,,,,,,,
,,merged rows,,merged rows+cols,merged rows+cols,merged rows+cols,,,,,,,
with trailing space   ,,merged rows,,merged rows+cols,merged rows+cols,merged rows+cols,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,1,2,3,4,5,,,,,,
,,,6,7,8,9,10,,,,,,
,,,11,12,13,14,15,,,,,,
,,,16,17,18,19,20,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,last cell
*/
}

//output content using single phase only - it will be without merged cells info - only first cell from range will be output
func ExampleSheetReadStream_SinglePhased() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	sheet := xl.SheetReader(0, false)

	defer xl.Close()
	defer sheet.Close()
	fmt.Println(sheet.Dimension())

	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}

/*
Output:
14 28
,,,,,,,,,,,,,
    with leading space,,merged rows,,merged cols,,,,,,,,,
,,,,,,,,,,,,,
,,,,merged rows+cols,,,,,,,,,
with trailing space   ,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,1,2,3,4,5,,,,,,
,,,6,7,8,9,10,,,,,,
,,,11,12,13,14,15,,,,,,
,,,16,17,18,19,20,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,last cell
*/
}

//output content using multi phases it must be same as for normal mode
func ExampleSheetReadStream_MultiPhased() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	sheet := xl.SheetReader(0, true)

	defer xl.Close()
	defer sheet.Close()
	fmt.Println(sheet.Dimension())

	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}

/*
Output:
14 28
,,,,,,,,,,,,,
    with leading space,,merged rows,,merged cols,merged cols,merged cols,,,,,,,
,,merged rows,,,,,,,,,,,
,,merged rows,,merged rows+cols,merged rows+cols,merged rows+cols,,,,,,,
with trailing space   ,,merged rows,,merged rows+cols,merged rows+cols,merged rows+cols,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,1,2,3,4,5,,,,,,
,,,6,7,8,9,10,,,,,,
,,,11,12,13,14,15,,,,,,
,,,16,17,18,19,20,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,
,,,,,,,,,,,,,last cell
*/
}

func TestSheetReadStream_notImplemented(t *testing.T) {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}
	defer xl.Close()

	sheet := xl.SheetReader(0, true)
	defer sheet.Close()

	require.Panics(t, func() { sheet.Col(0)})
	require.Panics(t, func() { sheet.Cols()})
	require.Panics(t, func() { sheet.InsertCol(0)})
	require.Panics(t, func() { sheet.InsertRow(0)})
	require.Panics(t, func() { sheet.DeleteRow(0)})
	require.Panics(t, func() { sheet.DeleteCol(0)})
	require.Panics(t, func() { sheet.SetDimension(100, 100)})
	require.Panics(t, func() { sheet.SetActive()})
	require.Panics(t, func() { sheet.Set(options.NewSheetOptions(options.Sheet.Visibility(types.VisibilityTypeVisible)))})
	require.Panics(t, func() { sheet.SetName("aaa")})
}

func TestSheetReadStream_access(t *testing.T) {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()

	sheet := xl.SheetReader(0, true)
	defer sheet.Close()

	require.Equal(t, "8", sheet.CellByRef("F11").Value())
	require.Equal(t, "", sheet.CellByRef("F10").Value())
	require.Equal(t, "8", sheet.Cell(5, 10).Value())
	require.Equal(t, []string{"", "", "", "", "", "6","7","8","9","10","11","12","13","14","15","16","17","18","19","20"}, sheet.Range("D10:H13").Values())
}