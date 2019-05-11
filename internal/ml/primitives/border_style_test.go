package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBorderStyle(t *testing.T) {
	type Entity struct {
		Attribute primitives.BorderStyleType `xml:"attribute,attr"`
	}

	list := map[string]primitives.BorderStyleType{
		"":                 primitives.BorderStyleType(0),
		"none":             styles.BorderStyleNone,
		"thin":             styles.BorderStyleThin,
		"medium":           styles.BorderStyleMedium,
		"dashed":           styles.BorderStyleDashed,
		"dotted":           styles.BorderStyleDotted,
		"thick":            styles.BorderStyleThick,
		"double":           styles.BorderStyleDouble,
		"hair":             styles.BorderStyleHair,
		"mediumDashed":     styles.BorderStyleMediumDashed,
		"dashDot":          styles.BorderStyleDashDot,
		"mediumDashDot":    styles.BorderStyleMediumDashDot,
		"dashDotDot":       styles.BorderStyleDashDotDot,
		"mediumDashDotDot": styles.BorderStyleMediumDashDotDot,
		"slantDashDot":     styles.BorderStyleSlantDashDot,
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
