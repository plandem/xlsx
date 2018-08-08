package types_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/internal/ml/types"
	"github.com/stretchr/testify/require"
	"testing"
	"github.com/plandem/xlsx/options"
)

func TestVisibility(t *testing.T) {
	type Entity struct {
		Attribute types.VisibilityType `xml:"attribute,attr"`
	}

	list := map[string]types.VisibilityType{
		"":           types.VisibilityType(0),
		"visible":    options.VisibilityTypeVisible,
		"hidden":     options.VisibilityTypeHidden,
		"veryHidden": options.VisibilityTypeVeryHidden,
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
