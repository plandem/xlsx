package format

import (
	"github.com/plandem/xlsx/internal/ml"
)

//List of all possible values for UnderlineType
const (
	UnderlineTypeSingle           ml.UnderlineType = "single"
	UnderlineTypeDouble           ml.UnderlineType = "double"
	UnderlineTypeSingleAccounting ml.UnderlineType = "singleAccounting"
	UnderlineTypeDoubleAccounting ml.UnderlineType = "doubleAccounting"
	UnderlineTypeNone             ml.UnderlineType = "none"
)