package format_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBorderStyle(t *testing.T) {
	type Entity struct {
		Attribute ml.BorderStyleType `xml:"attribute,attr"`
	}

	list := map[string]ml.BorderStyleType{
		"":                 ml.BorderStyleType(0),
		"none":             format.BorderStyleNone,
		"thin":             format.BorderStyleThin,
		"medium":           format.BorderStyleMedium,
		"dashed":           format.BorderStyleDashed,
		"dotted":           format.BorderStyleDotted,
		"thick":            format.BorderStyleThick,
		"double":           format.BorderStyleDouble,
		"hair":             format.BorderStyleHair,
		"mediumDashed":     format.BorderStyleMediumDashed,
		"dashDot":          format.BorderStyleDashDot,
		"mediumDashDot":    format.BorderStyleMediumDashDot,
		"dashDotDot":       format.BorderStyleDashDotDot,
		"mediumDashDotDot": format.BorderStyleMediumDashDotDot,
		"slantDashDot":     format.BorderStyleSlantDashDot,
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
