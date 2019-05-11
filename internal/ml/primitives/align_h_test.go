package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignH(t *testing.T) {
	type Entity struct {
		Attribute primitives.HAlignType `xml:"attribute,attr"`
	}

	list := map[string]primitives.HAlignType{
		"":                 primitives.HAlignType(0),
		"general":          styles.HAlignGeneral,
		"left":             styles.HAlignLeft,
		"center":           styles.HAlignCenter,
		"right":            styles.HAlignRight,
		"fill":             styles.HAlignFill,
		"justify":          styles.HAlignJustify,
		"centerContinuous": styles.HAlignCenterContinuous,
		"distributed":      styles.HAlignDistributed,
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
