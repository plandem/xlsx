package code

import (
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/styles"
)

func Example_mergedCells() {
	xl := xlsx.New()

	sheet := xl.AddSheet("")

	//range 1
	sheet.RangeByRef("B4:E4").Merge()
	sheet.CellByRef("B4").SetValue("Merged Range 1")
	sheet.CellByRef("B4").SetStyles(xl.AddStyles(
		styles.New(
			styles.Alignment.VAlign(styles.VAlignCenter),
			styles.Alignment.HAlign(styles.HAlignCenter),
			styles.Fill.Type(styles.PatternTypeSolid),
			styles.Fill.Color("#FFFF00"),
		),
	))

	//range 2
	sheet.RangeByRef("B6:B9").Merge()
	sheet.CellByRef("B6").SetValue("Merged Range 2")
	sheet.CellByRef("B6").SetStyles(xl.AddStyles(
		styles.New(
			styles.Alignment.VAlign(styles.VAlignCenter),
			styles.Alignment.HAlign(styles.HAlignCenter),
			styles.Alignment.WrapText,
			styles.Fill.Type(styles.PatternTypeSolid),
			styles.Fill.Color("#EEAA00"),
		),
	))

	//range 3
	sheet.RangeByRef("D6:E9").Merge()
	sheet.CellByRef("D6").SetValue("Merged Range 3")
	sheet.CellByRef("D6").SetStyles(xl.AddStyles(
		styles.New(
			styles.Alignment.VAlign(styles.VAlignCenter),
			styles.Alignment.HAlign(styles.HAlignCenter),
			styles.Fill.Type(styles.PatternTypeSolid),
			styles.Fill.Color("#FF7070"),
		),
	))

	xl.SaveAs("./foo.xlsx")
}
