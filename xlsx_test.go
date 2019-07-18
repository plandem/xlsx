// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx_test

import (
	"bytes"
	"github.com/plandem/xlsx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	xl := xlsx.New()
	assert.NotNil(t, xl)
	assert.IsType(t, &xlsx.Spreadsheet{}, xl)
}

func TestOpening(t *testing.T) {
	//can't open
	xl, err := xlsx.Open("./test_files/unknown_file.xlsx")
	assert.Nil(t, xl)
	assert.NotNil(t, err)

	//non zip
	xl, err = xlsx.Open("./xlsx.go")
	assert.Nil(t, xl)
	assert.NotNil(t, err)

	//ok
	xl, err = xlsx.Open("./test_files/example_simple.xlsx")
	assert.NotNil(t, xl)
	assert.Nil(t, err)
	assert.IsType(t, &xlsx.Spreadsheet{}, xl)

	//non zip
	nonZipFile, err := os.Open("./xlsx.go")
	assert.Nil(t, err)
	xl, err = xlsx.Open(nonZipFile)
	assert.NotNil(t, err)
	assert.Nil(t, xl)

	//zip file
	zipFile, err := os.Open("./test_files/example_simple.xlsx")
	assert.NotNil(t, zipFile)
	assert.Nil(t, err)

	//ok
	xl, err = xlsx.Open(zipFile)
	assert.NotNil(t, xl)
	assert.Nil(t, err)
	assert.IsType(t, &xlsx.Spreadsheet{}, xl)
}

func TestSaving(t *testing.T) {
	xl := xlsx.New()
	assert.NotNil(t, xl)
	assert.IsType(t, &xlsx.Spreadsheet{}, xl)

	//no filename
	err := xl.Save()
	assert.NotNil(t, err)

	//no sheets
	err = xl.SaveAs("./test_files/tmp.xlsx")
	assert.NotNil(t, err)

	//ok
	xl.AddSheet("new sheet")
	err = xl.SaveAs("./test_files/tmp.xlsx")
	assert.Nil(t, err)

	//save to io.Writer
	xl, err = xlsx.Open("./test_files/example_simple.xlsx")
	require.NotNil(t, xl)
	require.Nil(t, err)
	require.IsType(t, &xlsx.Spreadsheet{}, xl)
	err = xl.SaveAs(bytes.NewBuffer(nil))
	require.Nil(t, err)

	//save under differ name
	xl, err = xlsx.Open("./test_files/example_simple.xlsx")
	require.NotNil(t, xl)
	require.Nil(t, err)
	require.IsType(t, &xlsx.Spreadsheet{}, xl)
	err = xl.SaveAs("./test_files/tmp.xlsx")
	require.Nil(t, err)

	//open saved
	xl, err = xlsx.Open("./test_files/tmp.xlsx")
	require.NotNil(t, xl)
	require.Nil(t, err)
	require.IsType(t, &xlsx.Spreadsheet{}, xl)
	assert.Equal(t, []string{"Sheet1"}, xl.SheetNames())

	//save with same name
	xl.AddSheet("new sheet")
	err = xl.Save()
	assert.Nil(t, err)

	//open saved and check for a new sheet
	xl, err = xlsx.Open("./test_files/tmp.xlsx")
	require.NotNil(t, xl)
	require.Nil(t, err)
	require.IsType(t, &xlsx.Spreadsheet{}, xl)
	require.Equal(t, []string{"Sheet1", "new sheet"}, xl.SheetNames())
}
