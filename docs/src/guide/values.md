# Values
[[toc]]

Working with cells' values is dead simple and in most cases it will be more than enough:
```go
	// set value, using unified method
	sheet.CellByRef("A1").SetValue("string")
	sheet.CellByRef("A2").SetValue(true)
	sheet.CellByRef("A3").SetValue(12345)
	sheet.CellByRef("A4").SetValue(123.123)
	sheet.CellByRef("A5").SetValue(time.Now())

	// get raw value
	var val string = sheet.CellByRef("A1").Value()
	
	// get string presentation of value
    val = sheet.CellByRef("A1").String()

```
::: tip Unification 
Xlsx2Go will automagically detect type of value and call required typed getter/setter. Read below about typed values.
:::

::: note Raw value
Raw value is cell's value that returned as is - without any processing. It can be anything - reference to `Shared String` or even error value.
:::

### Numbers 
```go
	// set signed integer value
	sheet.CellByRef("A1").SetInt(-12345)
	// set unsigned integer value
	sheet.CellByRef("A2").SetUint(12345)
	// set floating number value
	sheet.CellByRef("A3").SetFloat(12345.12345)

	// get signed integer value or error
	si, err := sheet.CellByRef("A1").Int()
	// get unsigned integer value or error
	ui, err := sheet.CellByRef("A2").Uint()
	// get floating number value or error
	f, err := sheet.CellByRef("A3").Float()
```

### Dates
```go
	now := time.Now()
	
	// set date+time value
	sheet.CellByRef("A1").SetDateTime(now)
	//set date value
	sheet.CellByRef("A2").SetDate(now)
	//set time value
	sheet.CellByRef("A3").SetTime(now)
	//set delta time value
	sheet.CellByRef("A4").SetDeltaTime(now)
	
	// get time.Time value or error 
	d, err := sheet.CellByRef("A1").Date()
```

::: note N.B.
Technically, any date or time related value is stored as number with additional format code information
:::

### Booleans
```go
	// set boolean value
	sheet.CellByRef("A1").SetBool(true)

	// get boolean value or error
	b, err := sheet.CellByRef("A1").Bool()
``` 

### Texts
Excel uses mechanism to reduce required memory while working with texts, due to fact that some text values can be repeated as is multiply times. So when string will be used more than once, Excel stores it as `Shared String` and in that case does not matter how many times user used that string, only one will be stored in memory and cell will hold only reference to that string. But for some cases, user wants to enforce Excel to store text directly in cell and for these cases Excel stores text as `Inline String`.

Xlsx2Go supports both types and user should decide by own, what type to use.  
```go
	// set text as `Shared String`
	sheet.CellByRef("A1").SetText("string")

	// set text as `Inline String`
	sheet.CellByRef("A2").SetInlineText("inline string")

	// get text value
	s := sheet.CellByRef("A1").String()
```

::: note Rich Text
Keep in mind, that text can be a simple string, as well as full featured rich text.
 
::: right
Check [Rich Text](/guide/rich-text.md) for more information about rich texts.
:::

### Number Formats
In some cases we want to set value, but also use format how to display that value.
 
```go
	//set floating number value and related format code
	sheet.CellByRef("A1").SetValueWithFormat(12345.12345, "0.00")
```

::: tip Format Codes
Check [Number Format](/guide/number_format.md) for more information about codes for number format.
:::