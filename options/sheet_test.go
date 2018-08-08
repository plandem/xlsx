package options

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSheetOptions(t *testing.T) {
	o := NewSheetOptions(
		Sheet.Visibility(VisibilityTypeVeryHidden),
	)

	require.IsType(t, &SheetOptions{}, o)
	require.Equal(t, &SheetOptions{
		Visibility: VisibilityTypeVeryHidden,
	}, o)
}
