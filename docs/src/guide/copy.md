# Copying
Xlsx2Go has support of copying cell's information, such as value and styles.

::: warning Merged Cells, Hyperlinks, Comments
Information about merged cells, hyperlinks or comments will not be copied
:::

### Rows
```go
	// Copy row at index 0 to index 5 without copying row's settings
	sheet.Row(0).CopyTo(4, false)
```

### Cols
```go
	// Copy column at index 0 to index 5 without copying column's settings
	sheet.Col(0).CopyTo(3, false)
```

### Ranges
```go
	// Copy range to another range that started at indexes
	sheet.RangeByRef("A1:B3").CopyTo(3, 0)

	// Copy range to another range that started at reference
	sheet.RangeByRef("A1:B3").CopyToRef("I4")
```
