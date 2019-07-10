// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

	list := map[primitives.IconSetType]string{
		primitives.IconSetType(0):       "",
		rule.IconSetType3Arrows:         rule.IconSetType3Arrows.String(),
		rule.IconSetType3ArrowsGray:     rule.IconSetType3ArrowsGray.String(),
		rule.IconSetType3Flags:          rule.IconSetType3Flags.String(),
		rule.IconSetType3TrafficLights1: rule.IconSetType3TrafficLights1.String(),
		rule.IconSetType3TrafficLights2: rule.IconSetType3TrafficLights2.String(),
		rule.IconSetType3Signs:          rule.IconSetType3Signs.String(),
		rule.IconSetType3Symbols:        rule.IconSetType3Symbols.String(),
		rule.IconSetType3Symbols2:       rule.IconSetType3Symbols2.String(),
		rule.IconSetType4Arrows:         rule.IconSetType4Arrows.String(),
		rule.IconSetType4ArrowsGray:     rule.IconSetType4ArrowsGray.String(),
		rule.IconSetType4RedToBlack:     rule.IconSetType4RedToBlack.String(),
		rule.IconSetType4Rating:         rule.IconSetType4Rating.String(),
		rule.IconSetType4TrafficLights:  rule.IconSetType4TrafficLights.String(),
		rule.IconSetType5Arrows:         rule.IconSetType5Arrows.String(),
		rule.IconSetType5ArrowsGray:     rule.IconSetType5ArrowsGray.String(),
		rule.IconSetType5Rating:         rule.IconSetType5Rating.String(),
		rule.IconSetType5Quarters:       rule.IconSetType5Quarters.String(),
	}

	for v, s := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if v == 0 {
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
