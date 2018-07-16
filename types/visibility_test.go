package types_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVisibility(t *testing.T) {
	type Entity struct {
		Attribute types.VisibilityType `xml:"attribute,attr"`
	}

	entity := Entity{Attribute: types.VisibilityTypeVeryHidden}
	encoded, err := xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity attribute="veryHidden"></Entity>`, string(encoded))

	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
