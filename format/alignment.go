package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type alignmentOption byte

//Alignment is a 'namespace' for all possible settings for alignment
var Alignment alignmentOption

func (f *alignmentOption) VAlign(va primitives.VAlignType) option {
	return func(s *StyleFormat) {
		s.alignment.Vertical = va
	}
}

func (f *alignmentOption) HAlign(ha primitives.HAlignType) option {
	return func(s *StyleFormat) {
		s.alignment.Horizontal = ha
	}
}

func (f *alignmentOption) TextRotation(angle int) option {
	return func(s *StyleFormat) {
		s.alignment.TextRotation = angle
	}
}

func (f *alignmentOption) WrapText(s *StyleFormat) {
	s.alignment.WrapText = true
}

func (f *alignmentOption) Indent(i int) option {
	return func(s *StyleFormat) {
		s.alignment.Indent = i
	}
}

func (f *alignmentOption) RelativeIndent(i int) option {
	return func(s *StyleFormat) {
		s.alignment.RelativeIndent = i
	}
}

func (f *alignmentOption) JustifyLastLine(s *StyleFormat) {
	s.alignment.JustifyLastLine = true
}

func (f *alignmentOption) ShrinkToFit(s *StyleFormat) {
	s.alignment.ShrinkToFit = true
}

func (f *alignmentOption) ReadingOrder(i int) option {
	return func(s *StyleFormat) {
		s.alignment.ReadingOrder = i
	}
}
