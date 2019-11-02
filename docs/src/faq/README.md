# Frequently Asked Questions

#### How to ignore sheet's dimension?
```go
	sheet = xl.Sheet(0, SheetModeIgnoreDimension)
```

### Is there an “AutoFit” option for columns and rows?
Unfortunately, there is no way to specify “AutoFit” for a column or row. This feature is only available at runtime from within Excel.