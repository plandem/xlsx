package xlsx_test

import (
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize"
	"math/rand"
	"testing"
	"time"

	ooxml "github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format"
)

const simpleFile = "./test_files/simple.xlsx"
const bigFile = "./test_files/example_big.xlsx"
const hugeFile = "./test_files/example_huge.xlsx"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type openFileFn func(fileName string) (interface{}, interface{})

func tealegOpen(fileName string) (interface{}, interface{}) {
	xl, err := xlsx.OpenFile(fileName)
	if err != nil {
		panic(err)
	}

	return xl, xl.Sheets[0]
}

func excelizeOpen(fileName string) (interface{}, interface{}) {
	xl, err := excelize.OpenFile(fileName)
	if err != nil {
		panic(err)
	}

	return xl, "Sheet1"
}

func xlsxOpen(fileName string) (interface{}, interface{}) {
	xl, err := ooxml.Open(fileName)
	if err != nil {
		panic(err)
	}

	return xl, xl.Sheet(0)
}

func xlsxReadStream(fileName string) (interface{}, interface{}) {
	xl, err := ooxml.Open(fileName)
	if err != nil {
		panic(err)
	}

	return xl, xl.SheetReader(0, false)
}

func BenchmarkLibsRandomGet(b *testing.B) {
	benchmarks := []struct {
		name     string
		open     openFileFn
		callback func(f interface{}, s interface{}, value *string, x int, y int)
	}{
		{"excelize", excelizeOpen, func(f interface{}, s interface{}, value *string, maxCols, maxRows int) {
			xl := f.(*excelize.File)
			axis, _ := excelize.CoordinatesToCellName(rand.Intn(maxCols), 1+rand.Intn(maxRows))
			*value, _ = xl.GetCellValue("Sheet1", axis)
		}},
		{"tealeg", tealegOpen, func(f interface{}, s interface{}, value *string, maxCols, maxRows int) {
			sheet := s.(*xlsx.Sheet)
			*value = sheet.Cell(rand.Intn(maxCols), rand.Intn(maxRows)).Value
		}},
		{"xlsx", xlsxOpen, func(f interface{}, s interface{}, value *string, maxCols, maxRows int) {
			sheet := s.(ooxml.Sheet)
			*value = sheet.Cell(rand.Intn(maxCols), rand.Intn(maxRows)).Value()
		}},
	}

	const maxCols = 100
	const maxRows = 100
	var value string

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			f, sheet := bm.open(simpleFile)
			for i := 0; i < b.N; i++ {
				bm.callback(f, sheet, &value, maxCols, maxRows)
			}
		})
	}
}

func BenchmarkLibsRandomSet(b *testing.B) {
	benchmarks := []struct {
		name     string
		open     openFileFn
		callback func(f interface{}, s interface{}, value *string, x int, y int)
	}{
		{"excelize", excelizeOpen, func(f interface{}, s interface{}, value *string, maxCols, maxRows int) {
			xl := f.(*excelize.File)
			axis, _ := excelize.CoordinatesToCellName(rand.Intn(maxCols), 1+rand.Intn(maxRows))
			xl.SetCellValue("Sheet1", axis, rand.Intn(100))
		}},
		{"tealeg", tealegOpen, func(f interface{}, s interface{}, value *string, maxCols, maxRows int) {
			sheet := s.(*xlsx.Sheet)
			sheet.Cell(rand.Intn(maxCols), rand.Intn(maxRows)).SetValue(rand.Intn(100))
		}},
		{"xlsx", xlsxOpen, func(f interface{}, s interface{}, value *string, maxCols, maxRows int) {
			sheet := s.(ooxml.Sheet)
			sheet.Cell(rand.Intn(maxCols), rand.Intn(maxRows)).SetValue(rand.Intn(100))
		}},
	}

	const maxCols = 100
	const maxRows = 100
	var value string

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			f, sheet := bm.open(simpleFile)
			for i := 0; i < b.N; i++ {
				bm.callback(f, sheet, &value, maxCols, maxRows)
			}
		})
	}
}

func BenchmarkLibsRandomSetStyle(b *testing.B) {
	benchmarks := []struct {
		name     string
		open     openFileFn
		create   func(f interface{}) interface{}
		callback func(f interface{}, s interface{}, ss interface{}, x int, y int)
	}{
		{"excelize", excelizeOpen, func(f interface{}) interface{} {
			xl := f.(*excelize.File)
			style, err := xl.NewStyle(`{"custom_number_format": "[$-380A]dddd\\,\\ dd\" de \"mmmm\" de \"yyyy;@"}`)
			if err != nil {
				panic(err)
			}

			return style
		}, func(f interface{}, s interface{}, ss interface{}, maxCols, maxRows int) {
			xl := f.(*excelize.File)
			styleId := ss.(int)

			axis, _ := excelize.CoordinatesToCellName(rand.Intn(maxCols), 1+rand.Intn(maxRows))
			xl.SetCellStyle("Sheet1", axis, axis, styleId)
		}},
		{"tealeg", tealegOpen, func(f interface{}) interface{} {
			style := xlsx.NewStyle()
			font := *xlsx.NewFont(12, "Verdana")
			font.Bold = true
			font.Italic = true
			font.Underline = true
			style.Font = font
			fill := *xlsx.NewFill("solid", "00FF0000", "FF000000")
			style.Fill = fill
			border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
			style.Border = border
			style.ApplyBorder = true

			return style
		}, func(f interface{}, s interface{}, ss interface{}, maxCols, maxRows int) {
			sheet := s.(*xlsx.Sheet)
			style := ss.(*xlsx.Style)
			sheet.Cell(rand.Intn(maxCols), rand.Intn(maxRows)).SetStyle(style)
		}},
		{"xlsx", xlsxOpen, func(f interface{}) interface{} {
			xl := f.(*ooxml.Spreadsheet)

			style := format.New(
				format.Font.Name("Calibri"),
				format.Font.Size(12),
				format.Font.Color("#FF0000"),
				format.Font.Scheme(format.FontSchemeMinor),
				format.Font.Family(format.FontFamilySwiss),

				format.Fill.Type(format.PatternTypeNone),

				format.Alignment.VAlign(format.VAlignBottom),
				format.Alignment.HAlign(format.HAlignFill),
				format.Border.Color("#ff00ff"),
				format.Border.Type(format.BorderStyleDashDot),
				format.Protection.Hidden,
				format.Protection.Locked,
				//format.NumberFormat("#.### usd"),
				format.Fill.Type(format.PatternTypeDarkDown),
				format.Fill.Color("#FFFFFF"),
				format.Fill.Background("#FF0000"),
			)

			return xl.AddFormatting(style)
		}, func(f interface{}, s interface{}, ss interface{}, maxCols, maxRows int) {
			sheet := s.(ooxml.Sheet)
			styleId := ss.(format.DirectStyleID)
			sheet.Cell(rand.Intn(maxCols), rand.Intn(maxRows)).SetFormatting(styleId)
		}},
	}

	const maxCols = 100
	const maxRows = 100

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			f, sheet := bm.open(simpleFile)
			style := bm.create(f)
			for i := 0; i < b.N; i++ {
				bm.callback(f, sheet, style, maxCols, maxRows)
			}
		})
	}
}

func BenchmarkLibsReadBigFile(b *testing.B) {
	benchmarks := []struct {
		name     string
		open     openFileFn
		callback func(f interface{}, s interface{}, value *string)
	}{
		{"excelize", excelizeOpen, func(f interface{}, s interface{}, value *string) {
			xl := f.(*excelize.File)
			rows, _ := xl.GetRows("Sheet1")

			for _, row := range rows {
				*value = row[0]
			}
		}},
		{"tealeg", tealegOpen, func(f interface{}, s interface{}, value *string) {
			sheet := s.(*xlsx.Sheet)
			for row_i, row_max := 0, len(sheet.Rows); row_i < row_max; row_i++ {
				*value = sheet.Cell(0, row_i).Value
			}
		}},
		{"xlsx", xlsxOpen, func(f interface{}, s interface{}, value *string) {
			sheet := s.(ooxml.Sheet)
			_, row_max := sheet.Dimension()
			for row_i := 0; row_i < row_max; row_i++ {
				*value = sheet.Cell(0, row_i).Value()
			}
		}},
		{"xlsx-stream", xlsxReadStream, func(f interface{}, s interface{}, value *string) {
			sheet := s.(ooxml.Sheet)
			_, row_max := sheet.Dimension()
			for row_i := 0; row_i < row_max; row_i++ {
				*value = sheet.Cell(0, row_i).Value()
			}
		}},
	}

	var value string
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				f, s := bm.open(bigFile)
				bm.callback(f, s, &value)
			}
		})
	}
}

func BenchmarkLibsReadHugeFile(b *testing.B) {
	b.ReportAllocs()
	benchmarks := []struct {
		name     string
		open     openFileFn
		callback func(f interface{}, s interface{}, value *string)
	}{
		{"excelize", excelizeOpen, func(f interface{}, s interface{}, value *string) {
			xl := f.(*excelize.File)
			rows, _ := xl.GetRows("Sheet1")

			for _, row := range rows {
				*value = row[0]
			}
		}},
		{"tealeg", tealegOpen, func(f interface{}, s interface{}, value *string) {
			sheet := s.(*xlsx.Sheet)
			for row_i, row_max := 0, len(sheet.Rows); row_i < row_max; row_i++ {
				*value = sheet.Cell(0, row_i).Value
			}
		}},
		{"xlsx", xlsxOpen, func(f interface{}, s interface{}, value *string) {
			sheet := s.(ooxml.Sheet)
			_, row_max := sheet.Dimension()
			for row_i := 0; row_i < row_max; row_i++ {
				*value = sheet.Cell(0, row_i).Value()
			}
		}},
		{"xlsx-stream", xlsxReadStream, func(f interface{}, s interface{}, value *string) {
			sheet := s.(ooxml.Sheet)
			_, row_max := sheet.Dimension()
			for row_i := 0; row_i < row_max; row_i++ {
				*value = sheet.Cell(0, row_i).Value()
			}
		}},
	}

	var value string
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				f, s := bm.open(hugeFile)
				bm.callback(f, s, &value)
			}
		})
	}
}

func BenchmarkLibsUpdateBigFile(b *testing.B) {
	benchmarks := []struct {
		name     string
		open     openFileFn
		callback func(f interface{}, s interface{}, value *string)
		close    func(f interface{})
	}{
		{"excelize", excelizeOpen, func(f interface{}, s interface{}, value *string) {
			xl := f.(*excelize.File)
			rows, _ := xl.GetRows("Sheet1")

			for _, row := range rows {
				*value = row[0]
			}
		}, func(f interface{}) {
			xl := f.(*excelize.File)
			xl.SaveAs("saved_big_excelize.xlsx")
		}},
		{"tealeg", tealegOpen, func(f interface{}, s interface{}, value *string) {
			sheet := s.(*xlsx.Sheet)
			for row_i, row_max := 0, len(sheet.Rows); row_i < row_max; row_i++ {
				*value = sheet.Cell(0, row_i).Value
			}
		}, func(f interface{}) {
			xl := f.(*xlsx.File)
			xl.Save("saved_big_tealeg.xlsx")
		}},
		{"xlsx", xlsxOpen, func(f interface{}, s interface{}, value *string) {
			sheet := s.(ooxml.Sheet)
			_, row_max := sheet.Dimension()
			for row_i := 0; row_i < row_max; row_i++ {
				*value = sheet.Cell(0, row_i).Value()
			}
		}, func(f interface{}) {
			xl := f.(*ooxml.Spreadsheet)
			xl.SaveAs("saved_big_xlsx.xlsx")
		}},
	}

	var value string
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				f, s := bm.open(bigFile)
				bm.callback(f, s, &value)
				bm.close(f)
			}
		})
	}
}

func BenchmarkLibsUpdateHugeFile(b *testing.B) {
	benchmarks := []struct {
		name     string
		open     openFileFn
		callback func(f interface{}, s interface{}, value *string)
		close    func(f interface{})
	}{
		{"excelize", excelizeOpen, func(f interface{}, s interface{}, value *string) {
			xl := f.(*excelize.File)
			rows, _ := xl.GetRows("Sheet1")

			for _, row := range rows {
				*value = row[0]
			}
		}, func(f interface{}) {
			xl := f.(*excelize.File)
			xl.SaveAs("saved_huge_excelize.xlsx")
		}},
		//{"tealeg", tealegOpen, func(f interface{}, s interface{}, value *string) {
		//	sheet := s.(*xlsx.Sheet)
		//	for row_i, row_max := 0, len(sheet.Rows); row_i < row_max; row_i++ {
		//		*value = sheet.Cell(0, row_i).Value
		//	}
		//}, func(f interface{}) {
		//	xl := f.(*xlsx.File)
		//	xl.Save("saved_huge_tealeg.xlsx")
		//}},
		{"xlsx", xlsxOpen, func(f interface{}, s interface{}, value *string) {
			sheet := s.(ooxml.Sheet)
			_, row_max := sheet.Dimension()
			for row_i := 0; row_i < row_max; row_i++ {
				*value = sheet.Cell(0, row_i).Value()
			}
		}, func(f interface{}) {
			xl := f.(*ooxml.Spreadsheet)
			xl.SaveAs("saved_huge_xlsx.xlsx")
		}},
	}

	var value string
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				f, s := bm.open(hugeFile)
				bm.callback(f, s, &value)
				bm.close(f)
			}
		})
	}
}
