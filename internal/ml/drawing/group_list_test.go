// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing_test

import (
	"github.com/plandem/xlsx/internal/ml/drawing"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroupList_Add(t *testing.T) {
	shapes := drawing.GroupList{}
	shapes.Add(1)
	require.Equal(t, drawing.GroupList{1}, shapes)
	shapes.Add(2)
	require.Equal(t, drawing.GroupList{1, 2}, shapes)
	shapes.Add(3)
	require.Equal(t, drawing.GroupList{1, 2, 3}, shapes)
	shapes.Add(1)
	require.Equal(t, drawing.GroupList{1, 2, 3, 1}, shapes)
}
