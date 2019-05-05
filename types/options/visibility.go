package options

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for VisibilityType
const (
	_ primitives.VisibilityType = iota
	Visible
	Hidden
	VeryHidden
)

func init() {
	primitives.FromVisibilityType = map[primitives.VisibilityType]string{
		Visible:    "visible",
		Hidden:     "hidden",
		VeryHidden: "veryHidden",
	}

	primitives.ToVisibilityType = make(map[string]primitives.VisibilityType, len(primitives.FromVisibilityType))
	for k, v := range primitives.FromVisibilityType {
		primitives.ToVisibilityType[v] = k
	}
}
