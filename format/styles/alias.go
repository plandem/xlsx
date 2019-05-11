package styles

import (
	"github.com/plandem/xlsx/internal/ml"
)

//DirectStyleID is alias of original ml.DirectStyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for Info via style-sheet
// 4) put everything related to stylesheet to same package
type DirectStyleID = ml.DirectStyleID

//DiffStyleID is alias of original ml.DiffStyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for Info via style-sheet
// 4) put everything related to stylesheet to same package
type DiffStyleID = ml.DiffStyleID

//NamedStyleID is alias of original ml.NamedStyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for Info via style-sheet
// 4) put everything related to stylesheet to same package
type NamedStyleID = ml.NamedStyleID
