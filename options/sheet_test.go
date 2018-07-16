package options

import (
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSheetOptions(t *testing.T) {
	o := NewSheetOptions(
		Sheet.Visibility(types.VisibilityTypeVeryHidden),
	)

	require.IsType(t, &SheetOptions{}, o)
	require.Equal(t, &SheetOptions{
		Visibility: types.VisibilityTypeVeryHidden,
	}, o)
}
