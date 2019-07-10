// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type alignmentOption byte

//Alignment is a 'namespace' for all possible settings for alignment
var Alignment alignmentOption

func (f *alignmentOption) VAlign(va primitives.VAlignType) Option {
	return func(s *Info) {
		s.styleInfo.Alignment.Vertical = va
	}
}

func (f *alignmentOption) HAlign(ha primitives.HAlignType) Option {
	return func(s *Info) {
		s.styleInfo.Alignment.Horizontal = ha
	}
}

func (f *alignmentOption) TextRotation(angle int) Option {
	return func(s *Info) {
		s.styleInfo.Alignment.TextRotation = angle
	}
}

func (f *alignmentOption) WrapText(s *Info) {
	s.styleInfo.Alignment.WrapText = true
}

func (f *alignmentOption) Indent(i int) Option {
	return func(s *Info) {
		s.styleInfo.Alignment.Indent = i
	}
}

func (f *alignmentOption) RelativeIndent(i int) Option {
	return func(s *Info) {
		s.styleInfo.Alignment.RelativeIndent = i
	}
}

func (f *alignmentOption) JustifyLastLine(s *Info) {
	s.styleInfo.Alignment.JustifyLastLine = true
}

func (f *alignmentOption) ShrinkToFit(s *Info) {
	s.styleInfo.Alignment.ShrinkToFit = true
}

func (f *alignmentOption) ReadingOrder(i int) Option {
	return func(s *Info) {
		s.styleInfo.Alignment.ReadingOrder = i
	}
}
