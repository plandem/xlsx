// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives_test

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRefList(t *testing.T) {
	//list of Ref into RefList
	refList := primitives.RefListFromRefs("A10", "A11:B20")
	require.Equal(t, primitives.RefList("A10 A11:B20"), refList)

	//RefList info list of Ref
	require.Equal(t, []primitives.Ref{"A10", "A11:B20"}, refList.ToRefs())

	//RefList into BoundsList
	require.Equal(t, primitives.BoundsList{
		primitives.BoundsFromIndexes(0, 9, 0, 9),
		primitives.BoundsFromIndexes(0, 10, 1, 19),
	}, refList.ToBoundsList())
}
