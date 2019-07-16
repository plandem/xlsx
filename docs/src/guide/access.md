# Access
[[toc]]

In general access is possible by indexes (0-based) or by references (alphanumeric presentation, e.g. `A1`). Type of reference depends on situation. 
::: tip Type of reference
- **CellRef** - single cell. E.g. `A1`
- **Ref** - range. E.g. `A1:B2`
- **RefList** - few ranges same time, E.g. `A1:B2 C3 D1:E4`

N.B.: There is special case when **Ref** points to single cell, e.g. `A1:A1` is same as `A1`
:::

Xlsx2Go has few helper types, such as **Bounds** and **BoundsList** to simplify and unify access logic and few helper functions.
```go
package main

import (
	"fmt"
	"github.com/plandem/xlsx/types"
)

func main() {
	//to create Ref from CellRefs
	var ref types.Ref = types.RefFromCellRefs("A1", "B2")
	
	//to create CellRef from indexes
	var cRef types.CellRef = types.CellRefFromIndexes(0, 1)
	
	//to create Bounds from indexes
	var bounds types.Bounds = types.BoundsFromIndexes(0, 1, 1, 2)
}

```

Check [Api Doc](https://godoc.org/github.com/plandem/xlsx/types) for more information about types package

### Sheet
```go
	xl, err := xlsx.Open("./foo.xlsx")
	
	// Get sheet by 0-based index
	sheet := xl.Sheet(0)

	// Create a new sheet:
	sheet = xl.AddSheet("New Sheet Name")
```

### Cell
```go
	// Get cell by 0-based indexes
	cell := sheet.Cell(13, 27)

	// Get cell by reference
	cell = sheet.CellByRef("N28")
```

### Row
```go
	// Get row by 0-based index
	row := sheet.Row(9)

	// Get cell of row at 0-based col index
	cell := row.Cell(0)
```
	
### Column
```go
	// Get col by 0-based index
	col := sheet.Col(3)

	// Get cell of col at 0-based row index
	cell = col.Cell(0)
```

### Range
```go
	// Get range by references
	area := sheet.RangeByRef("D10:E10")
	
	// Get range indexes: fromCol, fromRow, toCol, toRow
	area = sheet.Range(3, 9, 4, 9)
```