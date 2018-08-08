package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFillPattern(t *testing.T) {
	type Entity struct {
		Attribute primitives.PatternType `xml:"attribute,attr"`
	}

	list := map[string]primitives.PatternType{
		"":                primitives.PatternType(0),
		"none":            format.PatternTypeNone,
		"solid":           format.PatternTypeSolid,
		"mediumGray":      format.PatternTypeMediumGray,
		"darkGray":        format.PatternTypeDarkGray,
		"lightGray":       format.PatternTypeLightGray,
		"darkHorizontal":  format.PatternTypeDarkHorizontal,
		"darkVertical":    format.PatternTypeDarkVertical,
		"darkDown":        format.PatternTypeDarkDown,
		"darkUp":          format.PatternTypeDarkUp,
		"darkGrid":        format.PatternTypeDarkGrid,
		"darkTrellis":     format.PatternTypeDarkTrellis,
		"lightHorizontal": format.PatternTypeLightHorizontal,
		"lightVertical":   format.PatternTypeLightVertical,
		"lightDown":       format.PatternTypeLightDown,
		"lightUp":         format.PatternTypeLightUp,
		"lightGrid":       format.PatternTypeLightGrid,
		"lightTrellis":    format.PatternTypeLightTrellis,
		"gray125":         format.PatternTypeGray125,
		"gray0625":        format.PatternTypeGray0625,
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
