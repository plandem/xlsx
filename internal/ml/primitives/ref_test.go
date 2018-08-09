package primitives_test

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRef(t *testing.T) {
	ref := primitives.RefFromCellRefs("A10", "B20")
	require.Equal(t, primitives.Ref("A10:B20"), ref)

	fromCel, toCel := ref.ToCellRefs()
	require.Equal(t, primitives.CellRef("A10"), fromCel)
	require.Equal(t, primitives.CellRef("B20"), toCel)

	ref = primitives.Ref("B20")
	require.Equal(t, primitives.Ref("B20"), ref)

	fromCel, toCel = ref.ToCellRefs()
	require.Equal(t, primitives.CellRef("A1"), fromCel)
	require.Equal(t, primitives.CellRef("B20"), toCel)

}
