// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/xlsx/internal/number_format/convert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var now time.Time

func init() {
	now = time.Now()
}

func checkCellValues(sheet Sheet, rowIdx int, t *testing.T) {
	require.Equal(t, "bla-bla", sheet.Cell(1, rowIdx).Value())
	require.Equal(t, "string", sheet.Cell(2, rowIdx).Value())
	require.Equal(t, "inline string", sheet.Cell(3, rowIdx).Value())
	require.Equal(t, "1", sheet.Cell(4, rowIdx).Value())
	require.Equal(t, "12345", sheet.Cell(5, rowIdx).Value())
	require.Equal(t, "123.123", sheet.Cell(6, rowIdx).Value())
	require.Equal(t, now.Format(convert.ISO8601), sheet.Cell(7, rowIdx).Value())
}

func checkCellTypedValues(sheet Sheet, rowIdx int, t *testing.T) {
	//get values by type
	require.Equal(t, "bla-bla", sheet.Cell(1, rowIdx).String())
	require.Equal(t, "string", sheet.Cell(2, rowIdx).String())
	require.Equal(t, "inline string", sheet.Cell(3, rowIdx).String())

	//must be bool
	b, err := sheet.Cell(4, rowIdx).Bool()
	require.Nil(t, err)
	require.Equal(t, true, b)

	//bool can not be int
	_, err = sheet.Cell(4, rowIdx).Int()
	require.NotNil(t, err)

	//bool can not be float
	_, err = sheet.Cell(4, rowIdx).Float()
	require.NotNil(t, err)

	//bool can not be date
	_, err = sheet.Cell(4, rowIdx).Date()
	require.NotNil(t, err)

	//int number must be int
	i, err := sheet.Cell(5, rowIdx).Int()
	require.Nil(t, err)
	require.Equal(t, 12345, i)

	//int number can be float
	f, err := sheet.Cell(5, rowIdx).Float()
	require.Nil(t, err)
	require.Equal(t, 12345.0, f)

	//int number can not be bool
	_, err = sheet.Cell(5, rowIdx).Bool()
	require.NotNil(t, err)

	//float number must be float
	f, err = sheet.Cell(6, rowIdx).Float()
	require.Nil(t, err)
	require.Equal(t, 123.123, f)

	//float number can not be int
	_, err = sheet.Cell(6, rowIdx).Int()
	require.NotNil(t, err)

	//float number can not be bool
	_, err = sheet.Cell(6, rowIdx).Bool()
	require.NotNil(t, err)

	//date must be date
	d, err := sheet.Cell(7, rowIdx).Date()
	require.Nil(t, err)
	require.Equal(t, now.Format(convert.ISO8601), d.Format(convert.ISO8601))

	//date can not be int
	_, err = sheet.Cell(7, rowIdx).Int()
	require.NotNil(t, err)

	//date can not be bool
	_, err = sheet.Cell(7, rowIdx).Bool()
	require.NotNil(t, err)

	//date can not be float
	_, err = sheet.Cell(7, rowIdx).Float()
	require.NotNil(t, err)
}

func TestCell_write(t *testing.T) {
	xl := New()
	sheet := xl.AddSheet("test cell")

	//headers
	sheet.CellByRef("A1").SetValue("header:")
	sheet.CellByRef("B1").SetValue("General")
	sheet.CellByRef("C1").SetValue("String")
	sheet.CellByRef("D1").SetValue("Inline String")
	sheet.CellByRef("E1").SetValue("Bool")
	sheet.CellByRef("F1").SetValue("Integer")
	sheet.CellByRef("G1").SetValue("Float")
	sheet.CellByRef("H1").SetValue("Date")

	//set values by dedicated method
	sheet.CellByRef("A2").SetValue("typed set:")
	sheet.CellByRef("B2").SetValue("bla-bla")
	sheet.CellByRef("C2").SetText("string")
	sheet.CellByRef("D2").SetInlineText("inline string")
	sheet.CellByRef("E2").SetBool(true)
	sheet.CellByRef("F2").SetInt(12345)
	sheet.CellByRef("G2").SetFloat(123.123)
	sheet.CellByRef("H2").SetDateTime(now)

	//get values raw value
	checkCellValues(sheet, 1, t)

	//get typed values
	checkCellTypedValues(sheet, 1, t)

	//set values by unified method
	sheet.CellByRef("A3").SetValue("unified set:")
	sheet.CellByRef("B3").SetValue("bla-bla")
	sheet.CellByRef("C3").SetValue("string")
	sheet.CellByRef("D3").SetValue("inline string")
	sheet.CellByRef("E3").SetValue(true)
	sheet.CellByRef("F3").SetValue(12345)
	sheet.CellByRef("G3").SetValue(123.123)
	sheet.CellByRef("H3").SetValue(time.Now())

	//get values raw value
	checkCellValues(sheet, 2, t)

	//get typed values
	checkCellTypedValues(sheet, 2, t)

	xl.SaveAs("./test_files/test_cell.xlsx")
	xl.Close()
}

func TestCell_reopen(t *testing.T) {
	xl, err := Open("./test_files/test_cell.xlsx")
	require.Nil(t, err)
	sheet := xl.Sheet(0)

	//get values raw value
	checkCellValues(sheet, 1, t)
	checkCellValues(sheet, 2, t)

	//get typed values
	checkCellTypedValues(sheet, 1, t)
	checkCellTypedValues(sheet, 2, t)

	xl.Close()
}
