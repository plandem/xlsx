package types_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVisibility(t *testing.T) {
	type Entity struct {
		Attribute types.VisibilityType `xml:"attribute,attr"`
	}

	list := map[string]types.VisibilityType{
		"_":          types.VisibilityType(0),
		"visible":    types.VisibilityTypeVisible,
		"hidden":     types.VisibilityTypeHidden,
		"veryHidden": types.VisibilityTypeVeryHidden,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if s == "_" {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
