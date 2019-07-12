# Introduction
Why another library to work with Excel XLSX in GO? 

Truth be told, developing of any library starts with some personal goals of author. Someone wants simple library to read Excel files, someone wants to create a new file, other wants to add charts. 

So what were the goals that time? It's a great pity, but I could not get a library that:

1) respects existing data/formatting - no corrupted files or lost formatting
> What if I need to open a well formatted file created with my favorite desktop application and update only one value?! I must get almost same file with just one updated value. None of existing library was able to do it. Corrupted file or lost formatting is common issue.

2) works with big files - reasonable speed and memory footprint
> Same here, someone could not open, others took forever to open with anomaly memory usage.

3) consistent and as small API as possible with enough features set to do most common tasks - learning curve means something 
> Why?! Because it's not rocket science - open/create file, create/read/update/delete sheets/rows/cols and use styles. XLSX is quite simple format to read/write and GO has quite powerful xml encoder/decoder, so the hardest part - that API. 

4) easy to read/understand source code, easy to maintain, easy to contribute - no shadow places/hacks/magic, just read and understand
> I was trying to contribute to existing libraries, but...actually it's much faster to create it from ground zero than to refactor existing and get satisfied results or fix some issues.

## Benchmarks
It was not a goal to make best of the best, but the same time it's interesting to know pros/cons. 
For some cases this library is second, for other - best, but in case of reading huge files - **the only**. 

|                | tealeg | excelize | xlsx |
|----------------|:------:|:--------:|:----:|
| RandomGet      |   1*   |     3    |   2  |
| RandomSet      |   1*   |     3    |   2  |
| RandomSetStyle |   1*   |     3    |   2  |
| ReadBigFile    |   2    |     3    |   1  |
| UpdateBigFile  |   2**  |     3    |   1  |
| ReadHugeFile   |   -    |     -    |   1  |
| UpdateHugeFile |   -    |     -    |   1  |

\* does not mutate information directly, so faster get/set, but slower read/write files - sometimes it can take forever to open file.

\** corrupted file after saving, lost styles/formatting

::: tip
You get more info about benchmarks at dedicated [Benchmarks Repository](https://github.com/plandem/xlsx-benchmarks)  
:::

## Roadmap
- [x] [copy cell, row, col](/guide/copy.md)
- [x] [read as stream](/guide/stream-read.md)
- [x] [merged cells](/guide/merged-cells.md)
- [x] [hyperlinks](/guide/hyperlinks.md)
- [x] [comments](/guide/comments.md)
- [x] [typed getter/setter for values](/guide/typed-values.md)
- [x] [conditional formatting](/guide/conditional-formatting.md)
- [x] [styles formatting](/guide/styles-formatting.md)
- [x] [rich text](/guide/rich-text.md)
- [ ] copy sheet
- [ ] write as stream
- [ ] formulas
- [ ] drawing
- [ ] unpack package to temp folder to reduce memory usage

## API Documentation
For detailed API documentation, you can check [godoc.org](https://godoc.org/github.com/plandem/xlsx)

### OOXML edition
XML is compliant with part 1 of the [5th edition](http://www.ecma-international.org/publications/standards/Ecma-376.htm) of the ECMA-376 Standard for Office Open XML

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fplandem%2Fxlsx.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fplandem%2Fxlsx?ref=badge_large)

