package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/conditional/rule"
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
		"3Arrows":         rule.IconSetType3Arrows,
		"3ArrowsGray":     rule.IconSetType3ArrowsGray,
		"3Flags":          rule.IconSetType3Flags,
		"3TrafficLights1": rule.IconSetType3TrafficLights1,
		"3TrafficLights2": rule.IconSetType3TrafficLights2,
		"3Signs":          rule.IconSetType3Signs,
		"3Symbols":        rule.IconSetType3Symbols,
		"3Symbols2":       rule.IconSetType3Symbols2,
		"4Arrows":         rule.IconSetType4Arrows,
		"4ArrowsGray":     rule.IconSetType4ArrowsGray,
		"4RedToBlack":     rule.IconSetType4RedToBlack,
		"4Rating":         rule.IconSetType4Rating,
		"4TrafficLights":  rule.IconSetType4TrafficLights,
		"5Arrows":         rule.IconSetType5Arrows,
		"5ArrowsGray":     rule.IconSetType5ArrowsGray,
		"5Rating":         rule.IconSetType5Rating,
		"5Quarters":       rule.IconSetType5Quarters,
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
