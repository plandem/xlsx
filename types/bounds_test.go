package types_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBounds(t *testing.T) {
	type Entity struct {
		Attribute types.Bounds `xml:"attribute,attr"`
	}

	entity := Entity{Attribute: types.BoundsFromIndexes(0, 0, 10, 10)}
	encoded, err := xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity attribute="A1:K11"></Entity>`, string(encoded))

	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
