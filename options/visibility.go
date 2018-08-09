package options

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for VisibilityType
const (
	_ primitives.VisibilityType = iota
	VisibilityTypeVisible
	VisibilityTypeHidden
	VisibilityTypeVeryHidden
)

func init() {
	primitives.FromVisibilityType = map[primitives.VisibilityType]string{
		VisibilityTypeVisible:    "visible",
		VisibilityTypeHidden:     "hidden",
		VisibilityTypeVeryHidden: "veryHidden",
	}

	primitives.ToVisibilityType = make(map[string]primitives.VisibilityType, len(primitives.FromVisibilityType))
	for k, v := range primitives.FromVisibilityType {
		primitives.ToVisibilityType[v] = k
	}
}
