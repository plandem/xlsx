// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx_test

import (
	"fmt"
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/styles"
	colOptions "github.com/plandem/xlsx/types/options/column"
	rowOptions "github.com/plandem/xlsx/types/options/row"
	sheetOptions "github.com/plandem/xlsx/types/options/sheet"
	"log"
	"os"
	"strings"
	"time"
)

// Demonstrates how to create/open/save XLSX files
func Example_files() {
	// Create a new XLSX file
	xl := xlsx.New()

	// Open the XLSX file using file name
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	// Open the XLSX file using file handler
	zipFile, err := os.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	xl, err = xlsx.Open(zipFile)
	if err != nil {
		log.Fatal(err)
	}

	// Update the existing XLSX file
	err = xl.Save()
	if err != nil {
		log.Fatal(err)
	}

	// Save the XLSX file under different name
	err = xl.SaveAs("new_file.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}

// Demonstrates how to access differ information
func Example_access() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	// Get sheet by 0-based index
	sheet := xl.Sheet(0)

	// Get cell by 0-based indexes
	cell := sheet.Cell(13, 27)
	fmt.Println(cell.Value())

	// Get cell by reference
	cell = sheet.CellByRef("N28")
	fmt.Println(cell.Value())

	// Get row by 0-based index
	row := sheet.Row(9)
	fmt.Println(strings.Join(row.Values(), ","))

	// Get cell of row at 0-based col index
	cell = row.Cell(0)
	fmt.Println(cell.Value())

	// Get col by 0-based index
	col := sheet.Col(3)
	fmt.Println(strings.Join(col.Values(), ","))

	// Get cell of col at 0-based row index
	cell = col.Cell(0)
	fmt.Println(cell.Value())

	// Get range by references
	area := sheet.RangeByRef("D10:H13")
	fmt.Println(strings.Join(area.Values(), ","))

	//Output:
	// last cell
	// last cell
	// ,,,1,2,3,4,5,,,,,,
	//
	// ,,,,,,,,,1,6,11,16,,,,,,,,,,,,,,,
	//
	// 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20
}

// Demonstrates how to iterate
func Example_iterate() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	// Get sheet by 0-based index
	sheet := xl.Sheet(0)

	// Iterate by indexes
	totalCols, totalRows := sheet.Dimension()
	for rIdx := 0; rIdx < totalRows; rIdx++ {
		for cIdx := 0; cIdx < totalCols; cIdx++ {
			fmt.Println(sheet.Cell(cIdx, rIdx).Value())
		}
	}

	// Iterate rows via iterator
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		for cells := row.Cells(); cells.HasNext(); {
			_, _, cell := cells.Next()
			fmt.Println(cell.Value())
		}
	}

	// Iterate cols via iterator
	for cols := sheet.Cols(); cols.HasNext(); {
		_, col := cols.Next()
		for cells := col.Cells(); cells.HasNext(); {
			_, _, cell := cells.Next()
			fmt.Println(cell.Value())
		}
	}

	// Iterate range's cells via iterator
	r := sheet.RangeByRef("A1:B3")
	for cells := r.Cells(); cells.HasNext(); {
		_, _, cell := cells.Next()
		fmt.Println(cell.Value())
	}

	// Iterate sheets via iterator
	for sheets := xl.Sheets(); sheets.HasNext(); {
		_, sheet := sheets.Next()
		fmt.Println(sheet.Name())
	}

	//Output:
	//Header 1
	//Header 2
	//Value 1-1
	//Value 2-1
	//Value 1-2
	//Value 2-2
	//Header 1
	//Header 2
	//Value 1-1
	//Value 2-1
	//Value 1-2
	//Value 2-2
	//Header 1
	//Value 1-1
	//Value 1-2
	//Header 2
	//Value 2-1
	//Value 2-2
	//Header 1
	//Header 2
	//Value 1-1
	//Value 2-1
	//Value 1-2
	//Value 2-2
	//First Sheet
	//Second Sheet
	//Last Sheet
}

// Demonstrate walk cells using callback
func Example_walk() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	// Get sheet by 0-based index
	sheet := xl.Sheet(0)

	// Walk through the cells of row
	row := sheet.Row(0)
	row.Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		fmt.Println(c.Value())
	})

	// Walk through the cells of col
	col := sheet.Col(0)
	col.Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		fmt.Println(c.Value())
	})

	// Walk through the cells of range
	area := sheet.RangeByRef("A1:B3")
	area.Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		fmt.Println(c.Value())
	})

	//Output:
	//Header 1
	//Header 2
	//Header 1
	//Value 1-1
	//Value 1-2
	//Header 1
	//Header 2
	//Value 1-1
	//Value 2-1
	//Value 1-2
	//Value 2-2
}

// Demonstrates how to update information
func Example_update() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	// Update value of cell
	cell := sheet.Cell(13, 27)
	fmt.Println(cell.Value())
	cell.SetValue("new value")
	fmt.Println(cell.Value())

	// Update value of cells in row
	row := sheet.Row(9)
	fmt.Println(strings.Join(row.Values(), ","))
	row.Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		c.SetValue(idx)
	})
	fmt.Println(strings.Join(row.Values(), ","))

	// Update value of cells in col
	col := sheet.Col(3)
	fmt.Println(strings.Join(col.Values(), ","))
	col.Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		c.SetValue(idx)
	})
	fmt.Println(strings.Join(col.Values(), ","))

	// Update value of cells in range
	area := sheet.RangeByRef("D10:H13")
	fmt.Println(strings.Join(area.Values(), ","))
	area.Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		c.SetValue(idx)
	})
	fmt.Println(strings.Join(area.Values(), ","))

	//Output:
	// last cell
	// new value
	// ,,,1,2,3,4,5,,,,,,
	// 0,1,2,3,4,5,6,7,8,9,10,11,12,13
	// ,,,,,,,,,3,6,11,16,,,,,,,,,,,,,,,
	// 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27
	// 9,4,5,6,7,10,7,8,9,10,11,12,13,14,15,12,17,18,19,20
	// 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19
}

// Demonstrates how to add style formatting
func Example_formatting() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	// Create a new format for a bold font with red color and yellow solid background
	redBold := styles.New(
		styles.Font.Bold,
		styles.Font.Color("#ff0000"),
		styles.Fill.Background("#ffff00"),
		styles.Fill.Type(styles.PatternTypeSolid),
	)

	// Add formatting to xlsx
	styleId := xl.AddStyles(redBold)

	sheet := xl.Sheet(0)

	// Set formatting for cell
	sheet.CellByRef("N28").SetStyles(styleId)

	// Set DEFAULT formatting for row. Affects cells not yet allocated in the row.
	// In other words, this style applies to new cells.
	sheet.Row(9).SetStyles(styleId)

	// Set DEFAULT formatting for col. Affects cells not yet allocated in the col.
	// In other words, this style applies to new cells.
	sheet.Col(3).SetStyles(styleId)

	//set formatting for all cells in range
	sheet.RangeByRef("D10:H13").SetStyles(styleId)
}

// Demonstrates how to set options of rows/cols/sheets
func Example_options() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	// set options for row
	ro := rowOptions.New(
		rowOptions.Hidden(true),
		rowOptions.Height(10.0),
		rowOptions.Collapsed(true),
	)
	sheet.Row(9).SetOptions(ro)

	// set options for col
	co := colOptions.New(
		colOptions.Hidden(true),
		colOptions.Width(10.0),
		colOptions.Collapsed(true),
	)
	sheet.Col(3).SetOptions(co)

	// set options for sheet
	so := sheetOptions.New(
		sheetOptions.Visibility(sheetOptions.VisibilityVeryHidden),
	)
	sheet.SetOptions(so)
}

// Demonstrates how to append cols/rows/sheets.
func Example_append() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	// To append a new col/row, simple request it - sheet will be auto expanded.
	// E.g.: we have 14 cols, 28 rows.
	fmt.Println(sheet.Dimension())

	// Append 72 rows
	sheet.Row(99)
	fmt.Println(sheet.Dimension())

	// Append 36 cols
	sheet.Col(49)
	fmt.Println(sheet.Dimension())

	// Append 3 sheet
	fmt.Println(strings.Join(xl.SheetNames(), ","))
	xl.AddSheet("new sheet")
	xl.AddSheet("new sheet")
	xl.AddSheet("new sheet")
	fmt.Println(strings.Join(xl.SheetNames(), ","))

	//Output:
	// 14 28
	// 14 100
	// 50 100
	// Sheet1
	// Sheet1,new sheet,new sheet1,new sheet2
}

// Demonstrates how to insert cols/rows
func Example_insert() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	fmt.Println(sheet.Dimension())
	fmt.Println(strings.Join(sheet.Col(3).Values(), ","))

	// Insert a new col
	sheet.InsertCol(3)
	fmt.Println(sheet.Dimension())
	fmt.Println(strings.Join(sheet.Col(3).Values(), ","))
	fmt.Println(strings.Join(sheet.Col(4).Values(), ","))

	// Insert a new row
	fmt.Println(strings.Join(sheet.Row(9).Values(), ","))
	sheet.InsertRow(3)
	fmt.Println(sheet.Dimension())
	fmt.Println(strings.Join(sheet.Row(9).Values(), ","))
	fmt.Println(strings.Join(sheet.Row(10).Values(), ","))

	//Output:
	// 14 28
	// ,,,,,,,,,1,6,11,16,,,,,,,,,,,,,,,
	// 15 28
	// ,,,,,,,,,,,,,,,,,,,,,,,,,,,
	// ,,,,,,,,,1,6,11,16,,,,,,,,,,,,,,,
	// ,,,,1,2,3,4,5,,,,,,
	// 15 29
	// ,,,,,,,,,,,,,,
	// ,,,,1,2,3,4,5,,,,,,
}

// Demonstrates how to delete information
func Example_delete() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	fmt.Println(sheet.Dimension())

	// Delete col
	fmt.Println(strings.Join(sheet.Col(3).Values(), ","))
	sheet.DeleteCol(3)
	fmt.Println(sheet.Dimension())
	fmt.Println(strings.Join(sheet.Col(3).Values(), ","))

	// Delete row
	fmt.Println(strings.Join(sheet.Row(3).Values(), ","))
	sheet.DeleteRow(3)
	fmt.Println(sheet.Dimension())
	fmt.Println(strings.Join(sheet.Row(3).Values(), ","))

	// Delete sheet
	fmt.Println(strings.Join(xl.SheetNames(), ","))
	xl.DeleteSheet(0)
	fmt.Println(strings.Join(xl.SheetNames(), ","))

	//Output:
	// 14 28
	// ,,,,,,,,,1,6,11,16,,,,,,,,,,,,,,,
	// 13 28
	// ,merged cols,,merged rows+cols,,,,,,2,7,12,17,,,,,,,,,,,,,,,
	// ,,merged rows,merged rows+cols,,,,,,,,,
	// 13 27
	// with trailing space   ,,merged rows,,,,,,,,,,
	// Sheet1
	//
}

// Demonstrates how to open sheet in streaming mode
func Example_streams() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	// Open sheet in stream reading mode with single phase.
	// Some meta information is NOT available (e.g. merged cells).
	sheet := xl.Sheet(0, xlsx.SheetModeStream)
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}
	sheet.Close()

	// Open sheet in stream reading mode with multi phases.
	// Meta information is available.
	sheet = xl.Sheet(0, xlsx.SheetModeStream, xlsx.SheetModeMultiPhase)
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}
	sheet.Close()
}

// Demonstrates how to copy information in sheet
func Example_copy() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}

	// Copy row to another row with index
	row := sheet.Row(0)
	row.CopyTo(4, false)
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}

	// Copy col to another col with index
	col := sheet.Col(0)
	col.CopyTo(3, false)
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}

	// Copy range to another range that started at indexes
	r := sheet.RangeByRef("A1:B3")
	r.CopyTo(3, 0)
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}

	// Copy range to another range that started at ref
	r.CopyToRef("I4")
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))
	}

	//Output:
	//Header 1,Header 2
	//Value 1-1,Value 2-1
	//Value 1-2,Value 2-2
	//Header 1,Header 2
	//Value 1-1,Value 2-1
	//Value 1-2,Value 2-2
	//,
	//Header 1,Header 2
	//Header 1,Header 2,,Header 1
	//Value 1-1,Value 2-1,,Value 1-1
	//Value 1-2,Value 2-2,,Value 1-2
	//,,,
	//Header 1,Header 2,,Header 1
	//Header 1,Header 2,,Header 1,Header 2
	//Value 1-1,Value 2-1,,Value 1-1,Value 2-1
	//Value 1-2,Value 2-2,,Value 1-2,Value 2-2
	//,,,,
	//Header 1,Header 2,,Header 1,
	//Header 1,Header 2,,Header 1,Header 2,,,,,
	//Value 1-1,Value 2-1,,Value 1-1,Value 2-1,,,,,
	//Value 1-2,Value 2-2,,Value 1-2,Value 2-2,,,,,
	//,,,,,,,,Header 1,Header 2
	//Header 1,Header 2,,Header 1,,,,,Value 1-1,Value 2-1
	//,,,,,,,,Value 1-2,Value 2-2
}

// Demonstrates how to get/set value for cell
func Example_gettersAndSetters() {
	xl := xlsx.New()
	defer xl.Close()

	sheet := xl.AddSheet("test sheet")

	now, _ := time.Parse("02 Jan 06 15:04 MST", time.RFC822)

	//set values by typed method
	sheet.CellByRef("A1").SetText("string")
	sheet.CellByRef("B1").SetInlineText("inline string")
	sheet.CellByRef("C1").SetBool(true)
	sheet.CellByRef("D1").SetInt(12345)
	sheet.CellByRef("E1").SetFloat(123.123)
	sheet.CellByRef("F1").SetDateTime(now)
	sheet.CellByRef("G1").SetDate(now)
	sheet.CellByRef("H1").SetTime(now)
	sheet.CellByRef("I1").SetDeltaTime(now)

	//set values by unified method
	sheet.CellByRef("A2").SetValue("string")
	sheet.CellByRef("B2").SetValue(true)
	sheet.CellByRef("C2").SetValue(12345)
	sheet.CellByRef("D2").SetValue(123.123)
	sheet.CellByRef("E2").SetValue(now)

	//get raw values that were set via typed setter
	fmt.Println(sheet.CellByRef("A1").Value())
	fmt.Println(sheet.CellByRef("B1").Value())
	fmt.Println(sheet.CellByRef("C1").Value())
	fmt.Println(sheet.CellByRef("D1").Value())
	fmt.Println(sheet.CellByRef("E1").Value())
	fmt.Println(sheet.CellByRef("F1").Value())
	fmt.Println(sheet.CellByRef("G1").Value())
	fmt.Println(sheet.CellByRef("H1").Value())
	fmt.Println(sheet.CellByRef("I1").Value())

	//get raw values that were set that via general setter
	fmt.Println(sheet.CellByRef("A2").Value())
	fmt.Println(sheet.CellByRef("B2").Value())
	fmt.Println(sheet.CellByRef("C2").Value())
	fmt.Println(sheet.CellByRef("D2").Value())
	fmt.Println(sheet.CellByRef("E2").Value())

	//get typed values and error if invalid type (values were set via typed setter)
	_ = sheet.CellByRef("A1").String()
	_ = sheet.CellByRef("B1").String()
	_, _ = sheet.CellByRef("C1").Bool()
	_, _ = sheet.CellByRef("D1").Int()
	_, _ = sheet.CellByRef("E1").Float()
	_, _ = sheet.CellByRef("F1").Date()

	//get typed values and error if invalid type (values were set via general setter)
	_ = sheet.CellByRef("A2").String()
	_, _ = sheet.CellByRef("B2").Bool()
	_, _ = sheet.CellByRef("C2").Int()
	_, _ = sheet.CellByRef("D2").Float()
	_, _ = sheet.CellByRef("E2").Date()

	//Output:
	//string
	//inline string
	//1
	//12345
	//123.123
	//2006-01-02T15:04:00
	//2006-01-02T15:04:00
	//2006-01-02T15:04:00
	//2006-01-02T15:04:00
	//string
	//1
	//12345
	//123.123
	//2006-01-02T15:04:00
}
