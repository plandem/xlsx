package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignV(t *testing.T) {
	type Entity struct {
		Attribute primitives.VAlignType `xml:"attribute,attr"`
	}

	list := map[primitives.VAlignType]string{
		primitives.VAlignType(0): "",
		styles.VAlignTop:         styles.VAlignTop.String(),
		styles.VAlignCenter:      styles.VAlignCenter.String(),
		styles.VAlignBottom:      styles.VAlignBottom.String(),
		styles.VAlignJustify:     styles.VAlignJustify.String(),
		styles.VAlignDistributed: styles.VAlignDistributed.String(),
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
