// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberFormat(t *testing.T) {
	style := New(
		NumberFormatID(8),
	)

	require.IsType(t, &Info{}, style)
	require.Equal(t, createStylesAndFill(func(f *Info) {
		f.styleInfo.NumberFormat = &ml.NumberFormat{
			ID:   8,
			Code: "($#,##0.00_);[RED]($#,##0.00_)",
		}
	}), style)

	style = New(
		NumberFormat(`$0.00" usd"`),
	)

	require.Equal(t, createStylesAndFill(func(f *Info) {
		f.styleInfo.NumberFormat = &ml.NumberFormat{
			ID:   -1,
			Code: `$0.00" usd"`,
		}
	}), style)
}
