# Iterate
Does not matter how you will access data inside of sheet(by indexes, reference or via iterators), in all cases library will auto expand sheet's dimension to required size.
```go
	xl := xlsx.New()
	sheet := xl.AddSheet("sheet name")

	// Append 99 rows, 100 in total
	sheet.Row(99)
	
	// Append 49 cols, 50 in total
	sheet.Col(49)
```
Sheet is a grid, so if you request something outside of current dimension of grid, library will auto expand to required dimension (new sheets have 1x1 dimension by default).

### Iterate by indexes
```go
	totalCols, totalRows := sheet.Dimension()
	for rIdx := 0; rIdx < totalRows; rIdx++ {
		for cIdx := 0; cIdx < totalCols; cIdx++ {
			cell := sheet.Cell(cIdx, rIdx)
		}
	}
```

### Iterators
Sometimes it's easier to iterate through the data using iterators

N.B.: Internally, Row and Column are subtypes of Range, so they inherit Range's functionality. 

#### Sheet
```go
	xl, err := xlsx.Open("./foo.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	// Iterate sheets via iterator
	for sheets := xl.Sheets(); sheets.HasNext(); {
		_, sheet := sheets.Next()
	}
```

#### Rows
```go
	// Iterate rows via iterator
	for rows := sheet.Rows(); rows.HasNext(); {
		_, row := rows.Next()
		
		// Iterate row's cells via iterator
		for cells := row.Cells(); cells.HasNext(); {
			_, _, cell := cells.Next()
		}
	}
```

#### Columns
```go
	// Iterate cols via iterator
	for cols := sheet.Cols(); cols.HasNext(); {
		_, col := cols.Next()
		// Iterate col's cells via iterator
		for cells := col.Cells(); cells.HasNext(); {
			_, _, cell := cells.Next()
		}
	}
```

#### Range
```go
	// Iterate range's cells via iterator
	for cells := sheet.RangeByRef("A1:B3").Cells(); cells.HasNext(); {
		_, _, cell := cells.Next()
	}
```

### Walk
Or you can `Walk` method to process cells in range
```go
	//Walk through the cells of row
	sheet.Row(0).Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		fmt.Println(c.Value())
	})

	// Walk through the cells of col
	sheet.Col(0).Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		fmt.Println(c.Value())
	})

	// Walk through the cells of range
	sheet.RangeByRef("A1:B3").Walk(func(idx, cIdx, rIdx int, c *xlsx.Cell) {
		fmt.Println(c.Value())
	})
```