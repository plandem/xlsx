package xlsx

import (
	"github.com/plandem/ooxml"
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
	require.Equal(t, 0, ss.add("new value"))
	require.Equal(t, 1, len(ss.index))
	require.Equal(t, 0, ss.add("new value"))
	require.Equal(t, 0, ss.add("new value"))
	require.Equal(t, 0, ss.add("new value"))
	require.Equal(t, 1, len(ss.index))
	require.Equal(t, map[string]int{"new value": 0}, ss.index)

	require.Equal(t, 1, ss.add("another value"))
	require.Equal(t, 2, len(ss.index))
	require.Equal(t, map[string]int{"new value": 0, "another value": 1}, ss.index)

	require.Equal(t, "new value", ss.get(0))
	require.Equal(t, "another value", ss.get(1))
}
