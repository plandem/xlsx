package format

import (
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type fontOption byte

//Font is a 'namespace' for all possible settings for font
var Font fontOption

func (f *fontOption) Name(name string) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Font.Name = ml.Property(name)
	}
}

func (f *fontOption) Default(s *StyleFormat) {
	s.styleInfo.Font.Family = FontFamilySwiss
	s.styleInfo.Font.Scheme = FontSchemeMinor
	s.styleInfo.Font.Name = "Calibri"
	s.styleInfo.Font.Size = 11.0
	//s.font.Color  =  Color{Theme: 1}
}

func (f *fontOption) Bold(s *StyleFormat) {
	s.styleInfo.Font.Bold = true
}

func (f *fontOption) Italic(s *StyleFormat) {
	s.styleInfo.Font.Italic = true
}

func (f *fontOption) Strikeout(s *StyleFormat) {
	s.styleInfo.Font.Strike = true
}

func (f *fontOption) Shadow(s *StyleFormat) {
	s.styleInfo.Font.Shadow = true
}

func (f *fontOption) Condense(s *StyleFormat) {
	s.styleInfo.Font.Condense = true
}

func (f *fontOption) Extend(s *StyleFormat) {
	s.styleInfo.Font.Extend = true
}

func (f *fontOption) Family(family primitives.FontFamilyType) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Font.Family = family
	}
}

func (f *fontOption) Color(rgb string) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Font.Color = color.New(rgb)
	}
}

func (f *fontOption) Size(size float64) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Font.Size = ml.PropertyDouble(size)
	}
}

func (f *fontOption) Underline(ut primitives.UnderlineType) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Font.Underline = ut
	}
}

func (f *fontOption) VAlign(va primitives.FontVAlignType) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Font.VAlign = va
	}
}

func (f *fontOption) Scheme(sn primitives.FontSchemeType) styleOption {
	return func(s *StyleFormat) {
		s.styleInfo.Font.Scheme = sn
	}
}

func (f *fontOption) Charset(charset FontCharsetType) styleOption {
	return func(s *StyleFormat) {
		if charset >= FontCharsetANSI && charset <= FontCharsetOEM {
			//FIXME: right now it's not possible to encode 'Ansi' charset with 'omitempty'
			s.styleInfo.Font.Charset = primitives.FontCharsetType(charset)
		}
	}
}
