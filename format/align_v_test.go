package format_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignV(t *testing.T) {
	type Entity struct {
		Attribute format.VAlignType `xml:"attribute,attr"`
	}

	list := map[string]format.VAlignType{
		"_":           format.VAlignType(0),
		"top":         format.VAlignTop,
		"center":      format.VAlignCenter,
		"bottom":      format.VAlignBottom,
		"justify":     format.VAlignJustify,
		"distributed": format.VAlignDistributed,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if s == "_" {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
