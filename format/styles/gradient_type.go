// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

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
