// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSheetOptions(t *testing.T) {
	o := New(
		Visibility(VisibilityVeryHidden),
	)

	require.IsType(t, &Info{}, o)
	require.Equal(t, &Info{
		Visibility: VisibilityVeryHidden,
	}, o)
}
