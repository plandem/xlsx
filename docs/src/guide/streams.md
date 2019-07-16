# Streaming
[[toc]]

In some cases, when input file is really huge, more advanced techniques required to process data. Streaming is a most commonly used way to lazy load and process by blocks to:
 * reduce overall memory usage and 
 * reduce time of initial loading 

Full stream support of xlsx files can be challenging task. More over, in real world only sheets holds huge data to process. Taking this into account, the real world task will be to stream sheets, but not xlsx file entirely.   

::: warning Limits
Xlsx2Go supports only limited set of features during sheet streaming.
:::

### Reading sheet
To stream sheet, you should open it in `Stream` mode. After that you can access cells in sheet in a normal way.

::: note N.B.
Sheets that were opened as stream, should be closed to free allocated resources.
:::
```go
	sheet := xl.Sheet(0, xlsx.SheetModeStream)
	
	//close sheet to free allocated resources
	defer sheet.Close()
```

::: warning Stream Mode
If sheet previously was opened in normal mode, then you can't open it in stream mode. The reason is simple - streaming is for optimization, if information was totally loaded, then nothing to optimize. 
:::

::: tip Merged cells, Hyperlinks and etc
By default, there is no access to merged cells, hyperlinks and conditional formatting information, to get access you should open sheet in `Multi Phase` mode. In that case sheet will be processed few times - first one to load `later` information, such as merged cells, and last phase for actual streaming cells info.
::: 

```go
	sheet := xl.Sheet(0, xlsx.SheetModeStream, xlsx.SheetModeMultiPhase)
	defer sheet.Close()
```

#### Example 
<<< @/src/code/stream_read_test.go