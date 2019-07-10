// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives_test

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCelRef(t *testing.T) {
	require.Equal(t, primitives.CellRef(""), primitives.CellRefFromIndexes(-1, -1))
	require.Equal(t, primitives.CellRef(""), primitives.CellRefFromIndexes(0, -1))
	require.Equal(t, primitives.CellRef(""), primitives.CellRefFromIndexes(-1, 0))
	require.Equal(t, primitives.CellRef("A1"), primitives.CellRefFromIndexes(0, 0))

	ref := primitives.CellRefFromIndexes(100, 100)
	require.Equal(t, primitives.CellRef("CW101"), ref)

	col, row := ref.ToIndexes()
	require.Equal(t, 100, col)
	require.Equal(t, 100, row)
}
