// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type fontOption byte

//Font is a 'namespace' for all possible settings for font
var Font fontOption

func (f *fontOption) Name(name string) Option {
	return func(s *Info) {
		s.styleInfo.Font.Name = ml.Property(name)
	}
}

func (f *fontOption) Default(s *Info) {
	s.styleInfo.Font.Family = FontFamilySwiss
	s.styleInfo.Font.Scheme = FontSchemeMinor
	s.styleInfo.Font.Name = "Calibri"
	s.styleInfo.Font.Size = 11.0
	//s.font.Color  =  Color{Theme: 1}
}

func (f *fontOption) Bold(s *Info) {
	s.styleInfo.Font.Bold = true
}

func (f *fontOption) Italic(s *Info) {
	s.styleInfo.Font.Italic = true
}

func (f *fontOption) Strikeout(s *Info) {
	s.styleInfo.Font.Strike = true
}

func (f *fontOption) Superscript(s *Info) {
	s.styleInfo.Font.Effect = primitives.FontEffectSuperscript
}

func (f *fontOption) Subscript(s *Info) {
	s.styleInfo.Font.Effect = primitives.FontEffectSubscript
}

func (f *fontOption) Shadow(s *Info) {
	s.styleInfo.Font.Shadow = true
}

func (f *fontOption) Condense(s *Info) {
	s.styleInfo.Font.Condense = true
}

func (f *fontOption) Extend(s *Info) {
	s.styleInfo.Font.Extend = true
}

func (f *fontOption) Family(family primitives.FontFamilyType) Option {
	return func(s *Info) {
		s.styleInfo.Font.Family = family
	}
}

func (f *fontOption) Color(rgb string) Option {
	return func(s *Info) {
		s.styleInfo.Font.Color = color.New(rgb)
	}
}

func (f *fontOption) Size(size float64) Option {
	return func(s *Info) {
		s.styleInfo.Font.Size = ml.PropertyDouble(size)
	}
}

func (f *fontOption) Underline(underline primitives.UnderlineType) Option {
	return func(s *Info) {
		s.styleInfo.Font.Underline = underline
	}
}

func (f *fontOption) Scheme(scheme primitives.FontSchemeType) Option {
	return func(s *Info) {
		s.styleInfo.Font.Scheme = scheme
	}
}

func (f *fontOption) Charset(charset FontCharsetType) Option {
	return func(s *Info) {
		if charset >= FontCharsetANSI && charset <= FontCharsetOEM {
			//FIXME: right now it's not possible to encode 'Ansi' charset with 'omitempty'
			s.styleInfo.Font.Charset = primitives.FontCharsetType(charset)
		}
	}
}
