package format_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/format"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignH(t *testing.T) {
	type Entity struct {
		Attribute format.HAlignType `xml:"attribute,attr"`
	}

	entity := Entity{Attribute: format.HAlignGeneral}
	encoded, err := xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity attribute="general"></Entity>`, string(encoded))

	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
