package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCellType(t *testing.T) {
	type Entity struct {
		Attribute primitives.CellType `xml:"attribute,attr"`
	}

	list := map[primitives.CellType]string{
		types.CellTypeGeneral:      types.CellTypeGeneral.String(),
		types.CellTypeBool:         types.CellTypeBool.String(),
		types.CellTypeDate:         types.CellTypeDate.String(),
		types.CellTypeNumber:       types.CellTypeNumber.String(),
		types.CellTypeError:        types.CellTypeError.String(),
		types.CellTypeSharedString: types.CellTypeSharedString.String(),
		types.CellTypeFormula:      types.CellTypeFormula.String(),
		types.CellTypeInlineString: types.CellTypeInlineString.String(),
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

			if s == "" {
				require.Equal(tt, Entity{}, decoded)
			} else {
				require.Equal(tt, entity, decoded)
			}

			require.Equal(tt, s, decoded.Attribute.String())
		})
	}
}
