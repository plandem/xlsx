package primitives

import (
	"github.com/plandem/ooxml/ml"
)

//OptionalBool is helper alias for ml.OptionalBool from core package
func OptionalBool(v bool) *bool {
	return &v
}

//OptionalIndex is helper alias for ml.OptionalIndex from core package
func OptionalIndex(v int) ml.OptionalIndex {
	return ml.OptionalIndex(&v)
}
