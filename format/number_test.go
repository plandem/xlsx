package format

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberFormat(t *testing.T) {
	style := New(
		NumberFormatID(8),
		//NumberFormat("#.### usd"),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, &StyleFormat{
		key: "b828191babe328f056db7e2fd933ee1f",
		numFormat: ml.NumberFormat{
			ID:   8,
			Code: "($#,##0.00_);[RED]($#,##0.00_)",
		},
	}, style)

	style = New(
		NumberFormat("#.### usd"),
	)

	require.Equal(t, &StyleFormat{
		key: "c002d336f2642f8c69a240eb35011015",
		numFormat: ml.NumberFormat{
			ID:   -1,
			Code: "#.### usd",
		},
	}, style)
}
