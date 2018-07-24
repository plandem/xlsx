package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
)

//List of all possible values for GradientType
const (
	GradientTypeLinear styles.GradientType = iota
	GradientTypePath
)

func init() {
	styles.FromGradientType = map[styles.GradientType]string{
		GradientTypeLinear: "linear",
		GradientTypePath:   "path",
	}

	styles.ToGradientType = make(map[string]styles.GradientType, len(styles.FromGradientType))
	for k, v := range styles.FromGradientType {
		styles.ToGradientType[v] = k
	}
}
