package code

import (
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/types/hyperlink"
)

func Example_hyperlinks() {
	xl := xlsx.New()

	sheet := xl.AddSheet("")

	//hyperlink to other excel file with reference to C3 at Sheet1
	link := hyperlink.New(
		hyperlink.ToFile(`C:\example_simple.xlsx`),
		hyperlink.ToRef("C3", "Sheet1"),
		hyperlink.Tooltip("That's a tooltip"),
		hyperlink.Display("Something to display"), //Cell still holds own value
		hyperlink.Styles(styles.New(
			styles.Font.Bold,
			styles.Font.Color("#ff0000"),
		)),
	)

	sheet.CellByRef("A2").SetHyperlink(link)

	xl.SaveAs("./foo.xlsx")
}
