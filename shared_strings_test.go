package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal/hash"
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

	require.Equal(t, 0, len(ss.index))
	require.Equal(t, 0, ss.addString("new value"))
	require.Equal(t, 1, len(ss.index))
	require.Equal(t, 0, ss.addString("new value"))
	require.Equal(t, 0, ss.addString("new value"))
	require.Equal(t, 0, ss.addString("new value"))
	require.Equal(t, 1, len(ss.index))
	require.Equal(t, map[uint64]int{
		hash.StringItem(&ml.StringItem{Text: "new value"}).Hash(): 0,
	}, ss.index)

	require.Equal(t, 1, ss.addString("another value"))
	require.Equal(t, 2, len(ss.index))
	require.Equal(t, map[uint64]int{
		hash.StringItem(&ml.StringItem{Text: "new value"}).Hash():     0,
		hash.StringItem(&ml.StringItem{Text: "another value"}).Hash(): 1,
	}, ss.index)

	text := &ml.StringItem{
		RichText: &[]*ml.RichText{
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
	require.Equal(t, 3, len(ss.index))
	require.Equal(t, map[uint64]int{
		hash.StringItem(&ml.StringItem{Text: "new value"}).Hash():     0,
		hash.StringItem(&ml.StringItem{Text: "another value"}).Hash(): 1,
		hash.StringItem(text).Hash():                                  2,
	}, ss.index)
	require.Equal(t, 2, ss.addText(&(*text)))

	require.Equal(t, "new value", fromRichText(ss.get(0)))
	require.Equal(t, "another value", fromRichText(ss.get(1)))
	require.Equal(t, "part1part2", fromRichText(ss.get(2)))
}
