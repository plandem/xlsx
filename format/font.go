package format

import (
	"github.com/plandem/xlsx/internal/ml"
	sharedML "github.com/plandem/ooxml/ml"
)

type fontOption byte

//Font is a 'namespace' for all possible settings for font
var Font fontOption

func (f *fontOption) Name(name string) option {
	return func(s *StyleFormat) {
		s.Font.Name = sharedML.Property(name)
	}
}

func (f *fontOption) Bold(s *StyleFormat) {
	s.Font.Bold = true
}

func (f *fontOption) Italic(s *StyleFormat) {
	s.Font.Italic = true
}

func (f *fontOption) Strikeout(s *StyleFormat) {
	s.Font.Strike = true
}

func (f *fontOption) Shadow(s *StyleFormat) {
	s.Font.Shadow = true
}

func (f *fontOption) Condense(s *StyleFormat) {
	s.Font.Condense = true
}

func (f *fontOption) Extend(s *StyleFormat) {
	s.Font.Extend = true
}

func (f *fontOption) Family(family ml.FontFamilyType) option {
	return func(s *StyleFormat) {
		s.Font.Family = family
	}
}

func (f *fontOption) Color(color string) option {
	return func(s *StyleFormat) {
		//s.Font.Color = newColor(color)
	}
}

func (f *fontOption) Size(size float64) option {
	return func(s *StyleFormat) {
		s.Font.Size = sharedML.PropertyDouble(size)
	}
}

func (f *fontOption) Underline(ut ml.UnderlineType) option {
	return func(s *StyleFormat) {
		s.Font.Underline = ut
	}
}

func (f *fontOption) VAlign(va ml.FontVAlignType) option {
	return func(s *StyleFormat) {
		s.Font.VAlign = va
	}
}

func (f *fontOption) Scheme(sn ml.FontSchemeType) option {
	return func(s *StyleFormat) {
		s.Font.Scheme = sn
	}
}
