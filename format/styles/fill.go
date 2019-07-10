// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type patternOption byte
type gradientOption byte

//N.B.: only one kind of fill is allowed by standard
type fillOption struct {
	Pattern  patternOption
	Gradient gradientOption
}

//Fill is a 'namespace' for all possible settings for fill
var Fill fillOption

func (f *fillOption) Color(rgb string) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Pattern.Color = color.New(rgb)
		s.styleInfo.Fill.Gradient = &ml.GradientFill{}
	}
}

func (f *fillOption) Background(rgb string) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Pattern.Background = color.New(rgb)
		s.styleInfo.Fill.Gradient = &ml.GradientFill{}
	}
}

func (f *fillOption) Type(pt primitives.PatternType) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Pattern.Type = pt
		s.styleInfo.Fill.Gradient = &ml.GradientFill{}
	}
}

func (p *patternOption) Color(rgb string) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Pattern.Color = color.New(rgb)
		s.styleInfo.Fill.Gradient = &ml.GradientFill{}
	}
}

func (p *patternOption) Background(rgb string) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Pattern.Background = color.New(rgb)
		s.styleInfo.Fill.Gradient = &ml.GradientFill{}
	}
}

func (p *patternOption) Type(pt primitives.PatternType) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Pattern.Type = pt
		s.styleInfo.Fill.Gradient = &ml.GradientFill{}
	}
}

func (g *gradientOption) Type(gt primitives.GradientType) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Gradient.Type = gt
		s.styleInfo.Fill.Pattern = &ml.PatternFill{}
	}
}

func (g *gradientOption) Degree(degree float64) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Gradient.Degree = degree
		s.styleInfo.Fill.Pattern = &ml.PatternFill{}
	}
}

func (g *gradientOption) Left(left float64) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Gradient.Left = left
		s.styleInfo.Fill.Pattern = &ml.PatternFill{}
	}
}

func (g *gradientOption) Right(right float64) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Gradient.Right = right
		s.styleInfo.Fill.Pattern = &ml.PatternFill{}
	}
}

func (g *gradientOption) Top(top float64) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Gradient.Top = top
		s.styleInfo.Fill.Pattern = &ml.PatternFill{}
	}
}

func (g *gradientOption) Bottom(bottom float64) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Gradient.Bottom = bottom
		s.styleInfo.Fill.Pattern = &ml.PatternFill{}
	}
}

func (g *gradientOption) Stop(position float64, rgb string) Option {
	return func(s *Info) {
		s.styleInfo.Fill.Gradient.Stop = append(s.styleInfo.Fill.Gradient.Stop, &ml.GradientStop{Position: position, Color: color.New(rgb)})
		s.styleInfo.Fill.Pattern = &ml.PatternFill{}
	}
}
