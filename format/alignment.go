package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type alignmentOption byte

//Alignment is a 'namespace' for all possible settings for alignment
var Alignment alignmentOption

func (f *alignmentOption) VAlign(va primitives.VAlignType) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Alignment.Vertical = va
	}
}

func (f *alignmentOption) HAlign(ha primitives.HAlignType) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Alignment.Horizontal = ha
	}
}

func (f *alignmentOption) TextRotation(angle int) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Alignment.TextRotation = angle
	}
}

func (f *alignmentOption) WrapText(s *StyleFormat) {
	s.styleInfo.Alignment.WrapText = true
}

func (f *alignmentOption) Indent(i int) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Alignment.Indent = i
	}
}

func (f *alignmentOption) RelativeIndent(i int) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Alignment.RelativeIndent = i
	}
}

func (f *alignmentOption) JustifyLastLine(s *StyleFormat) {
	s.styleInfo.Alignment.JustifyLastLine = true
}

func (f *alignmentOption) ShrinkToFit(s *StyleFormat) {
	s.styleInfo.Alignment.ShrinkToFit = true
}

func (f *alignmentOption) ReadingOrder(i int) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Alignment.ReadingOrder = i
	}
}
