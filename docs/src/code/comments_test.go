package code

import (
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/types/comment"
)

func Example_comments() {
	xl := xlsx.New()

	sheet := xl.AddSheet("")

	//red alert box comment
	cmt := comment.New(
		comment.Width(250),
		comment.Height(50),
		comment.Background("#FFAAAA"),
		comment.Shadow("#FF0000"),
		comment.Stroke("#CC0000"),
		comment.Author("Gate Keeper"),
		comment.Text(
			styles.New(
				styles.Font.Bold,
				styles.Font.Size(16),
			),
			"STOP",
			"\nDanger zone, do not proceed",
		),
	)

	sheet.CellByRef("A2").SetValue("Hello?")
	sheet.CellByRef("A2").SetComment(cmt)

	xl.SaveAs("./foo.xlsx")
}
