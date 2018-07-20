package types_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestText(t *testing.T) {
	type Entity struct {
		Text types.Text `xml:"text"`
	}

	entity := Entity{Text: "common text"}
	encoded, err := xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity><text>common text</text></Entity>`, string(encoded))

	entity = Entity{Text: " text with space"}
	encoded, err = xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity><text xml:space="preserve"> text with space</text></Entity>`, string(encoded))

	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
