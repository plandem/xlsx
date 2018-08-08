package primitives_test

import (
	"encoding/xml"
	"github.com/stretchr/testify/require"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"testing"
)

func TestText(t *testing.T) {
	type Entity struct {
		Text primitives.Text `xml:"text"`
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
