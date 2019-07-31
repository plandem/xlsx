# Number Format Codes
[[toc]]
Number format code is way to show numeric value. It controls whether a number is displayed as an integer, a floating number, a date, a currency value or some other user defined format.

::: tip
Xlsx2Go supports built-in Excel codes and will try to detect/convert to proper internal ID, as well as custom codes - use it as is without worrying.
:::

### Built-in Codes
::: note General type
* `@`
* `General`
:::

::: note Integer number	
* `0`
* `0%`
* `(#,##0_);(#,##0)`
* `(#,##0_);[RED](#,##0)`
:::		

::: note Float number
* `0.00`
* `#,##0`
* `#,##0.00`
* `($#,##0_);($#,##0)`
* `($#,##0_);[RED]($#,##0)`
* `($#,##0.00_);($#,##0.00_)`
* `($#,##0.00_);[RED]($#,##0.00_)`
* `0.00%`
* `0.00E+00`
* `# ?/?`
* `# ??/??`
* `(#,##0.00);(#,##0.00)`
* `(#,##0.00);[RED](#,##0.00)`
* `_(*#,##0_);_(*(#,##0);_(*"-"_);_(@_)`
* `_($*#,##0_);_($*(#,##0);_(*"-"_);_(@_)`
* `_(*#,##0.00_);_(*(#,##0.00);_(*"-"??_);_(@_)`
* `_($*#,##0.00_);_($*(#,##0.00);_(*"-"??_);_(@_)`
* `##0.0E+0`
:::

::: note Date
* `m-d-yy`
* `d-mmm-yy`
* `d-mmm`
* `mmm-yy`	
:::		

::: note Time
* `h:mm AM/PM`
* `h:mm:ss AM/PM`
* `h:mm`
* `h:mm:ss`	 
:::		 

::: note Date+Time
* `m-d-yy h:mm`
:::

::: note DeltaTime
* `mm:ss`
* `[h]:mm:ss`
* `mm:ss.0`
:::

### Custom Codes
Format code can control any aspect of number formatting allowed by Excel:

::: tip Currency
The `$` in format appears as the local currency symbol.
:::

::: tip Colors
The color format should have one of the following values: 

`[Black]` `[Blue]` `[Cyan]` `[Green]` `[Magenta]` `[Red]` `[White]` `[Yellow]`
:::

#### Examples
<table>
<tr>
	<th style="text-align: right;">Number code</th>
	<th style="text-align: right;">Go Value</th>
	<th style="text-align: right;">Excel Output</th>
</th>
<tr>
	<td style="text-align: right;"><code>dd/mm/yyyy hh:mm AM/PM</code></td>
	<td style="text-align: right;">time.Now()</td>
	<td style="text-align: right;">18/07/2019 12:30 AM</td>
</tr>
<tr>
	<td style="text-align: right;"><code>mm/dd/yy</code></td>
	<td style="text-align: right;">time.Now()</td>
	<td style="text-align: right;">07/18/19</td>
</tr>
<tr>
	<td style="text-align: right;"><code>mmm d yyyy</code></td>
	<td style="text-align: right;">time.Now()</td>
	<td style="text-align: right;">Jul 18 2019</td>
</tr>
<tr>
	<td style="text-align: right;"><code>d mmmm yyyy</code></td>
	<td style="text-align: right;">time.Now()</td>
	<td style="text-align: right;">18 July 2019</td>
</tr>
<tr>
	<td style="text-align: right;"><code>0.000</code></td>
	<td style="text-align: right;">1.2345678</td>
	<td style="text-align: right;">1.235</td>
</tr>
<tr>
	<td style="text-align: right;"><code>#,##0</code></td>
	<td style="text-align: right;">1234.567</td>
	<td style="text-align: right;">1,235</td>
</tr>
<tr>
	<td style="text-align: right;"><code>0 "dollar and" .00 "cents"</code></td>
	<td style="text-align: right;">1.87</td>
	<td style="text-align: right;">1 dollar and .87 cents</td>
</tr>
<tr>
	<td style="text-align: right;" rowspan="3"><code>[Green]General;[Red]-General;General</code></td>
	<td style="text-align: right;">12345</td>
	<td style="text-align: right;"><span style="color:green">1235</span></td>
</tr>
<tr>
	<td style="text-align: right;">-12345</td>
	<td style="text-align: right;"><span style="color:red">-12345</span></td>
</tr>
<tr>
	<td style="text-align: right;">0</td>
	<td style="text-align: right;">0</td>
</tr>
</table>

::: note

::: right
For more information about custom formats, check [Microsoft Documentation](https://support.office.com/en-us/article/create-a-custom-number-format-78f2a361-936b-4c03-8772-09fab54be7f4?ui=en-US&rs=en-US&ad=US)
:::

