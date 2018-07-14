package xlsx_test

import (
	"fmt"
	"github.com/plandem/xlsx"
	"strings"
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
