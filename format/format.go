package format

import (
	"github.com/plandem/xlsx/format/internal/hash"
	"github.com/plandem/xlsx/internal/ml"
	"reflect"
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
	s := &StyleFormat{}
	s.Set(options...)
	return s
}

//Key returns unique hash for style settings
func (s *StyleFormat) Key() string {
	return s.key
}

//beforeSet init nested data to simplify code around it
func (s *StyleFormat) beforeSet() {
	//unpack fill
	if s.fill.Pattern == nil {
		s.fill.Pattern = &ml.PatternFill{}
	}

	if s.fill.Gradient == nil {
		s.fill.Gradient = &ml.GradientFill{}
	}

	//unpack border
	if s.border.Left == nil {
		s.border.Left = &ml.BorderSegment{}
	}

	if s.border.Right == nil {
		s.border.Right = &ml.BorderSegment{}
	}

	if s.border.Top == nil {
		s.border.Top = &ml.BorderSegment{}
	}

	if s.border.Bottom == nil {
		s.border.Bottom = &ml.BorderSegment{}
	}

	if s.border.Diagonal == nil {
		s.border.Diagonal = &ml.BorderSegment{}
	}

	if s.border.Vertical == nil {
		s.border.Vertical = &ml.BorderSegment{}
	}

	if s.border.Horizontal == nil {
		s.border.Horizontal = &ml.BorderSegment{}
	}
}

//afterSet remove empty nested structures
func (s *StyleFormat) afterSet() {
	//pack fill
	if s.fill.Pattern != nil && *s.fill.Pattern == (ml.PatternFill{}) {
		s.fill.Pattern = nil
	}

	if s.fill.Gradient != nil && reflect.DeepEqual(s.fill.Gradient, &ml.GradientFill{}) {
		s.fill.Gradient = nil
	}

	//pack border
	if s.border.Left != nil && *s.border.Left == (ml.BorderSegment{}) {
		s.border.Left = nil
	}

	if s.border.Right != nil && *s.border.Right == (ml.BorderSegment{}) {
		s.border.Right = nil
	}

	if s.border.Top != nil && *s.border.Top == (ml.BorderSegment{}) {
		s.border.Top = nil
	}

	if s.border.Bottom != nil && *s.border.Bottom == (ml.BorderSegment{}) {
		s.border.Bottom = nil
	}

	if s.border.Diagonal != nil && *s.border.Diagonal == (ml.BorderSegment{}) {
		s.border.Diagonal = nil
	}

	if s.border.Vertical != nil && *s.border.Vertical == (ml.BorderSegment{}) {
		s.border.Vertical = nil
	}

	if s.border.Horizontal != nil && *s.border.Horizontal == (ml.BorderSegment{}) {
		s.border.Horizontal = nil
	}
}

//Set sets new options for style
func (s *StyleFormat) Set(options ...option) {
	//N.B.: performance for this package is not so critical, so let's unpack/pack nested structures each call
	s.beforeSet()
	for _, o := range options {
		o(s)
	}
	s.afterSet()

	s.key = hash.Style(&s.font, &s.fill, &s.alignment, &s.numFormat, &s.protection, &s.border)
}

//Pack pack current style settings and returns only non-empty objects
func (s *StyleFormat) Pack() (font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, number *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border) {
	if (s.font != ml.Font{} && s.font != ml.Font{Size: 0, Family: 0, Charset: 0}) {
		font = &ml.Font{}
		*font = s.font
	}

	return
}
