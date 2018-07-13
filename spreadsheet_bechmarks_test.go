package xlsx_test

import (
	"github.com/plandem/xlsx"
	"log"
	"testing"
)

func readSheet(sheet xlsx.Sheet) {
	_, iMaxRow := sheet.Dimension()
	var value string

	for iRow := 0; iRow < iMaxRow; iRow++ {
		value = sheet.Cell(0, iRow).Value()
	}

	log.Println(value)
	_ = value
}

func BenchmarkSpreadsheet_Sheet(b *testing.B) {
	b.ReportAllocs()
	xl, err := xlsx.Open(hugeFile)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		sheet := xl.Sheet(0)
		readSheet(sheet)
		sheet.Close()
	}
}

func BenchmarkSpreadsheet_SheetReader(b *testing.B) {
	b.ReportAllocs()
	xl, err := xlsx.Open(hugeFile)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		sheet := xl.SheetReader(0)
		readSheet(sheet)
		sheet.Close()
	}
}
