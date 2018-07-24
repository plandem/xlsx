package format

import (
	"github.com/plandem/xlsx/format/internal/hash"
	"github.com/plandem/xlsx/internal/ml"
)

//StyleRefID is alias of original ml.StyleRefID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for StyleFormat via style-sheet
// 4) put everything related to stylesheet to same package
type StyleRefID ml.StyleRefID

//StyleFormat is objects that holds combined information about cell styling
type StyleFormat struct {
	key string

	Font       ml.Font
	Fill       ml.Fill
	Alignment  ml.CellAlignment
	NumFormat  ml.NumberFormat
	Protection ml.CellProtection
	Border     ml.Border
}

type option func(o *StyleFormat)

//New creates and returns StyleFormat object with requested options
func New(options ...option) *StyleFormat {
	s := &StyleFormat{
		Fill: ml.Fill{
			Pattern:  &ml.PatternFill{},
			Gradient: &ml.GradientFill{},
		},
		Border: ml.Border{
			Left:       &ml.BorderSegment{},
			Right:      &ml.BorderSegment{},
			Top:        &ml.BorderSegment{},
			Bottom:     &ml.BorderSegment{},
			Diagonal:   &ml.BorderSegment{},
			Vertical:   &ml.BorderSegment{},
			Horizontal: &ml.BorderSegment{},
		},
	}
	s.Set(options...)
	return s
}

//Key returns unique hash for style settings
func (s *StyleFormat) Key() string {
	return s.key
}

//Set sets new options for style
func (s *StyleFormat) Set(options ...option) {
	for _, o := range options {
		o(s)
	}

	s.key = hash.Style(&s.Font, &s.Fill, &s.Alignment, &s.NumFormat, &s.Protection, &s.Border)
}
