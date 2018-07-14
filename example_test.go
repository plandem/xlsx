package xlsx_test

import (
	"fmt"
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/options"
	"github.com/plandem/xlsx/types"
	"log"
	"os"
	"strings"
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

	// Get col by 0-based index
	col := sheet.Col(3)
	fmt.Println(strings.Join(col.Values(), ","))

	// Get range by references
	area := sheet.Range("D10:H13")
	fmt.Println(strings.Join(area.Values(), ","))

	//Output:
	// last cell
	// last cell
	// ,,,1,2,3,4,5,,,,,,
	// ,,,,,,,,,1,6,11,16,,,,,,,,,,,,,,,
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
	r := sheet.Range("A1:B3")
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
	area := sheet.Range("A1:B3")
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
	area := sheet.Range("D10:H13")
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
	redBold := format.New(
		format.Font.Bold,
		format.Font.Color("#ff0000"),
		format.Fill.Background("#ffff00"),
		format.Fill.Type(format.PatternTypeSolid),
	)

	// Add formatting to xlsx
	styleId := xl.AddFormatting(redBold)

	sheet := xl.Sheet(0)

	// Set formatting for cell
	sheet.CellByRef("N28").SetFormatting(styleId)

	// Set DEFAULT formatting for row. Affects cells not yet allocated in the row.
	// In other words, this style applies to new cells.
	sheet.Row(9).SetFormatting(styleId)

	// Set DEFAULT formatting for col. Affects cells not yet allocated in the col.
	// In other words, this style applies to new cells.
	sheet.Col(3).SetFormatting(styleId)

	//set formatting for all cells in range
	sheet.Range("D10:H13").SetFormatting(styleId)
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
	rowOptions := options.NewRowOptions(
		options.Row.Hidden(true),
		options.Row.Height(10.0),
		options.Row.Collapsed(true),
	)
	sheet.Row(9).Set(rowOptions)

	// set options for col
	colOptions := options.NewColumnOptions(
		options.Column.Hidden(true),
		options.Column.Width(10.0),
		options.Column.Collapsed(true),
	)
	sheet.Col(3).Set(colOptions)

	// set options for sheet
	sheetOptions := options.NewSheetOptions(
		options.Sheet.Visibility(types.VisibilityTypeVeryHidden),
	)
	sheet.Set(sheetOptions)
}

// Demonstrates how to append cols/rows/sheets.
func Example_append() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	// To append a new col/row, simple request it - sheet will be auto expanded. E.g.: we have 14 cols, 28 rows.
	fmt.Println(sheet.Dimension())

	// Append 72 rows
	sheet.Row(99)
	fmt.Println(sheet.Dimension())

	// Append 36 cols
	sheet.Col(49)
	fmt.Println(sheet.Dimension())

	// Append 3 sheet
	fmt.Println(strings.Join(xl.GetSheetNames(), ","))
	xl.AddSheet("new sheet")
	xl.AddSheet("new sheet")
	xl.AddSheet("new sheet")
	fmt.Println(strings.Join(xl.GetSheetNames(), ","))

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
	fmt.Println(strings.Join(xl.GetSheetNames(), ","))
	xl.DeleteSheet(0)
	fmt.Println(strings.Join(xl.GetSheetNames(), ","))

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
