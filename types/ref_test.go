package types_test

import (
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRef(t *testing.T) {
	ref := types.RefFromCellRefs("A10", "B20")
	require.Equal(t, types.Ref("A10:B20"), ref)

	fromCel, toCel := ref.ToCellRefs()
	require.Equal(t, types.CellRef("A10"), fromCel)
	require.Equal(t, types.CellRef("B20"), toCel)


	ref = types.Ref("B20")
	require.Equal(t, types.Ref("B20"), ref)

	fromCel, toCel = ref.ToCellRefs()
	require.Equal(t, types.CellRef("A1"), fromCel)
	require.Equal(t, types.CellRef("B20"), toCel)

}
