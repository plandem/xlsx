package format

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//DirectStyleID is alias of original ml.DirectStyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for StyleFormat via style-sheet
// 4) put everything related to stylesheet to same package
type DirectStyleID = ml.DirectStyleID

//DiffStyleID is alias of original ml.DiffStyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for StyleFormat via style-sheet
// 4) put everything related to stylesheet to same package
type DiffStyleID = ml.DiffStyleID

//NamedStyleID is alias of original ml.NamedStyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for StyleFormat via style-sheet
// 4) put everything related to stylesheet to same package
type NamedStyleID = ml.NamedStyleID

//ConditionType is alias of original primitives.ConditionType
type ConditionType = primitives.ConditionType

//ConditionOperatorType is alias of original primitives.ConditionOperatorType
type ConditionOperatorType = primitives.ConditionOperatorType

//TimePeriodType is alias of original primitives.TimePeriodType
type TimePeriodType = primitives.TimePeriodType

//Formula is alias of original primitives.Formula
type Formula = primitives.Formula
