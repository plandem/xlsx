// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFont(t *testing.T) {
	style := New(
		Font.Name("Calibri"),
		Font.Size(10),
		Font.Bold,
		Font.Italic,
		Font.Strikeout,
		Font.Shadow,
		Font.Condense,
		Font.Extend,
		Font.Family(FontFamilyDecorative),
		Font.Color("#FF00FF"),
		Font.Underline(UnderlineTypeSingle),
		Font.VAlign(FontVAlignBaseline),
		Font.Scheme(FontSchemeMinor),
	)

	require.IsType(t, &Info{}, style)
	require.Equal(t, createStylesAndFill(func(f *Info) {
		f.styleInfo.Font = &ml.Font{
			Name:      "Calibri",
			Bold:      true,
			Italic:    true,
			Strike:    true,
			Shadow:    true,
			Condense:  true,
			Extend:    true,
			Size:      10.0,
			Color:     color.New("FFFF00FF"),
			Family:    FontFamilyDecorative,
			Underline: UnderlineTypeSingle,
			VAlign:    FontVAlignBaseline,
			Scheme:    FontSchemeMinor,
		}
	}), style)
}
