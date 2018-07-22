package format

import "github.com/plandem/xlsx/internal/ml"

type alignment struct {
	Horizontal      ml.HAlignType
	Vertical        ml.VAlignType
	TextRotation    int
	WrapText        bool
	Indent          int
	RelativeIndent  int
	JustifyLastLine bool
	ShrinkToFit     bool
	ReadingOrder    int
}

type alignmentOption byte

//Alignment is a 'namespace' for all possible settings for alignment
var Alignment alignmentOption

func (f *alignmentOption) VAlign(va ml.VAlignType) option {
	return func(s *StyleFormat) {
		s.Alignment.Vertical = va
	}
}

func (f *alignmentOption) HAlign(ha ml.HAlignType) option {
	return func(s *StyleFormat) {
		s.Alignment.Horizontal = ha
	}
}

func (f *alignmentOption) TextRotation(angle int) option {
	return func(s *StyleFormat) {
		s.Alignment.TextRotation = angle
	}
}

func (f *alignmentOption) WrapText(s *StyleFormat) {
	s.Alignment.WrapText = true
}

func (f *alignmentOption) Indent(i int) option {
	return func(s *StyleFormat) {
		s.Alignment.Indent = i
	}
}

func (f *alignmentOption) RelativeIndent(i int) option {
	return func(s *StyleFormat) {
		s.Alignment.RelativeIndent = i
	}
}

func (f *alignmentOption) JustifyLastLine(s *StyleFormat) {
	s.Alignment.JustifyLastLine = true
}

func (f *alignmentOption) ShrinkToFit(s *StyleFormat) {
	s.Alignment.ShrinkToFit = true
}

func (f *alignmentOption) ReadingOrder(i int) option {
	return func(s *StyleFormat) {
		s.Alignment.ReadingOrder = i
	}
}
