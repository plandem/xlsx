package types_test

import (
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCelRef(t *testing.T) {
	require.Equal(t, types.CellRef(""), types.CellRefFromIndexes(-1, -1))
	require.Equal(t, types.CellRef(""), types.CellRefFromIndexes(0, -1))
	require.Equal(t, types.CellRef(""), types.CellRefFromIndexes(-1, 0))
	require.Equal(t, types.CellRef("A1"), types.CellRefFromIndexes(0, 0))

	ref := types.CellRefFromIndexes(100, 100)
	require.Equal(t, types.CellRef("CW101"), ref)

	col, row := ref.ToIndexes()
	require.Equal(t, 100, col)
	require.Equal(t, 100, row)
}
