package code

import (
	"github.com/plandem/xlsx/format/styles"
)

func Example_stylesFont() {
	//all possible settings for font
	styles.New(
		styles.Font.Name("Courier New"),
		styles.Font.Bold,
		styles.Font.Italic,
		styles.Font.Strikeout,
		styles.Font.Superscript,
		styles.Font.Subscript,
		styles.Font.Shadow,
		styles.Font.Condense,
		styles.Font.Extend,
		styles.Font.Family(styles.FontFamilyRoman),
		styles.Font.Color("#FF0000"),
		styles.Font.Size(16),
		styles.Font.Underline(styles.UnderlineTypeSingle),
		styles.Font.Scheme(styles.FontSchemeMinor),
		styles.Font.Charset(styles.FontCharsetMAC),
	)
}
