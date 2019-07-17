package code

import (
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/conditional"
	"github.com/plandem/xlsx/format/conditional/rule"
	"github.com/plandem/xlsx/format/styles"
)

func Example() {
	xl := xlsx.New()
	defer xl.Close()

	//create a new sheet
	sheet := xl.AddSheet("The first sheet")

	//access by ref
	cell := sheet.CellByRef("A2")

	//set value
	cell.SetValue("Easy Peasy")

	//set cool styles
	cell.SetStyles(styles.New(
		styles.Font.Bold,
		styles.Font.Color("#ff0000"),
		styles.Fill.Type(styles.PatternTypeSolid),
		styles.Fill.Color("#ffff00"),
		styles.Border.Color("#009000"),
		styles.Border.Type(styles.BorderStyleMedium),
	))

	//add comment
	cell.SetComment("No Comment!")

	//add hyperlink
	sheet.CellByRef("A4").SetValueWithHyperlink("wikipedia", "http://google.com")

	//merge cells
	sheet.RangeByRef("A6:A7").Merge()
	sheet.CellByRef("A6").SetValue("merged cell")

	//iterating
	for iRow := 1; iRow < 7; iRow++ {
		//access by indexes
		cell := sheet.Cell(1, iRow)
		cell.SetValue(iRow)
	}

	//add conditional formatting
	sheet.AddConditional(conditional.New(
		conditional.AddRule(
			rule.Value.Between(1, 3, styles.New(
				styles.Font.Bold,
				styles.Font.Color("#ff0000"),
			)),
		),
		conditional.AddRule(
			rule.IconSet.Type(rule.IconSetType3Arrows),
		),
	), "B2:B7")

	xl.SaveAs("./foo.xlsx")
}
