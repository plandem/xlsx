// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/types/options/column"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCol(t *testing.T) {
	xl, err := Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()
	sheet := xl.Sheet(0)
	c := sheet.Col(5)

	o := options.New(
		options.Width(0),
		options.OutlineLevel(10),
		options.Hidden(true),
		options.Phonetic(true),
		options.Collapsed(true),
	)

	c.SetOptions(o)

	require.Equal(t, c.ml.OutlineLevel, o.OutlineLevel)
	require.Equal(t, c.ml.Hidden, o.Hidden)
	require.Equal(t, c.ml.Phonetic, o.Phonetic)
	require.Equal(t, c.ml.Collapsed, o.Collapsed)
	require.Equal(t, c.ml.Width, float32(0.0))
	require.Equal(t, c.ml.CustomWidth, false)
	require.Equal(t, c.ml.Style, styles.DirectStyleID(0))

	o = options.New(
		options.Width(100.0),
	)

	c.SetOptions(o)
	require.Equal(t, c.ml.OutlineLevel, o.OutlineLevel)
	require.Equal(t, c.ml.Hidden, o.Hidden)
	require.Equal(t, c.ml.Phonetic, o.Phonetic)
	require.Equal(t, c.ml.Collapsed, o.Collapsed)
	require.Equal(t, c.ml.Width, float32(100.0))
	require.Equal(t, c.ml.CustomWidth, true)
	require.Equal(t, c.ml.Style, styles.DirectStyleID(0))

	style := styles.New(
		styles.Font.Name("Calibri"),
		styles.Font.Size(12),
	)

	styleRef := xl.AddStyles(style)
	c.SetStyles(styleRef)

	require.Equal(t, c.ml.Style, styles.DirectStyleID(styleRef))
}
