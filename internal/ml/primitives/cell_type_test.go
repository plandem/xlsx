package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCellType(t *testing.T) {
	type Entity struct {
		Attribute primitives.CellType `xml:"attribute,attr"`
	}

	list := map[string]primitives.CellType{
		"":          primitives.CellTypeGeneral,
		"b":         primitives.CellTypeBool,
		"d":         primitives.CellTypeDate,
		"n":         primitives.CellTypeNumber,
		"e":         primitives.CellTypeError,
		"s":         primitives.CellTypeSharedString,
		"str":       primitives.CellTypeFormula,
		"inlineStr": primitives.CellTypeInlineString,
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

			if s == "" {
				require.Equal(tt, Entity{}, decoded)
			} else {
				require.Equal(tt, entity, decoded)
			}

			require.Equal(tt, s, decoded.Attribute.String())
		})
	}
}
