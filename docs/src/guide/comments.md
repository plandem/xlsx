# Comments
[[toc]]

Cell comments are a way of adding notation to cells. 

Xlsx2Go supports string version of comments, as well as custom version via special type. Read about types below.

```go
	//to add comment to cell
	sheet.CellByRef("N28").SetComment("Never hear the 'No comment' again.")

	//to remove comment
	sheet.CellByRef("N28").RemoveComment()

	//to retrieve text of comment
	comments := sheet.CellByRef("N28").Comment()
```

### String comment
As was shown, the simplest way to add comment is to use string version. At the same time, that version has some limitations - you can't use rich text, set author, width, height and etc.

```go
	//to add string comment
	sheet.CellByRef("B2").SetValue("Any comments?")
	sheet.CellByRef("B2").SetComment("No comment!")
```
![](~@images/comments.png)

### Custom comment
While with string version of comment you can add comments really easy, sometimes we need additional settings like width, height, author or even rich text. For these cases you can use special type and configure comment as you wish.

![](~@images/comments-custom.png)

#### Example

<<< @/src/code/comments_test.go

::: note
Check [Rich Text](/guide/rich-text.md) for more information about rich texts.
:::