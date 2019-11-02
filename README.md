# Xlsx2Go
[![Build Status](https://travis-ci.org/plandem/xlsx.svg?branch=master)](https://travis-ci.org/plandem/xlsx)
[![Code Coverage](https://codecov.io/gh/plandem/xlsx/branch/master/graph/badge.svg)](https://codecov.io/gh/plandem/xlsx) 
[![Go Report Card](https://goreportcard.com/badge/github.com/plandem/xlsx)](https://goreportcard.com/report/github.com/plandem/xlsx)
[![GoDoc](https://godoc.org/github.com/plandem/xlsx?status.svg)](https://godoc.org/github.com/plandem/xlsx)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/plandem/xlsx/master/LICENSE) 
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fplandem%2Fxlsx.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fplandem%2Fxlsx?ref=badge_shield)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.me/gayvoronsky)

```go
package main

import (
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/conditional"
	"github.com/plandem/xlsx/format/conditional/rule"
	"github.com/plandem/xlsx/format/styles"
)

func main() {
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
```

## Documentation
* [Guide](https://plandem.github.io/xlsx/)
* [API Documentation](https://godoc.org/github.com/plandem/xlsx)
* [Benchmarks](https://github.com/plandem/xlsx-benchmarks)

# Roadmap
- [ ] sheet: copy
- [ ] sheet: custom filters
- [x] sheet: streaming
- [x] merged cells: merge/split for ranges, cols, rows
- [x] hyperlinks: for cells, ranges, cols, rows
- [x] range: copy
- [x] row: copy
- [x] col: copy
- [x] cell: comments
- [ ] cell: formulas
- [x] cell: typed getter/setter for values
- [x] other: conditional formatting
- [x] other: rich texts
- [ ] other: drawing
- [ ] other: more optimization
- [ ] other: more tests

# Contribution 
- To prevent mess, sources have strict separation of markup and functionality. Document that describes OOXML is quite huge (about 6K pages), but the same time - functionality is not.
- All markup resides inside of 'ml' folders, only marshal/unmarshal is allowed here, no any functionality.
- Not every 'ml object' has related 'functional object' and vice versa.
- If you want some functionality, then wrap 'ml object' and do what you want.

### OOXML edition
XML is compliant with part 1 of the [5th edition](http://www.ecma-international.org/publications/standards/Ecma-376.htm) of the ECMA-376 Standard for Office Open XML

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fplandem%2Fxlsx.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fplandem%2Fxlsx?ref=badge_large)
