package xlsx_test

import (
	"fmt"
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format"
	"log"
	"strings"
)

func ExampleSpreadsheet_GetSheetNames() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	fmt.Println(xl.GetSheetNames())
	//Output:
	// [Sheet1]
}

func ExampleSpreadsheet_Sheet() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//nil, if there is no sheet with requested index
	if sheet := xl.Sheet(12345); sheet == nil {
		fmt.Println("Unknown sheet")
	}

	if sheet := xl.Sheet(0); sheet != nil {
		fmt.Println(sheet.Name())
	}

	//Output:
	// Unknown sheet
	// Sheet1
}

func ExampleSpreadsheet_AddSheet() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.AddSheet("New sheet")

	//now you can use sheet as always
	fmt.Println(sheet.Name())

	err = xl.SaveAs("new_file.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	//Output:
	// New sheet
}

func ExampleSpreadsheet_DeleteSheet() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//add a new sheet, next index is 1
	xl.AddSheet("New sheet")

	//delete a sheet with index 0
	xl.DeleteSheet(0)
}

func ExampleSpreadsheet_AddFormatting() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//create a new format for a bold font with red color and yellow solid background
	redBold := format.New(
		format.Font.Bold,
		format.Font.Color("#ff0000"),
		format.Fill.Background("#ffff00"),
		format.Fill.Type(format.PatternTypeSolid),
	)

	//add formatting to xlsx
	styleId := xl.AddFormatting(redBold)

	//now you can use this id wherever you need
	_ = styleId
}

func ExampleSpreadsheet_Sheets() {
	xl, err := xlsx.Open("./test_files/example_iteration.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//output names of sheets
	fmt.Println(strings.Join(xl.GetSheetNames(), ","))

	//iterate sheets via iterator
	for sheets := xl.Sheets(); sheets.HasNext(); {
		_, sheet := sheets.Next()
		fmt.Println(sheet.Name())
	}

	//Output:
	//First Sheet,Second Sheet,Last Sheet
	//First Sheet
	//Second Sheet
	//Last Sheet
}
