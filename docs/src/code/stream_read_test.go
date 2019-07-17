package code

import (
	"fmt"
	"github.com/plandem/xlsx"
)

func Example_readStream() {
	xl, err := xlsx.Open("./foo.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0, xlsx.SheetModeStream)
	defer sheet.Close()

	totalCols, totalRows := sheet.Dimension()
	for rIdx := 0; rIdx < totalRows; rIdx++ {
		for cIdx := 0; cIdx < totalCols; cIdx++ {
			cell := sheet.Cell(cIdx, rIdx)
			fmt.Println(cell.String())
		}
	}
}
