package types_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestObjects(t *testing.T) {
	type Entity struct {
		Attribute types.ObjectsType `xml:"attribute,attr"`
	}

	entity := Entity{Attribute: types.ObjectsTypePlaceholders}
	encoded, err := xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity attribute="placeholders"></Entity>`, string(encoded))

	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
