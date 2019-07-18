package code

import (
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/types"
)

func Example_writeStream() {
	xl := xlsx.New()

	sheet := xl.AddSheet("Sheet1", xlsx.SheetModeStream)

	for iRow, iRowMax := 0, 100; iRow < iRowMax; iRow++ {
		for iCol, iColMax := 0, 9; iCol < iColMax; iCol++ {
			sheet.Cell(iCol, iRow).SetValue(types.CellRefFromIndexes(iCol, iRow))
		}
	}

	xl.SaveAs("./foo.xlsx")
}
