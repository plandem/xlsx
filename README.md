# Xlsx2Go
[![Build Status](https://travis-ci.org/plandem/xlsx2go.svg?branch=master)](https://travis-ci.org/plandem/xlsx2go)
[![Code Coverage](https://codecov.io/gh/plandem/xlsx2go/branch/master/graph/badge.svg)](https://codecov.io/gh/plandem/xlsx2go) 
[![Go Report Card](https://goreportcard.com/badge/github.com/plandem/xlsx2go)](https://goreportcard.com/report/github.com/plandem/xlsx2go)
[![GoDoc](https://godoc.org/github.com/plandem/xlsx2go?status.svg)](https://godoc.org/github.com/plandem/xlsx2go)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/plandem/xlsx2go/master/LICENSE) 
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fplandem%2Fxlsx.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fplandem%2Fxlsx?ref=badge_shield)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.me/gayvoronsky)

**Note:** Github repository was renamed from `xlsx` to `xlsx2go` to make it more easier to distinct existing xlsx libraries. Previous address will be auto redirected, package will be named as before - xlsx.   

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

# Introduction
Why another library to work with Excel XLSX in GO? 

Truth be told, developing of any library starts with some personal goals of author. Someone wants simple library to read Excel files, someone wants to create a new file, other wants to add charts. 

So what were the goals that time? It's a great pity, but I could not get a library that:

1) respects existing data/formatting - no corrupted files or lost formatting

> What if I need to open a well formatted file created with my favorite desktop application and update only one value?! I must get almost same file with just one updated value. None of existing library was able to do it. Corrupted file or lost formatting is common issue.

2) works with big files - reasonable speed and memory footprint

> Same here, someone could not open, others took forever to open with anomaly memory usage.

3) consistent and as small API as possible with enough features set to do most common tasks - learning curve means something 

> Why?! Because it's not rocket science - open/create file, create/read/update/delete sheets/rows/cols and use styles. XLSX is quite simple format to read/write and GO has quite powerful xml encoder/decoder, so the hardest part - that API. 

4) easy to read/understand source code, easy to maintain, easy to contribute - no shadow places/hacks/magic, just read and understand

> I was trying to contribute to existing libraries, but...actually it's much faster to create it from ground zero than to refactor existing and get satisfied results or fix some issues.

## Documentation
* [Guide](https://plandem.github.io/xlsx2go/)
* [API Documentation](https://godoc.org/github.com/plandem/xlsx)

# Benchmarks
It was not a goal to make best of the best, but the same time it's interesting to know pros/cons. 
For some cases this library is second, for other - best, but in case of reading huge files - **the only**. 

|                | tealeg | excelize | xlsx |
|----------------|:------:|:--------:|:----:|
| RandomGet      |   1!   |     3    |   2  |
| RandomSet      |   1!   |     3    |   2  |
| RandomSetStyle |   1!   |     3    |   2  |
| ReadBigFile    |    2   |     3    |   1  |
| UpdateBigFile  |    2!! |     3    |   1  |
| ReadHugeFile   |    -   |     -    |   1  |
| UpdateHugeFile |    -   |     -    |   1  |

* ! - does not mutate information directly, so faster get/set, but slower read/write files - sometimes it can take forever to open file.
* !! - corrupted file after saving, lost styles/formatting
 
[Benchmarks report](BENCHMARKS.md) 

# Roadmap
- [ ] sheet: copy
- [x] sheet: read as stream
- [ ] sheet: custom filters
- [x] sheet: write as stream
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
- [ ] other: unpack package to temp folder to reduce memory usage
- [x] other: more tests

# Contribution 
- To prevent mess, sources have strict separation of markup and functionality. Document that describes OOXML is quite huge (about 6K pages), but the same time - functionality is not.
- All markup resides inside of 'ml' folders, only marshal/unmarshal is allowed here, no any functionality.
- Not every 'ml object' has related 'functional object' and vice versa.
- If you want some functionality, then wrap 'ml object' and do what you want.

### OOXML edition
XML is compliant with part 1 of the [5th edition](http://www.ecma-international.org/publications/standards/Ecma-376.htm) of the ECMA-376 Standard for Office Open XML

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fplandem%2Fxlsx.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fplandem%2Fxlsx?ref=badge_large)
