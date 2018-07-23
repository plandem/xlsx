package format

import (
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColor(t *testing.T) {
	indexedColor := 6
	require.Equal(t, &ml.Color{Indexed: sharedML.OptionalIndex(&indexedColor)}, newColor("#FF00FF"))
	require.Equal(t, &ml.Color{RGB: "FF112233"}, newColor("#112233"))
}
