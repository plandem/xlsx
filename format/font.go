package format

import (
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/format/internal/color"
	"github.com/plandem/xlsx/internal/ml/styles"
)

type fontOption byte

//Font is a 'namespace' for all possible settings for font
var Font fontOption

func (f *fontOption) Name(name string) option {
	return func(s *StyleFormat) {
		s.font.Name = ml.Property(name)
	}
}

func (f *fontOption) Bold(s *StyleFormat) {
	s.font.Bold = true
}

func (f *fontOption) Italic(s *StyleFormat) {
	s.font.Italic = true
}

func (f *fontOption) Strikeout(s *StyleFormat) {
	s.font.Strike = true
}

func (f *fontOption) Shadow(s *StyleFormat) {
	s.font.Shadow = true
}

func (f *fontOption) Condense(s *StyleFormat) {
	s.font.Condense = true
}

func (f *fontOption) Extend(s *StyleFormat) {
	s.font.Extend = true
}

func (f *fontOption) Family(family styles.FontFamilyType) option {
	return func(s *StyleFormat) {
		s.font.Family = family
	}
}

func (f *fontOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.font.Color = color.New(rgb)
	}
}

func (f *fontOption) Size(size float64) option {
	return func(s *StyleFormat) {
		s.font.Size = ml.PropertyDouble(size)
	}
}

func (f *fontOption) Underline(ut styles.UnderlineType) option {
	return func(s *StyleFormat) {
		s.font.Underline = ut
	}
}

func (f *fontOption) VAlign(va styles.FontVAlignType) option {
	return func(s *StyleFormat) {
		s.font.VAlign = va
	}
}

func (f *fontOption) Scheme(sn styles.FontSchemeType) option {
	return func(s *StyleFormat) {
		s.font.Scheme = sn
	}
}
