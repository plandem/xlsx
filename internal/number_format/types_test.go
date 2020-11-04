// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package number

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberFormat(t *testing.T) {
	require.Equal(t, true, IsBuiltIn(0))
	require.Equal(t, true, IsBuiltIn(163))
	require.Equal(t, false, IsBuiltIn(-1))
	require.Equal(t, false, IsBuiltIn(164))

	//built-in ID was provided, ignore built-in CODE
	require.Equal(t, builtIn[0x00], Resolve(ml.NumberFormat{ID: 0, Code: "0.00"}))

	//built-in ID was provided, ignore custom CODE and return general code/type
	unknownBuiltIn := &builtInFormat{ml.NumberFormat{ID: 162, Code: "General"}, General}
	require.Equal(t, unknownBuiltIn, Resolve(ml.NumberFormat(ml.NumberFormat{ID: 162, Code: "abcd"})))

	//built-in CODE was provided, ignore custom ID
	require.Equal(t, builtIn[0x02], Resolve(ml.NumberFormat{ID: 1000, Code: "0.00"}))

	//custom ID and CODE were provided
	require.Nil(t, Resolve(ml.NumberFormat{ID: 1000, Code: "abcde"}))

	//built-in ID was provided, ignore custom CODE
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: 0, Code: "General"}), New(0, "abcd"))

	//built-in ID was provided, ignore built-in CODE
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: 0, Code: "General"}), New(0, "0.00"))

	//built-in CODE was provided, ignore custom ID
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: 2, Code: "0.00"}), New(1000, "0.00"))

	//custom CODE was provided, ignore custom ID
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: -1, Code: "abcd"}), New(1000, "abcd"))

	//custom ID was provided
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: 1000, Code: ""}), New(1000, ""))
}

func TestDateFormat(t *testing.T) {
	value := "36892.521"
	var testCases = []struct {
		code     string
		expected string
	}{
		{"d/m/yyyy", "01/01/2001"},
		{"d-mmm-yy", "01-Jan-01"},
		{"d-mmm", "01-Jan"},
		{"mmm-yy", "Jan-01"},
		{"h:mm AM/PM", "12:30 PM"},
		{"h:mm:ss AM/PM", "12:30:14 PM"},
		{"h:mm", "12:30"},
		{"h:mm:ss", "12:30:14"},
		{"m/d/yy H:mm", "01/01/01 12:30"},
	}
	for _, test := range testCases {
		t.Run(test.code, func(t *testing.T) {
			result := Format(value, test.code, primitives.CellType(0))
			assert.Equal(t, test.expected, result)
		})
	}
}
