# Styles Formatting
[[toc]]

Styles can be defined through special type - general type for all available styles. But information that will be using to style object, depends on usage, e.g. to style cells everything will be use, to style rich texts - only font information. 

```go
	// create a new styles
	ss := styles.New(
		styles.Font.Bold,
		styles.Font.Color("#ff0000"),
	)
	
	//update styles
	ss.Set(
		styles.Border.Color("#009000"),
		styles.Border.Type(styles.BorderStyleMedium),
	)
```
::: warning Modify Styles
While you can modify created styles, you should keep in mind, that modifying will work only before applying styles to cell and any modifications after applying, will create new styles.
:::

```go
	ss := styles.New(
		styles.Font.Bold,
	))

	//font will be `bold`
	sheet.CellByRef("A1").SetStyles(ss)

	//modify styles
	ss.Set(
		styles.Font.Color("#ff0000"),
	)

	//`A2` - will be `bold and red` 
	//`A1` - will be only `bold` and without color
	sheet.CellByRef("A2").SetStyles(ss)
```

::: note Predefined values
Xlsx2Go defined all built-in values to use for styling. For more information, check [API documentation](https://godoc.org/github.com/plandem/xlsx/format/styles#pkg-constants)
::: 

### Font
::: warning
Excel can only display installed fonts, that's why using standard fonts(e.g.: `Calibri`, `Times New Roman` or `Courier New`) is highly recommended.
:::

::: note
The default font for cell is `Calibri` (Excel 2007+)
:::
<<< @/src/code/styles_font_test.go

### Fill

### Border

### Alignment

### Number Format

### Protection