package format_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/format"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFillPattern(t *testing.T) {
	type Entity struct {
		Attribute format.PatternType `xml:"attribute,attr"`
	}

	entity := Entity{Attribute: format.PatternTypeLightHorizontal}
	encoded, err := xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity attribute="lightHorizontal"></Entity>`, string(encoded))

	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
