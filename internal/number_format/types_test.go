package numberFormat

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberFormat(t *testing.T) {
	require.Equal(t, true, IsBuiltIn(0))
	require.Equal(t, true, IsBuiltIn(163))
	require.Equal(t, false, IsBuiltIn(-1))
	require.Equal(t, false, IsBuiltIn(164))

	//built-in ID was provided, ignore built-in CODE
	require.Equal(t, builtIn[0x00], Resolve(ml.NumberFormat{ID: 0, Code: "0.00"}))

	//built-in ID was provided, ignore custom CODE and return general code/type
	unknownBuiltIn := &BuiltInFormat{ml.NumberFormat{ID: 162, Code: "@"}, General}
	require.Equal(t, unknownBuiltIn, Resolve(ml.NumberFormat(ml.NumberFormat{ID: 162, Code: "abcd"})))

	//built-in CODE was provided, ignore custom ID
	require.Equal(t, builtIn[0x02], Resolve(ml.NumberFormat{ID: 1000, Code: "0.00"}))

	//custom ID and CODE were provided
	require.Nil(t, Resolve(ml.NumberFormat{ID: 1000, Code: "abcde"}))

	//built-in ID was provided, ignore custom CODE
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: 0, Code: "@"}), New(0, "abcd"))

	//built-in ID was provided, ignore custom CODE
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: 0, Code: "@"}), New(0, "abcd"))

	//built-in ID was provided, ignore built-in CODE
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: 0, Code: "@"}), New(0, "0.00"))

	//built-in CODE was provided, ignore custom ID
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: 2, Code: "0.00"}), New(1000, "0.00"))

	//custom CODE was provided, ignore custom ID
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: -1, Code: "abcd"}), New(1000, "abcd"))

	//custom ID was provided
	require.Equal(t, ml.NumberFormat(ml.NumberFormat{ID: 1000, Code: ""}), New(1000, ""))
}
