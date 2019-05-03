package xlsx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpreadsheet_Sheet(t *testing.T) {
	xl, err := Open("./test_files/example_simple.xlsx")
	assert.NotNil(t, xl)
	assert.Nil(t, err)
	assert.IsType(t, &Spreadsheet{}, xl)

	defer xl.Close()

	//stream mode should be new each time
	sheet := xl.Sheet(0, SheetModeStream | SheetModeMultiPhase)
	assert.Equal(t, sheetModeRead | SheetModeStream | SheetModeMultiPhase, sheet.mode())

	sheet = xl.Sheet(0, SheetModeStream | SheetModeIgnoreDimension)
	assert.Equal(t, sheetModeRead | SheetModeStream | SheetModeIgnoreDimension, sheet.mode())

	//normal mode should not be changed
	sheet = xl.Sheet(0, SheetModeIgnoreDimension)
	assert.Equal(t, sheetModeRead | sheetModeWrite | SheetModeIgnoreDimension, sheet.mode())

	sheet = xl.Sheet(0)
	assert.Equal(t, sheetModeRead | sheetModeWrite | SheetModeIgnoreDimension, sheet.mode())

	//stream mode should not work after normal mode
	assert.Panics(t, func() {
		sheet = xl.Sheet(0, SheetModeStream)
	})

	sheet = xl.AddSheet("a new sheet")
	assert.Equal(t, sheetModeRead | sheetModeWrite, sheet.mode())

	sheet = xl.Sheet(1)
	assert.Equal(t, sheetModeRead | sheetModeWrite, sheet.mode())

	sheet = xl.Sheet(1, SheetModeIgnoreDimension)
	assert.Equal(t, sheetModeRead | sheetModeWrite, sheet.mode())

}
