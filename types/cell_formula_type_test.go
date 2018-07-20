package types_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
	"fmt"
)

func TestCellFormulaType(t *testing.T) {
	type Entity struct {
		Attribute types.CellFormulaType `xml:"attribute,attr"`
	}

	list := map[string] types.CellFormulaType{
		"normal": types.CellFormulaTypeNormal,
		"array": types.CellFormulaTypeArray,
		"dataTable": types.CellFormulaTypeDataTable,
		"shared": types.CellFormulaTypeShared,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T){
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
