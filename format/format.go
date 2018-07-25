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
	key        string
	font       ml.Font
	fill       ml.Fill
	alignment  ml.CellAlignment
	numFormat  ml.NumberFormat
	protection ml.CellProtection
	border     ml.Border
}

type option func(o *StyleFormat)

//New creates and returns StyleFormat object with requested options
func New(options ...option) *StyleFormat {
	s := &StyleFormat{
		fill: ml.Fill{
			Pattern:  &ml.PatternFill{},
			Gradient: &ml.GradientFill{},
		},
		border: ml.Border{
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

	s.key = hash.Style(&s.font, &s.fill, &s.alignment, &s.numFormat, &s.protection, &s.border)
}

//Pack pack current style settings and returns only non-empty objects
func (s *StyleFormat) Pack() (font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, number *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border) {
	if (s.font != ml.Font{}) {
		*font = s.font
	}

	return
}
