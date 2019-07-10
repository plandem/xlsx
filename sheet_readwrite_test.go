// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx_test

import (
	"github.com/plandem/xlsx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSheetReadWrite_InsertCol(t *testing.T) {
	xl := xlsx.New()
	defer xl.Close()

	s := xl.AddSheet("The first sheet")

	require.Equal(t, []string{""}, s.Row(0).Values())

	//insert col at 0
	col := s.InsertCol(0)
	col.Cell(0).SetValue("0")
	require.Equal(t, []string{"0", ""}, s.Row(0).Values())

	col = s.InsertCol(0)
	col.Cell(0).SetValue("1")
	require.Equal(t, []string{"1", "0", ""}, s.Row(0).Values())

	col = s.InsertCol(0)
	col.Cell(0).SetValue("2")
	require.Equal(t, []string{"2", "1", "0", ""}, s.Row(0).Values())

	//insert col between 0 and 1
	col = s.InsertCol(1)
	col.Cell(0).SetValue("between")
	require.Equal(t, []string{"2", "between", "1", "0", ""}, s.Row(0).Values())

	//insert col after last
	col = s.InsertCol(5)
	col.Cell(0).SetValue("last")
	require.Equal(t, []string{"2", "between", "1", "0", "", "last"}, s.Row(0).Values())
}

func TestSheetReadWrite_InsertRow(t *testing.T) {
	xl := xlsx.New()
	defer xl.Close()

	s := xl.AddSheet("The first sheet")

	require.Equal(t, []string{""}, s.Row(0).Values())

	//insert row at 0
	row := s.InsertRow(0)
	row.Cell(0).SetValue("0")
	require.Equal(t, []string{"0", ""}, s.Col(0).Values())

	row = s.InsertRow(0)
	row.Cell(0).SetValue("1")
	require.Equal(t, []string{"1", "0", ""}, s.Col(0).Values())

	row = s.InsertRow(0)
	row.Cell(0).SetValue("2")
	require.Equal(t, []string{"2", "1", "0", ""}, s.Col(0).Values())

	//insert row between 0 and 1
	row = s.InsertRow(1)
	row.Cell(0).SetValue("between")
	require.Equal(t, []string{"2", "between", "1", "0", ""}, s.Col(0).Values())

	//insert after last row
	row = s.InsertRow(5)
	row.Cell(0).SetValue("last")
	require.Equal(t, []string{"2", "between", "1", "0", "", "last"}, s.Col(0).Values())
}
