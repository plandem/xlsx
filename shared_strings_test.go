// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSharedStrings(t *testing.T) {
	pkg := ooxml.NewPackage(nil)
	doc := &Spreadsheet{
		pkg:           pkg,
		Package:       pkg,
		relationships: ooxml.NewRelationships("not matter the name", pkg),
	}

	ss := newSharedStrings("xl/sharedStrings.xml", doc)

	require.NotNil(t, pkg)
	require.NotNil(t, ss)

	require.Equal(t, 0, ss.index.Count())
	require.Equal(t, 0, ss.addString("new value"))
	require.Equal(t, 1, ss.index.Count())
	require.Equal(t, 0, ss.addString("new value"))
	require.Equal(t, 0, ss.addString("new value"))
	require.Equal(t, 0, ss.addString("new value"))
	require.Equal(t, 1, ss.index.Count())
	require.Equal(t, 1, ss.addString("another value"))
	require.Equal(t, 2, ss.index.Count())

	text := &ml.StringItem{
		RichText: []*ml.RichText{
			{
				Text: "part1",
				Font: &ml.RichFont{
					Color: &ml.Color{RGB: "FFFF1122"},
				},
			},
			{
				Text: "part2",
				Font: &ml.RichFont{
					Size:  8,
					Color: &ml.Color{RGB: "FFFF3344"},
				},
			},
		},
	}

	require.Equal(t, 2, ss.addText(text))
	require.Equal(t, 3, ss.index.Count())
	require.Equal(t, 2, ss.addText(&(*text)))

	require.Equal(t, "new value", fromRichText(ss.get(0)))
	require.Equal(t, "another value", fromRichText(ss.get(1)))
	require.Equal(t, "part1part2", fromRichText(ss.get(2)))
}
