package code

import (
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/types/options/column"
)

func Example_richText() {
	xl := xlsx.New()

	sheet := xl.AddSheet("")
	sheet.Col(1).SetOptions(
		options.New(
			options.Width(48),
		),
	)

	//Text with bold and italic parts
	sheet.CellByRef("B2").SetText(
		"Text with ",
		styles.New(
			styles.Font.Bold,
		),
		"bold",
		" and ",
		styles.New(
			styles.Font.Italic,
		),
		"italic",
		" parts",
	)

	//Text with red and green parts
	sheet.CellByRef("B4").SetText(
		"Text with ",
		styles.New(
			styles.Font.Color("#ff0000"),
		),
		"red", " and ",
		styles.New(
			styles.Font.Color("#00ff00"),
		),
		"green",
		" parts",
	)

	//Centered text with underlined part
	sheet.CellByRef("B6").SetText(
		"Centered text with ",
		styles.New(
			styles.Font.Underline(styles.UnderlineTypeSingle),
		),
		"underlined",
		" part",
		styles.New(
			styles.Alignment.HAlign(styles.HAlignCenter),
		),
	)

	//E=mc2
	sheet.CellByRef("B8").SetText(
		"E=mc",
		styles.New(
			styles.Font.VAlign(styles.FontVAlignSuperscript),
		),
		"2",
		styles.New(
			styles.Alignment.HAlign(styles.HAlignRight),
		),
	)

	xl.SaveAs("./foo.xlsx")
}
