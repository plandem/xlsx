package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCellFormulaType(t *testing.T) {
	type Entity struct {
		Attribute primitives.CellFormulaType `xml:"attribute,attr"`
	}

	list := map[primitives.CellFormulaType]string{
		primitives.CellFormulaType(0):       "",
		primitives.CellFormulaTypeNormal:    primitives.CellFormulaTypeNormal.String(),
		primitives.CellFormulaTypeArray:     primitives.CellFormulaTypeArray.String(),
		primitives.CellFormulaTypeDataTable: primitives.CellFormulaTypeDataTable.String(),
		primitives.CellFormulaTypeShared:    primitives.CellFormulaTypeShared.String(),
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
