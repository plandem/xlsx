package internal

//Excel has some hardcoded limits:
// https://support.office.com/en-us/article/excel-specifications-and-limits-1672b34d-7043-467e-8e27-269d656771c3

//Total number of characters that a cell can contain
const ExcelCellLimit = 32767

//Total number of rows on a worksheet
const ExcelRowLimit = 1048576

//Total number of columns on a worksheet
const ExcelColumnLimit = 16384

//Maximum column width
const ExcelColumnWidthLimit = 255

//Maximum row height
const ExcelRowWidthLimit = 409

//Total number of characters that a header/footer can contain
const ExcelHeaderFooterLimit = 255

//Total number of hyperlinks in a worksheet
const ExcelHyperlinkLimit = 66530

//Total number of characters that a url can contain
const ExcelUrlLimit = 255

//Total number of characters that a sheet name can contain
const ExcelSheetNameLimit = 31
