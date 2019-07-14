# Rich Text
Rich text is text with multiply style formatting. The rule oh thumb is to break the string into parts and put a style object before the part that you want to format. String parts that don’t have a style are given a default styles.  

For example to get 'This is text with **bold** part':

```go  
// unformatted string
'This is text with bold part'

// break it into parts
'This is text with ', 'bold', ' part'

// add styles before the parts that should be formatted
'This is text with ', styles, 'bold', ' part'
```

::: danger
Excel does not allow the use of two consecutive formats in a rich text or an empty string part. 
:::

::: tip Styling Limits
In Excel, only the font properties of the styles are applied to the string parts in a rich text. Other features such as border, background, text wrap and alignment must be applied to the cell. You can add `last argument` with `styles for cell`
:::

::: warning Text limits
Excel has built-in limits for cell's text value. Check [Excel Limits](/guide/limits.md) for more information about built-in limits
:::

### Example
![](/rich-text.png)

<<< @/src/code/rich_text_test.go