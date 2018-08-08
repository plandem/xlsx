package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for GradientType
const (
	GradientTypeLinear primitives.GradientType = iota
	GradientTypePath
)

func init() {
	primitives.FromGradientType = map[primitives.GradientType]string{
		GradientTypeLinear: "linear",
		GradientTypePath:   "path",
	}

	primitives.ToGradientType = make(map[string]primitives.GradientType, len(primitives.FromGradientType))
	for k, v := range primitives.FromGradientType {
		primitives.ToGradientType[v] = k
	}
}
