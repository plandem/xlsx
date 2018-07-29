package numberFormat

import (
	"github.com/plandem/xlsx/internal/ml"
)

var (
	builtIn     map[int]*builtInFormat
	typeDefault map[Type]int
)

func init() {
	typeDefault = map[Type]int{
		General:   0x00,
		Integer:   0x01,
		Float:     0x02,
		Date:      0x0e,
		Time:      0x14,
		DateTime:  0x16,
		DeltaTime: 0x2d,
	}

	builtIn = map[int]*builtInFormat{
		0x00: {ml.NumberFormat{ID: 0x00, Code: `@`}, General},
		0x01: {ml.NumberFormat{ID: 0x01, Code: `0`}, Integer},
		0x02: {ml.NumberFormat{ID: 0x02, Code: `0.00`}, Float},
		0x03: {ml.NumberFormat{ID: 0x03, Code: `#,##0`}, Float},
		0x04: {ml.NumberFormat{ID: 0x04, Code: `#,##0.00`}, Float},
		0x05: {ml.NumberFormat{ID: 0x05, Code: `($#,##0_);($#,##0)`}, Float},
		0x06: {ml.NumberFormat{ID: 0x06, Code: `($#,##0_);[RED]($#,##0)`}, Float},
		0x07: {ml.NumberFormat{ID: 0x07, Code: `($#,##0.00_);($#,##0.00_)`}, Float},
		0x08: {ml.NumberFormat{ID: 0x08, Code: `($#,##0.00_);[RED]($#,##0.00_)`}, Float},
		0x09: {ml.NumberFormat{ID: 0x09, Code: `0%`}, Integer},
		0x0a: {ml.NumberFormat{ID: 0x0a, Code: `0.00%`}, Float},
		0x0b: {ml.NumberFormat{ID: 0x0b, Code: `0.00E+00`}, Float},
		0x0c: {ml.NumberFormat{ID: 0x0c, Code: `# ?/?`}, Float},
		0x0d: {ml.NumberFormat{ID: 0x0d, Code: `# ??/??`}, Float},
		0x0e: {ml.NumberFormat{ID: 0x0e, Code: `m-d-yy`}, Date},
		0x0f: {ml.NumberFormat{ID: 0x0f, Code: `d-mmm-yy`}, Date},
		0x10: {ml.NumberFormat{ID: 0x10, Code: `d-mmm`}, Date},
		0x11: {ml.NumberFormat{ID: 0x11, Code: `mmm-yy`}, Date},
		0x12: {ml.NumberFormat{ID: 0x12, Code: `h:mm AM/PM`}, Time},
		0x13: {ml.NumberFormat{ID: 0x13, Code: `h:mm:ss AM/PM`}, Time},
		0x14: {ml.NumberFormat{ID: 0x14, Code: `h:mm`}, Time},
		0x15: {ml.NumberFormat{ID: 0x15, Code: `h:mm:ss`}, Time},
		0x16: {ml.NumberFormat{ID: 0x16, Code: `m-d-yy h:mm`}, DateTime},
		//...
		0x25: {ml.NumberFormat{ID: 0x25, Code: `(#,##0_);(#,##0)`}, Integer},
		0x26: {ml.NumberFormat{ID: 0x26, Code: `(#,##0_);[RED](#,##0)`}, Integer},
		0x27: {ml.NumberFormat{ID: 0x27, Code: `(#,##0.00);(#,##0.00)`}, Float},
		0x28: {ml.NumberFormat{ID: 0x28, Code: `(#,##0.00);[RED](#,##0.00)`}, Float},
		0x29: {ml.NumberFormat{ID: 0x29, Code: `_(*#,##0_);_(*(#,##0);_(*"-"_);_(@_)`}, Float},
		0x2a: {ml.NumberFormat{ID: 0x2a, Code: `_($*#,##0_);_($*(#,##0);_(*"-"_);_(@_)`}, Float},
		0x2b: {ml.NumberFormat{ID: 0x2b, Code: `_(*#,##0.00_);_(*(#,##0.00);_(*"-"??_);_(@_)`}, Float},
		0x2c: {ml.NumberFormat{ID: 0x2c, Code: `_($*#,##0.00_);_($*(#,##0.00);_(*"-"??_);_(@_)`}, Float},
		0x2d: {ml.NumberFormat{ID: 0x2d, Code: `mm:ss`}, DeltaTime},
		0x2e: {ml.NumberFormat{ID: 0x2e, Code: `[h]:mm:ss`}, DeltaTime},
		0x2f: {ml.NumberFormat{ID: 0x2f, Code: `mm:ss.0`}, DeltaTime},
		0x30: {ml.NumberFormat{ID: 0x30, Code: `##0.0E+0`}, Float},
		0x31: {ml.NumberFormat{ID: 0x31, Code: `@`}, General},
	}
}
