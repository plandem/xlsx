package format

import (
	"github.com/plandem/xlsx/internal/ml"
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

//StyleFormat is objects that holds combined information about cell styling
type StyleFormat struct {
	styleInfo *ml.DiffStyle
	namedInfo *ml.NamedStyleInfo
}

type option func(o *StyleFormat)

//New creates and returns StyleFormat object with requested options
func New(options ...option) *StyleFormat {
	s := &StyleFormat{
		&ml.DiffStyle{
			NumberFormat: &ml.NumberFormat{},
			Font:         &ml.Font{},
			Fill: &ml.Fill{
				Pattern:  &ml.PatternFill{},
				Gradient: &ml.GradientFill{},
			},
			Border: &ml.Border{
				Left:       &ml.BorderSegment{},
				Right:      &ml.BorderSegment{},
				Top:        &ml.BorderSegment{},
				Bottom:     &ml.BorderSegment{},
				Diagonal:   &ml.BorderSegment{},
				Vertical:   &ml.BorderSegment{},
				Horizontal: &ml.BorderSegment{},
			},
			Alignment:  &ml.CellAlignment{},
			Protection: &ml.CellProtection{},
		},
		&ml.NamedStyleInfo{},
	}
	s.Set(options...)
	return s
}

//Set sets new options for style
func (s *StyleFormat) Set(options ...option) {
	for _, o := range options {
		o(s)
	}
}
