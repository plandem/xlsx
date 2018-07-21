package types_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestObjects(t *testing.T) {
	type Entity struct {
		Attribute types.ObjectsType `xml:"attribute,attr"`
	}

	list := map[string]types.ObjectsType{
		"":             types.ObjectsType(0),
		"all":          types.ObjectsTypeAll,
		"placeholders": types.ObjectsTypePlaceholders,
		"none":         types.ObjectsTypeNone,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if s == "" {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
			require.Equal(tt, s, decoded.Attribute.String())
		})
	}
}
