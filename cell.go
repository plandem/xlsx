package xlsx

import (
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"strconv"
)

//max length that excel cell can hold
const cellStringValueLimit = 32767

//Cell is a higher level object that wraps ml.Cell with functionality
type Cell struct {
	ml    *ml.Cell
	sheet *Sheet
}

//Type returns current type of cell
func (c *Cell) Type() types.CellType {
	return c.ml.Type
}

//GetValue returns current value of cell
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
	c.ml.Formula = nil

	if len(value) == 0 {
		c.ml.Value = value
		c.ml.Type = types.CellTypeGeneral
		c.ml.InlineStr = nil
		return
	}

	c.ml.Type = types.CellTypeInlineString
	c.ml.Value = ""
	c.ml.Formula = nil
	c.ml.InlineStr = &ml.StringItem{Text: types.Text(c.truncateIfRequired(value))}
}

//SetString sets value as shared string
func (c *Cell) SetString(value string) {
	c.ml.Formula = nil

	if len(value) == 0 {
		c.ml.Value = value
		c.ml.Type = types.CellTypeGeneral
		return
	}

	sid := c.sheet.workbook.doc.sharedStrings.add(c.truncateIfRequired(value))
	c.ml.Type = types.CellTypeSharedString
	c.ml.Value = strconv.Itoa(sid)
}

//SetInt sets an integer value
func (c *Cell) SetInt(value int) {
	c.setGeneral(strconv.Itoa(value))
}

//SetFloat sets a float value
func (c *Cell) SetFloat(value float64) {
	c.setGeneral(strconv.FormatFloat(value, 'f', -1, 64))
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
	case nil:
		c.Reset()
	default:
		c.SetString(fmt.Sprintf("%v", value))
	}
}

//Reset resets current current cell information
func (c *Cell) Reset() {
	c.ml = &ml.Cell{}
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

//SetFormatting sets style format to requested styleRef
func (c *Cell) SetFormatting(styleRef format.StyleRefID) {
	c.ml.Style = styleRef
}
