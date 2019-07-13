// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

//Excel has some hardcoded limits:
// https://support.office.com/en-us/article/excel-specifications-and-limits-1672b34d-7043-467e-8e27-269d656771c3

//ExcelCellLimit - total number of characters that a cell can contain
const ExcelCellLimit = 32767

//ExcelRowLimit - total number of rows on a worksheet
const ExcelRowLimit = 1048576

//ExcelColumnLimit - total number of columns on a worksheet
const ExcelColumnLimit = 16384

//ExcelColumnWidthLimit - maximum column width
const ExcelColumnWidthLimit = 255

//ExcelRowWidthLimit - maximum row height
const ExcelRowWidthLimit = 409

//ExcelHeaderFooterLimit - total number of characters that a header/footer can contain
const ExcelHeaderFooterLimit = 255

//ExcelHyperlinkLimit - total number of hyperlinks in a worksheet
const ExcelHyperlinkLimit = 66530

//ExcelSheetNameLimit - total number of characters that a sheet name can contain
const ExcelSheetNameLimit = 31

//ExcelFormulaLimit - total number of characters that a cell formula can contain
const ExcelFormulaLimit = 255

//UrlLimit - total number of characters that an url can contain
// https://stackoverflow.com/questions/417142/what-is-the-maximum-length-of-a-url-in-different-browsers
const UrlLimit = 2000

//FilePathLimit - total number of characters that a file path can contain
// http://msdn.microsoft.com/en-us/library/aa365247(VS.85).aspx#maxpath
const FilePathLimit = 32767
