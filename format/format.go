package format

import (
	"github.com/plandem/xlsx/internal/ml"
	"reflect"
)

//StyleID is alias of original ml.StyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for StyleFormat via style-sheet
// 4) put everything related to stylesheet to same package
type StyleID = ml.StyleID

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
	//N.B.: performance for this package is not so critical, so let's init/de-init nested structures each call
	s.beforeSet()
	for _, o := range options {
		o(s)
	}
	s.afterSet()
}

//Settings checks current style settings and returns copies of non-empty objects
func (s *StyleFormat) Settings() (font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, numFormat *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border) {
	//copy non-empty alignment
	if s.alignment != (ml.CellAlignment{}) {
		alignment = &ml.CellAlignment{}
		*alignment = s.alignment
	}

	//copy non-empty border
	if s.border != (ml.Border{}) {
		border = &ml.Border{}
		*border = s.border

		if s.border.Left != nil && !reflect.DeepEqual(border.Left, &ml.BorderSegment{}) {
			border.Left = &ml.BorderSegment{}
			*border.Left = *s.border.Left
		}

		if s.border.Right != nil && !reflect.DeepEqual(border.Right, &ml.BorderSegment{}) {
			border.Right = &ml.BorderSegment{}
			*border.Right = *s.border.Right
		}

		if s.border.Top != nil && !reflect.DeepEqual(border.Top, &ml.BorderSegment{}) {
			border.Top = &ml.BorderSegment{}
			*border.Top = *s.border.Top
		}

		if s.border.Bottom != nil && !reflect.DeepEqual(border.Bottom, &ml.BorderSegment{}) {
			border.Bottom = &ml.BorderSegment{}
			*border.Bottom = *s.border.Bottom
		}

		if s.border.Diagonal != nil && !reflect.DeepEqual(border.Diagonal, &ml.BorderSegment{}) {
			border.Diagonal = &ml.BorderSegment{}
			*border.Diagonal = *s.border.Diagonal
		}

		if s.border.Vertical != nil && !reflect.DeepEqual(border.Vertical, &ml.BorderSegment{}) {
			border.Vertical = &ml.BorderSegment{}
			*border.Vertical = *s.border.Vertical
		}

		if s.border.Horizontal != nil && !reflect.DeepEqual(border.Horizontal, &ml.BorderSegment{}) {
			border.Horizontal = &ml.BorderSegment{}
			*border.Horizontal = *s.border.Horizontal
		}
	}

	//copy non-empty fill
	if s.fill != (ml.Fill{}) {
		fill = &ml.Fill{}

		//copy pattern
		if s.fill.Pattern != nil && !reflect.DeepEqual(s.fill.Pattern, &ml.PatternFill{}) {
			fill.Pattern = &ml.PatternFill{}
			*fill.Pattern = *s.fill.Pattern
		}

		//copy gradient
		if s.fill.Gradient != nil && !reflect.DeepEqual(s.fill.Gradient, &ml.GradientFill{}) {
			fill.Gradient = &ml.GradientFill{}
			*fill.Gradient = *s.fill.Gradient
			copy(fill.Gradient.Stop, s.fill.Gradient.Stop)
		}
	}

	//copy non-empty font
	if (s.font != ml.Font{} && s.font != ml.Font{Size: 0, Family: 0, Charset: 0}) {
		font = &ml.Font{}
		*font = s.font
	}

	//copy non-empty numFormat
	if s.numFormat != (ml.NumberFormat{}) {
		numFormat = &ml.NumberFormat{}
		*numFormat = s.numFormat
	}

	//copy non-empty protection
	if s.protection != (ml.CellProtection{}) {
		protection = &ml.CellProtection{}
		*protection = s.protection
	}

	return
}
