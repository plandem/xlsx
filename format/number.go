package format

import (
	"github.com/plandem/xlsx/internal/ml"
)

//NumberFormat is option to update StyleFormat with provided custom format of number
func NumberFormat(format string) func(*StyleFormat) {
	return func(s *StyleFormat) {
		s.NumFormat.ID = -1
		s.NumFormat.Code = format
	}
}

//NumberFormatID is option to update StyleFormat with provided id of number format
func NumberFormatID(id int) func(*StyleFormat) {
	return func(s *StyleFormat) {
		s.NumFormat.ID = id
		s.NumFormat.Code = ""
	}
}

var (
	BuiltInNumberFormat map[int]*ml.NumberFormat
)

func init() {

}
/*
package format

import (
"encoding/xml"
. "github.com/plandem/xlsx/format/internal/number_format"
)

//NumberFormat is a direct mapping of XSD CT_NumFmt
type NumberFormat struct {
	ID   int    `xml:"numFmtId,attr"`
	Code string `xml:"formatCode,attr"`

	//Type is internal type of NumberFormat for built-in formats, used by Cell for custom logic
	Type Type
}

var (
	builtInNumberFormat map[int]*NumberFormat
)

//NumberFormatCode is option to update StyleFormat with provided custom format of number
func NumberFormatCode(code string) func(*StyleFormat) {
	return func(s *StyleFormat) {
		s.NumFormat.ID = -1
		s.NumFormat.Code = code
	}
}

//NumberFormatID is option to update StyleFormat with provided id of number format
func NumberFormatID(id int) func(*StyleFormat) {
	return func(s *StyleFormat) {
		s.NumFormat.ID = id
		s.NumFormat.Code = ""
	}
}

func (nf *NumberFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := d.DecodeElement(nf, &start)

	if err == nil {
		nf.resolve()
	}

	return err
}

//IsBuiltIn returns true if number format is one of built-in
func (nf *NumberFormat) IsBuiltIn() bool {
	return nf.ID >= 0 && nf.ID <= LastReservedID
}
//
//func (nf *NumberFormat) ResolveCode(code string) int {
//
//}

//resolve resolve id and code
func (nf *NumberFormat) resolve() {
	if nf.IsBuiltIn() {
		if knownFormat, ok := builtInNumberFormat[nf.ID]; ok {
			//known built-in format?
			*nf = *knownFormat
		} else {
			//general built-in format?
			nf.Code = builtInNumberFormat[0x00].Code
			nf.Type = General
		}
	} else {
		//if there is a known format code, then use that built-in type instead of code
		for _, knownFormat := range builtInNumberFormat {
			if knownFormat.Code == nf.Code {
				*nf = *knownFormat
				break
			}
		}
	}
}

func init() {
	builtInNumberFormat = map[int]*NumberFormat{
		0x00: {0x00, `@`, General},
		0x01: {0x01, `0`, Integer},
		0x02: {0x02, `0.00`, Float},
		0x03: {0x03, `#,##0`, Float},
		0x04: {0x04, `#,##0.00`, Float},
		0x05: {0x05, `($#,##0_);($#,##0)`, Float},
		0x06: {0x06, `($#,##0_);[RED]($#,##0)`, Float},
		0x07: {0x07, `($#,##0.00_);($#,##0.00_)`, Float},
		0x08: {0x08, `($#,##0.00_);[RED]($#,##0.00_)`, Float},
		0x09: {0x09, `0%`, Integer},
		0x0a: {0x0a, `0.00%`, Float},
		0x0b: {0x0b, `0.00E+00`, Float},
		0x0c: {0x0c, `# ?/?`, Float},
		0x0d: {0x0d, `# ??/??`, Float},
		0x0e: {0x0e, `m-d-yy`, Date},
		0x0f: {0x0f, `d-mmm-yy`, Date},
		0x10: {0x10, `d-mmm`, Date},
		0x11: {0x11, `mmm-yy`, Date},
		0x12: {0x12, `h:mm AM/PM`, Time},
		0x13: {0x13, `h:mm:ss AM/PM`, Time},
		0x14: {0x14, `h:mm`, Time},
		0x15: {0x15, `h:mm:ss`, Time},
		0x16: {0x16, `m-d-yy h:mm`, DateTime},

		0x25: {0x25, `(#,##0_);(#,##0)`, Integer},
		0x26: {0x26, `(#,##0_);[RED](#,##0)`, Integer},
		0x27: {0x27, `(#,##0.00);(#,##0.00)`, Float},
		0x28: {0x28, `(#,##0.00);[RED](#,##0.00)`, Float},
		0x29: {0x29, `_(*#,##0_);_(*(#,##0);_(*"-"_);_(@_)`, Float},
		0x2a: {0x2a, `_($*#,##0_);_($*(#,##0);_(*"-"_);_(@_)`, Float},
		0x2b: {0x2b, `_(*#,##0.00_);_(*(#,##0.00);_(*"-"??_);_(@_)`, Float},
		0x2c: {0x2c, `_($*#,##0.00_);_($*(#,##0.00);_(*"-"??_);_(@_)`, Float},
		0x2d: {0x2d, `mm:ss`, DeltaTime},
		0x2e: {0x2e, `[h]:mm:ss`, DeltaTime},
		0x2f: {0x2f, `mm:ss.0`, DeltaTime},
		0x30: {0x30, `##0.0E+0`, Float},
		0x31: {0x31, `@`, General},
	}
}

*/