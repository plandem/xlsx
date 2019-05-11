package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignV(t *testing.T) {
	type Entity struct {
		Attribute primitives.VAlignType `xml:"attribute,attr"`
	}

	list := map[string]primitives.VAlignType{
		"":            primitives.VAlignType(0),
		"top":         styles.VAlignTop,
		"center":      styles.VAlignCenter,
		"bottom":      styles.VAlignBottom,
		"justify":     styles.VAlignJustify,
		"distributed": styles.VAlignDistributed,
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
