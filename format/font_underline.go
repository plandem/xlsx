package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for UnderlineType
const (
	UnderlineTypeSingle           primitives.UnderlineType = "single"
	UnderlineTypeDouble           primitives.UnderlineType = "double"
	UnderlineTypeSingleAccounting primitives.UnderlineType = "singleAccounting"
	UnderlineTypeDoubleAccounting primitives.UnderlineType = "doubleAccounting"
	UnderlineTypeNone             primitives.UnderlineType = "none"
)
