package styles_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml/styles"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFillGradient(t *testing.T) {
	type Entity struct {
		Attribute styles.GradientType `xml:"attribute,attr"`
	}

	list := map[string]styles.GradientType{
		"":       styles.GradientType(0),
		"linear": format.GradientTypeLinear,
		"path":   format.GradientTypePath,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if s == "" {
				require.Equal(tt, `<Entity attribute="linear"></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)

			if s == "" {
				require.Equal(tt, "linear", decoded.Attribute.String())
			} else {
				require.Equal(tt, s, decoded.Attribute.String())
			}
		})
	}
}
