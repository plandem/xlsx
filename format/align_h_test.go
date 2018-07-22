package format_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignH(t *testing.T) {
	type Entity struct {
		Attribute ml.HAlignType `xml:"attribute,attr"`
	}

	list := map[string]ml.HAlignType{
		"":                 ml.HAlignType(0),
		"general":          format.HAlignGeneral,
		"left":             format.HAlignLeft,
		"center":           format.HAlignCenter,
		"right":            format.HAlignRight,
		"fill":             format.HAlignFill,
		"justify":          format.HAlignJustify,
		"centerContinuous": format.HAlignCenterContinuous,
		"distributed":      format.HAlignDistributed,
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
