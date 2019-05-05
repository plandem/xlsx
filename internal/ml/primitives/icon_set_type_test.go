package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIconSetType(t *testing.T) {
	type Entity struct {
		Attribute primitives.IconSetType `xml:"attribute,attr"`
	}

	list := map[string]primitives.IconSetType{
		"":                primitives.IconSetType(0),
		"3Arrows":         format.IconSetType3Arrows,
		"3ArrowsGray":     format.IconSetType3ArrowsGray,
		"3Flags":          format.IconSetType3Flags,
		"3TrafficLights1": format.IconSetType3TrafficLights1,
		"3TrafficLights2": format.IconSetType3TrafficLights2,
		"3Signs":          format.IconSetType3Signs,
		"3Symbols":        format.IconSetType3Symbols,
		"3Symbols2":       format.IconSetType3Symbols2,
		"4Arrows":         format.IconSetType4Arrows,
		"4ArrowsGray":     format.IconSetType4ArrowsGray,
		"4RedToBlack":     format.IconSetType4RedToBlack,
		"4Rating":         format.IconSetType4Rating,
		"4TrafficLights":  format.IconSetType4TrafficLights,
		"5Arrows":         format.IconSetType5Arrows,
		"5ArrowsGray":     format.IconSetType5ArrowsGray,
		"5Rating":         format.IconSetType5Rating,
		"5Quarters":       format.IconSetType5Quarters,
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
