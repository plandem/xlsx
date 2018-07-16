package format_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/format"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignV(t *testing.T) {
	type Entity struct {
		Attribute format.VAlignType `xml:"attribute,attr"`
	}

	entity := Entity{Attribute: format.VAlignDistributed}
	encoded, err := xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity attribute="distributed"></Entity>`, string(encoded))

	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
