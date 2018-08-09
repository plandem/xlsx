package xlsx

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/number_format"
	"github.com/plandem/xlsx/internal/number_format/convert"
	"github.com/plandem/xlsx/types"
	"math"
	"strconv"
	"time"
)

//max length that excel cell can hold
const cellStringValueLimit = 32767

//Cell is a higher level object that wraps ml.Cell with functionality
type Cell struct {
	ml    *ml.Cell
	sheet *sheetInfo
}

var (
	errTypeMismatch = errors.New("type mismatch")
)

//Type returns current type of cell
func (c *Cell) Type() types.CellType {
	return c.ml.Type
}

//Value returns current raw value of cell
func (c *Cell) Value() string {
	var value string

	switch c.ml.Type {
	case types.CellTypeInlineString:
		if c.ml.InlineStr != nil {
			value = string(c.ml.InlineStr.Text)
		}
	case types.CellTypeSharedString:
		var sid int

		if len(c.ml.Value) > 0 {
			sid, _ = strconv.Atoi(c.ml.Value)
		}

		value = c.sheet.workbook.doc.sharedStrings.get(sid)
	default:
		value = c.ml.Value
	}

	return value
}

//String returns formatted value as string respecting cell number format and type. Any errors ignored to conform String() interface.
func (c *Cell) String() string {
	//if cell has error, then just return value that Excel put here
	if c.ml.Type == types.CellTypeError {
		return c.ml.Value
	}

	code := c.sheet.workbook.doc.styleSheet.resolveNumberFormat(c.ml.Style)

	//N.B.: Maybe it's not a good idea to use resolved value (e.g. inline string) for conversion?!
	return numberFormat.Format(c.Value(), code, c.ml.Type)
}

//Date try to convert and return current raw value as time.Time
func (c *Cell) Date() (time.Time, error) {
	if c.ml.Type == types.CellTypeDate || c.ml.Type == types.CellTypeNumber || c.ml.Type == types.CellTypeGeneral {
		return convert.ToDate(c.ml.Value)
	}

	return time.Now(), errTypeMismatch
}

//Int try to convert and return current raw value as int
func (c *Cell) Int() (int, error) {
	if c.ml.Type == types.CellTypeNumber || c.ml.Type == types.CellTypeGeneral {
		return convert.ToInt(c.ml.Value)
	}

	return 0, errTypeMismatch
}

//Float try to convert and return current raw value as float64
func (c *Cell) Float() (float64, error) {
	if c.ml.Type == types.CellTypeNumber || c.ml.Type == types.CellTypeGeneral {
		return convert.ToFloat(c.ml.Value)
	}

	return math.NaN(), errTypeMismatch
}

//Bool try to convert and return current raw value as bool
func (c *Cell) Bool() (bool, error) {
	if c.ml.Type == types.CellTypeBool || c.ml.Type == types.CellTypeGeneral || c.ml.Type == types.CellTypeNumber {
		return convert.ToBool(c.ml.Value)
	}

	return false, errTypeMismatch
}

//setGeneral sets the value as general type
func (c *Cell) setGeneral(value string) {
	c.ml.Type = types.CellTypeGeneral
	c.ml.Value = value
	c.ml.Formula = nil
	c.ml.InlineStr = nil
}

//truncateIfRequired truncate string is exceeded allowed size
func (c *Cell) truncateIfRequired(value string) string {
	if len(value) > cellStringValueLimit {
		value = value[:cellStringValueLimit]
	}

	return value
}

//SetInlineString sets value as inline string
func (c *Cell) SetInlineString(value string) {
	if len(value) == 0 {
		c.setGeneral(value)
		return
	}

	c.ml.Type = types.CellTypeInlineString
	c.ml.Value = ""
	c.ml.Formula = nil
	c.ml.InlineStr = &ml.StringItem{Text: types.Text(c.truncateIfRequired(value))}
}

//SetString sets value as shared string
func (c *Cell) SetString(value string) {
	if len(value) == 0 {
		c.setGeneral(value)
		return
	}

	//we can update sharedStrings only when sheet is in write mode, to prevent pollution of sharedStrings with fake values
	if (c.sheet.mode() & sheetModeWrite) == 0 {
		panic(errorNotSupportedWrite)
	}

	//sharedStrings is the only place that can be mutated from the 'sheet' perspective
	sid := c.sheet.workbook.doc.sharedStrings.add(c.truncateIfRequired(value))
	c.ml.Formula = nil
	c.ml.Type = types.CellTypeSharedString
	c.ml.Value = strconv.Itoa(sid)
}

//SetInt sets an integer value
func (c *Cell) SetInt(value int) {
	c.ml.Type = types.CellTypeNumber
	c.ml.Value = strconv.Itoa(value)
	c.ml.Style = c.sheet.workbook.doc.styleSheet.typedStyles[numberFormat.Integer]
	c.ml.Formula = nil
	c.ml.InlineStr = nil
}

//SetFloat sets a float value
func (c *Cell) SetFloat(value float64) {
	c.ml.Type = types.CellTypeNumber
	c.ml.Value = strconv.FormatFloat(value, 'f', -1, 64)
	c.ml.Style = c.sheet.workbook.doc.styleSheet.typedStyles[numberFormat.Float]
	c.ml.Formula = nil
	c.ml.InlineStr = nil
}

//SetBool sets a bool value
func (c *Cell) SetBool(value bool) {
	c.ml.Type = types.CellTypeBool
	c.ml.Formula = nil
	c.ml.InlineStr = nil

	if value {
		c.ml.Value = "1"
	} else {
		c.ml.Value = "0"
	}
}

//setDate is a general setter for date types
func (c *Cell) setDate(value time.Time, t numberFormat.Type) {
	c.ml.Type = types.CellTypeDate
	c.ml.Value = value.Format(convert.ISO8601)
	c.ml.Style = c.sheet.workbook.doc.styleSheet.typedStyles[t]
	c.ml.Formula = nil
	c.ml.InlineStr = nil
}

//SetDateTime sets a time value with number format for datetime
func (c *Cell) SetDateTime(value time.Time) {
	c.setDate(value, numberFormat.DateTime)
}

//SetDate sets a time value with number format for date
func (c *Cell) SetDate(value time.Time) {
	c.setDate(value, numberFormat.Date)
}

//SetTime sets a time value with number format for time
func (c *Cell) SetTime(value time.Time) {
	c.setDate(value, numberFormat.Time)
}

//SetDeltaTime sets a time value with number format for delta time
func (c *Cell) SetDeltaTime(value time.Time) {
	c.setDate(value, numberFormat.DeltaTime)
}

//SetValue sets a value
func (c *Cell) SetValue(value interface{}) {
	switch t := value.(type) {
	case int:
		c.SetInt(value.(int))
	case int8:
		c.SetInt(int(value.(int8)))
	case int16:
		c.SetInt(int(value.(int16)))
	case int32:
		c.SetInt(int(value.(int32)))
	case int64:
		c.SetInt(int(value.(int64)))
	case float32:
		c.SetFloat(float64(value.(float32)))
	case float64:
		c.SetFloat(value.(float64))
	case string:
		c.SetString(t)
	case []byte:
		c.SetString(string(t))
	case bool:
		c.SetBool(bool(t))
	case time.Time:
		c.setDate(time.Time(t), numberFormat.DateTime)
	case nil:
		c.Reset()
	default:
		c.SetString(fmt.Sprintf("%v", value))
	}
}

//SetValueWithFormat is helper function that internally works as SetValue and SetFormatting
func (c *Cell) SetValueWithFormat(value interface{}, formatCode string) {
	//we can update styleSheet only when sheet is in write mode, to prevent pollution of styleSheet with fake values
	if (c.sheet.mode() & sheetModeWrite) == 0 {
		panic(errorNotSupportedWrite)
	}

	styleID := c.sheet.workbook.doc.styleSheet.addStyle(format.New(format.NumberFormat(formatCode)))

	c.SetValue(value)
	c.ml.Style = styleID
}

//Reset resets current current cell information
func (c *Cell) Reset() {
	*c.ml = ml.Cell{Ref: c.ml.Ref}
}

//Clear clears cell's value
func (c *Cell) Clear() {
	c.ml.Value = ""
}

//HasFormula returns true if cell has formula
func (c *Cell) HasFormula() bool {
	return c.ml.Formula != nil && (*c.ml.Formula != ml.CellFormula{})
}

//HasFormatting returns true if cell has styles
func (c *Cell) HasFormatting() bool {
	//0 is default style
	return c.ml.Style != 0
}

//SetFormatting sets style format to requested styleID
func (c *Cell) SetFormatting(styleID format.StyleID) {
	c.ml.Style = styleID
}
