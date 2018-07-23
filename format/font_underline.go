package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
)

//List of all possible values for UnderlineType
const (
	UnderlineTypeSingle           styles.UnderlineType = "single"
	UnderlineTypeDouble           styles.UnderlineType = "double"
	UnderlineTypeSingleAccounting styles.UnderlineType = "singleAccounting"
	UnderlineTypeDoubleAccounting styles.UnderlineType = "doubleAccounting"
	UnderlineTypeNone             styles.UnderlineType = "none"
)