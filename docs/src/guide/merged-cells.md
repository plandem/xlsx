# Merged Cells
[[toc]]

Short example is better any words

![](/merged-cells.png)

<<< @/src/code/merged_cells_test.go


To split previously merged cells:
```go
	sheet.RangeByRef("B4:E4").Split()
```
 
### Merge Rows
In case if you need to merge few rows:
```go
	//Merge 1-9 rows, internally it uses range: [0, 1, ExcelRowLimit, 9]
	sheet.MergeRows(1, 9)

	//Split 1-9 rows
	sheet.SplitRows(1, 9)
```

### Merge Cols
In case if you need to merge few columns:
```go
	//Split 1-9 cols, internally it uses range: [1, 0, 9, ExcelColumnLimit]
	sheet.MergeCols(1, 9)

	//Split 1-9 cols
	sheet.SplitCols(1, 9)
```

::: warning Columns and Rows limits
Excel has built-in limits for total number of rows and columns per sheet. Check [Excel Limits](/guide/limits.md) for more information about built-in limits
:::