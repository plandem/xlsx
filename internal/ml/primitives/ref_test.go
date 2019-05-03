package primitives_test

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRef(t *testing.T) {
	//A10:B20 => A10:B20
	ref := primitives.RefFromCellRefs("A10", "B20")
	require.Equal(t, primitives.Ref("A10:B20"), ref)

	//A10:B20 => A10:B20
	fromCel, toCel := ref.ToCellRefs()
	require.Equal(t, primitives.CellRef("A10"), fromCel)
	require.Equal(t, primitives.CellRef("B20"), toCel)

	//B20 => B20:B20
	ref = primitives.Ref("B20")
	fromCel, toCel = ref.ToCellRefs()
	require.Equal(t, primitives.CellRef("B20"), fromCel)
	require.Equal(t, primitives.CellRef("B20"), toCel)

	//A10:A10 => A10
	ref = primitives.RefFromCellRefs("A10", "A10")
	require.Equal(t, primitives.Ref("A10"), ref)

	//0x9 => A10
	ref = primitives.RefFromIndexes(0, 9)
	require.Equal(t, primitives.Ref("A10"), ref)
}
