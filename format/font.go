package format

import (
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml/primitives"
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

func (f *fontOption) Family(family primitives.FontFamilyType) option {
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

func (f *fontOption) Underline(ut primitives.UnderlineType) option {
	return func(s *StyleFormat) {
		s.font.Underline = ut
	}
}

func (f *fontOption) VAlign(va primitives.FontVAlignType) option {
	return func(s *StyleFormat) {
		s.font.VAlign = va
	}
}

func (f *fontOption) Scheme(sn primitives.FontSchemeType) option {
	return func(s *StyleFormat) {
		s.font.Scheme = sn
	}
}

func (f *fontOption) Charset(charset FontCharsetType) option {
	return func(s *StyleFormat) {
		if charset >= FontCharsetANSI && charset <= FontCharsetOEM {
			//FIXME: right now it's not possible to encode 'Ansi' charset with 'omitempty'
			s.font.Charset = primitives.FontCharsetType(charset)
		}
	}
}
