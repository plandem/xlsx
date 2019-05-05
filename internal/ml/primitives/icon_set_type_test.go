package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/conditional"
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
		"3Arrows":         conditional.IconSetType3Arrows,
		"3ArrowsGray":     conditional.IconSetType3ArrowsGray,
		"3Flags":          conditional.IconSetType3Flags,
		"3TrafficLights1": conditional.IconSetType3TrafficLights1,
		"3TrafficLights2": conditional.IconSetType3TrafficLights2,
		"3Signs":          conditional.IconSetType3Signs,
		"3Symbols":        conditional.IconSetType3Symbols,
		"3Symbols2":       conditional.IconSetType3Symbols2,
		"4Arrows":         conditional.IconSetType4Arrows,
		"4ArrowsGray":     conditional.IconSetType4ArrowsGray,
		"4RedToBlack":     conditional.IconSetType4RedToBlack,
		"4Rating":         conditional.IconSetType4Rating,
		"4TrafficLights":  conditional.IconSetType4TrafficLights,
		"5Arrows":         conditional.IconSetType5Arrows,
		"5ArrowsGray":     conditional.IconSetType5ArrowsGray,
		"5Rating":         conditional.IconSetType5Rating,
		"5Quarters":       conditional.IconSetType5Quarters,
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
