package color_test

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColor(t *testing.T) {
	require.Equal(t, &ml.Color{Indexed: primitives.OptionalIndex(6)}, color.New("#FF00FF"))
	require.Equal(t, &ml.Color{RGB: "FF112233"}, color.New("#112233"))
}
