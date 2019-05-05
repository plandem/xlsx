package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
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
		"none":            styles.PatternTypeNone,
		"solid":           styles.PatternTypeSolid,
		"mediumGray":      styles.PatternTypeMediumGray,
		"darkGray":        styles.PatternTypeDarkGray,
		"lightGray":       styles.PatternTypeLightGray,
		"darkHorizontal":  styles.PatternTypeDarkHorizontal,
		"darkVertical":    styles.PatternTypeDarkVertical,
		"darkDown":        styles.PatternTypeDarkDown,
		"darkUp":          styles.PatternTypeDarkUp,
		"darkGrid":        styles.PatternTypeDarkGrid,
		"darkTrellis":     styles.PatternTypeDarkTrellis,
		"lightHorizontal": styles.PatternTypeLightHorizontal,
		"lightVertical":   styles.PatternTypeLightVertical,
		"lightDown":       styles.PatternTypeLightDown,
		"lightUp":         styles.PatternTypeLightUp,
		"lightGrid":       styles.PatternTypeLightGrid,
		"lightTrellis":    styles.PatternTypeLightTrellis,
		"gray125":         styles.PatternTypeGray125,
		"gray0625":        styles.PatternTypeGray0625,
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
