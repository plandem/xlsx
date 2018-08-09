package types

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//CellType is alias of original primitives.CellType type to:
// 1) make it public
// 2) forbid usage of integers directly
type CellType = primitives.CellType

//List of all possible values for CellType
const (
	CellTypeGeneral CellType = iota
	CellTypeBool
	CellTypeDate
	CellTypeNumber
	CellTypeError
	CellTypeSharedString
	CellTypeFormula
	CellTypeInlineString
)

func init() {
	primitives.FromCellType = map[CellType]string{
		CellTypeBool:         "b",
		CellTypeDate:         "d",
		CellTypeNumber:       "n",
		CellTypeError:        "e",
		CellTypeSharedString: "s",
		CellTypeFormula:      "str",
		CellTypeInlineString: "inlineStr",
	}

	primitives.ToCellType = make(map[string]CellType, len(primitives.FromCellType))
	for k, v := range primitives.FromCellType {
		primitives.ToCellType[v] = k
	}
}
