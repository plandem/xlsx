package format

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

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, &StyleFormat{
		key: "0b65d5dc27be42d7f4826817e6650162",
		protection: ml.CellProtection{
			Locked: true,
			Hidden: true,
		},
	}, style)

}
