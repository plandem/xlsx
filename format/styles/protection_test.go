// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProtection(t *testing.T) {
	style := New(
		Protection.Hidden,
		Protection.Locked,
	)

	require.IsType(t, &Info{}, style)
	require.Equal(t, createStylesAndFill(func(f *Info) {
		f.styleInfo.Protection = &ml.CellProtection{
			Locked: true,
			Hidden: true,
		}
	}), style)
}
