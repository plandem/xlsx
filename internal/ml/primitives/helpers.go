// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"github.com/plandem/ooxml/ml"
)

//OptionalBool is helper function that allow encode/decode optional boolean, where false should not be omitted - mostly for booleans with default true value
func OptionalBool(v bool) *bool {
	return &v
}

//OptionalIndex is helper alias for ml.OptionalIndex from core package
func OptionalIndex(v int) ml.OptionalIndex {
	return ml.OptionalIndex(&v)
}
