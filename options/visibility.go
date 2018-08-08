package options

import (
	"github.com/plandem/xlsx/internal/ml/types"
)

//List of all possible values for VisibilityType
const (
	_ types.VisibilityType = iota
	VisibilityTypeVisible
	VisibilityTypeHidden
	VisibilityTypeVeryHidden
)

func init() {
	types.FromVisibilityType = map[types.VisibilityType]string{
		VisibilityTypeVisible:    "visible",
		VisibilityTypeHidden:     "hidden",
		VisibilityTypeVeryHidden: "veryHidden",
	}

	types.ToVisibilityType = make(map[string]types.VisibilityType, len(types.FromVisibilityType))
	for k, v := range types.FromVisibilityType {
		types.ToVisibilityType[v] = k
	}
}
