package xlsx_test

import (
	"fmt"
	"github.com/plandem/xlsx"
	"log"
	"strings"
)

func ExampleSheetReadWrite_SetActive() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//add a new sheet, next index is 1
	sheet := xl.AddSheet("New sheet")

	//set sheet as active
	sheet.SetActive()
}

func ExampleSheetReadWrite_Cell() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	//get cell by 0-based indexes, e.g.: 13,27 is same as N28
	cell := sheet.Cell(13, 27)

	fmt.Println(cell.Value())
	//Output:
	// last cell
}

func ExampleSheetReadWrite_CellByRef() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	//get cell by reference, e.g.: N28 is same as 13,27
	cell := sheet.CellByRef("N28")

	fmt.Println(cell.Value())
	//Output:
	// last cell
}

func ExampleSheetReadWrite_Row() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)
	_, totalRows := sheet.Dimension()
	for rIdx := 0; rIdx < totalRows; rIdx++ {
		row := sheet.Row(rIdx)
		fmt.Println(strings.Join(row.Values(), ","))

		for cells := row.Cells(); cells.HasNext(); {
			_, _, cell := cells.Next()
			fmt.Println(cell.Value())
		}
	}

	//Output:
	//Header 1,Header 2
	//Header 1
	//Header 2
	//Value 1-1,Value 2-1
	//Value 1-1
	//Value 2-1
	//Value 1-2,Value 2-2
	//Value 1-2
	//Value 2-2
}

func ExampleSheetReadWrite_Col() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)
	totalCols, _ := sheet.Dimension()
	for cIdx := 0; cIdx < totalCols; cIdx++ {
		col := sheet.Col(cIdx)
		fmt.Println(strings.Join(col.Values(), ","))

		for cells := col.Cells(); cells.HasNext(); {
			_, _, cell := cells.Next()
			fmt.Println(cell.Value())
		}
	}

	//Output:
	//Header 1,Value 1-1,Value 1-2
	//Header 1
	//Value 1-1
	//Value 1-2
	//Header 2,Value 2-1,Value 2-2
	//Header 2
	//Value 2-1
	//Value 2-2
}

func ExampleSheetReadWrite_Rows() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//get sheet by 0-based index
	sheet := xl.Sheet(0)

	//iterate rows via iterator
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		fmt.Println(strings.Join(row.Values(), ","))

		for cells := row.Cells(); cells.HasNext(); {
			_, _, cell := cells.Next()
			fmt.Println(cell.Value())
		}
	}

	//Output:
	//Header 1,Header 2
	//Header 1
	//Header 2
	//Value 1-1,Value 2-1
	//Value 1-1
	//Value 2-1
	//Value 1-2,Value 2-2
	//Value 1-2
	//Value 2-2
}

func ExampleSheetReadWrite_Cols() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//get sheet by 0-based index
	sheet := xl.Sheet(0)

	//iterate cols via iterator
	for cols := sheet.Cols(); cols.HasNext(); {
		_, col := cols.Next()
		fmt.Println(strings.Join(col.Values(), ","))

		for cells := col.Cells(); cells.HasNext(); {
			_, _, cell := cells.Next()
			fmt.Println(cell.Value())
		}
	}

	//Output:
	//Header 1,Value 1-1,Value 1-2
	//Header 1
	//Value 1-1
	//Value 1-2
	//Header 2,Value 2-1,Value 2-2
	//Header 2
	//Value 2-1
	//Value 2-2
}

func ExampleSheetReadWrite_Range() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//get sheet by 0-based index
	sheet := xl.Sheet(0)
	r := sheet.Range("A1:B3")
	for cells := r.Cells(); cells.HasNext(); {
		_, _, cell := cells.Next()
		fmt.Println(cell.Value())
	}
	//Output:
	//Header 1
	//Header 2
	//Value 1-1
	//Value 2-1
	//Value 1-2
	//Value 2-2
}

func ExampleSheetReadWrite_Walk() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//get sheet by 0-based index
	sheet := xl.Sheet(0)
	r := sheet.Range("A1:B3")
	r.Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		c.SetValue(idx)
	})

	fmt.Println(strings.Join(r.Values(), ","))
	//Output:
	//0,1,2,3,4,5
}
