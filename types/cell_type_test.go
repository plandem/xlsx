package types_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCellType(t *testing.T) {
	type Entity struct {
		Attribute types.CellType `xml:"attribute,attr"`
	}

	entity := Entity{Attribute: types.CellTypeInlineString}
	encoded, err := xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity attribute="inlineStr"></Entity>`, string(encoded))

	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
