// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRowOptions(t *testing.T) {
	o := New(
		OutlineLevel(5),
		Hidden(true),
		Phonetic(true),
		Height(45.5),
		Collapsed(true),
		Styles(styles.DirectStyleID(1)),
	)

	require.IsType(t, &Info{}, o)
	require.Equal(t, &Info{
		OutlineLevel: 5,
		Hidden:       true,
		Phonetic:     true,
		Height:       45.5,
		Collapsed:    true,
		Format:       styles.DirectStyleID(1),
	}, o)

	o = New(
		OutlineLevel(10),
		Hidden(false),
		Phonetic(true),
	)
	require.Equal(t, &Info{
		Hidden:   false,
		Phonetic: true,
	}, o)
}
