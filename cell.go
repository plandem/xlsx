// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/number_format"
	"github.com/plandem/xlsx/internal/number_format/convert"
	"github.com/plandem/xlsx/types"
	"github.com/plandem/xlsx/types/hyperlink"
	"math"
	"strconv"
	"time"
)

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
			value = fromRichText(c.ml.InlineStr)
		}
	case types.CellTypeSharedString:
		var sid int

		if len(c.ml.Value) > 0 {
			sid, _ = strconv.Atoi(c.ml.Value)
		}

		value = fromRichText(c.sheet.workbook.doc.sharedStrings.get(sid))
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
	return number.Format(c.Value(), code, c.ml.Type)
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
	if len(value) > internal.ExcelCellLimit {
		value = value[:internal.ExcelCellLimit]
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
	sid := c.sheet.workbook.doc.sharedStrings.addString(c.truncateIfRequired(value))
	c.ml.Formula = nil
	c.ml.Type = types.CellTypeSharedString
	c.ml.Value = strconv.Itoa(sid)
}

//SetText sets shared rich text
func (c *Cell) SetText(parts ...interface{}) error {
	//we can update sharedStrings only when sheet is in write mode, to prevent pollution of sharedStrings with fake values
	if (c.sheet.mode() & sheetModeWrite) == 0 {
		panic(errorNotSupportedWrite)
	}

	//sharedStrings is the only place that can be mutated from the 'sheet' perspective
	text, err := toRichText(parts...)
	if err == nil {
		sid := c.sheet.workbook.doc.sharedStrings.addText(text)
		c.ml.Formula = nil
		c.ml.Type = types.CellTypeSharedString
		c.ml.Value = strconv.Itoa(sid)
	}

	return err
}

//SetInlineText sets inline rich text
func (c *Cell) SetInlineText(parts ...interface{}) error {
	text, err := toRichText(parts...)

	if err == nil {
		c.ml.Type = types.CellTypeInlineString
		c.ml.Value = ""
		c.ml.Formula = nil
		c.ml.InlineStr = text
	}

	return err
}

//SetInt sets an integer value
func (c *Cell) SetInt(value int) {
	c.ml.Type = types.CellTypeNumber
	c.ml.Value = strconv.FormatInt(int64(value), 10)

	if c.ml.Style == styles.DirectStyleID(0) {
		c.ml.Style = c.sheet.workbook.doc.styleSheet.typedStyles[number.Integer]
	}

	c.ml.Formula = nil
	c.ml.InlineStr = nil
}

//SetUInt sets an unsigned integer value
func (c *Cell) SetUInt(value uint) {
	c.ml.Type = types.CellTypeNumber
	c.ml.Value = strconv.FormatUint(uint64(value), 10)

	if c.ml.Style == styles.DirectStyleID(0) {
		c.ml.Style = c.sheet.workbook.doc.styleSheet.typedStyles[number.Integer]
	}

	c.ml.Formula = nil
	c.ml.InlineStr = nil
}

//SetFloat sets a float value
func (c *Cell) SetFloat(value float64) {
	c.ml.Type = types.CellTypeNumber
	c.ml.Value = strconv.FormatFloat(value, 'f', -1, 64)

	if c.ml.Style == styles.DirectStyleID(0) {
		c.ml.Style = c.sheet.workbook.doc.styleSheet.typedStyles[number.Float]
	}

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
func (c *Cell) setDate(value time.Time, t number.Type) {
	c.ml.Type = types.CellTypeDate
	c.ml.Value = value.Format(convert.ISO8601)

	if c.ml.Style == styles.DirectStyleID(0) {
		c.ml.Style = c.sheet.workbook.doc.styleSheet.typedStyles[t]
	}

	c.ml.Formula = nil
	c.ml.InlineStr = nil
}

//SetDateTime sets a time value with number format for datetime
func (c *Cell) SetDateTime(value time.Time) {
	c.setDate(value, number.DateTime)
}

//SetDate sets a time value with number format for date
func (c *Cell) SetDate(value time.Time) {
	c.setDate(value, number.Date)
}

//SetTime sets a time value with number format for time
func (c *Cell) SetTime(value time.Time) {
	c.setDate(value, number.Time)
}

//SetDeltaTime sets a time value with number format for delta time
func (c *Cell) SetDeltaTime(value time.Time) {
	c.setDate(value, number.DeltaTime)
}

//SetValue sets a value
func (c *Cell) SetValue(value interface{}) {
	switch v := value.(type) {
	case int:
		c.SetInt(v)
	case int8:
		c.SetInt(int(v))
	case int16:
		c.SetInt(int(v))
	case int32:
		c.SetInt(int(v))
	case int64:
		c.SetInt(int(v))
	case uint:
		c.SetUInt(v)
	case uint8:
		c.SetUInt(uint(v))
	case uint16:
		c.SetUInt(uint(v))
	case uint32:
		c.SetUInt(uint(v))
	case uint64:
		c.SetUInt(uint(v))
	case float32:
		c.SetFloat(float64(v))
	case float64:
		c.SetFloat(v)
	case string:
		c.SetString(v)
	case []byte:
		c.SetString(string(v))
	case bool:
		c.SetBool(v)
	case time.Time:
		c.setDate(v, number.DateTime)
	case []interface{}:
		_ = c.SetText(v...)
	case nil:
		c.Reset()
	default:
		c.SetString(fmt.Sprintf("%v", value))
	}
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

//Styles returns DirectStyleID of active format for cell
func (c *Cell) Styles() styles.DirectStyleID {
	return c.ml.Style
}

//SetStyles sets style format to requested DirectStyleID
func (c *Cell) SetStyles(styleID styles.DirectStyleID) {
	c.ml.Style = styleID
}

//SetValueWithStyles is helper function that internally works as SetValue and SetStyles with NumberFormat
func (c *Cell) SetValueWithStyles(value interface{}, formatCode string) {
	//we can update styleSheet only when sheet is in write mode, to prevent pollution of styleSheet with fake values
	if (c.sheet.mode() & sheetModeWrite) == 0 {
		panic(errorNotSupportedWrite)
	}

	styleID := c.sheet.workbook.doc.styleSheet.addStyle(styles.New(styles.NumberFormat(formatCode)))

	c.SetValue(value)
	c.ml.Style = ml.DirectStyleID(styleID)
}

//Hyperlink returns resolved hyperlink.Info if there is any hyperlink or nil otherwise
func (c *Cell) Hyperlink() *hyperlink.Info {
	return c.sheet.hyperlinks.Get(c.ml.Ref)
}

//SetHyperlink sets hyperlink for cell, where link can be string or hyperlink.Info
func (c *Cell) SetHyperlink(link interface{}) error {
	styleID, err := c.sheet.hyperlinks.Add(types.RefFromIndexes(c.ml.Ref.ToIndexes()).ToBounds(), link)
	if err != nil {
		return err
	}

	c.SetStyles(styleID)
	return nil
}

//SetValueWithHyperlink is helper function that internally works as SetValue and SetHyperlink
func (c *Cell) SetValueWithHyperlink(value interface{}, link interface{}) error {
	err := c.SetHyperlink(link)

	if err == nil {
		c.SetValue(value)
	}

	return err
}

//RemoveHyperlink removes hyperlink from cell
func (c *Cell) RemoveHyperlink() {
	c.sheet.hyperlinks.Remove(types.RefFromIndexes(c.ml.Ref.ToIndexes()).ToBounds())
}

//SetComment sets comment for cell, where comment can be string or comment.Info
func (c *Cell) SetComment(comment interface{}) error {
	return c.sheet.comments.Add(types.RefFromIndexes(c.ml.Ref.ToIndexes()).ToBounds(), comment)
}

//RemoveComment removes comment from cell
func (c *Cell) RemoveComment() {
	c.sheet.comments.Remove(types.RefFromIndexes(c.ml.Ref.ToIndexes()).ToBounds())
}
