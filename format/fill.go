package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
	"github.com/plandem/xlsx/format/internal/color"
)

type fillOption byte

//Fill is a 'namespace' for all possible settings for fill
var Fill fillOption

func (p *fillOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.Fill.Pattern.Color = color.New(rgb)
	}
}

func (p *fillOption) Background(rgb string) option {
	return func(s *StyleFormat) {
		s.Fill.Pattern.Background = color.New(rgb)
	}
}

func (p *fillOption) Type(pt styles.PatternType) option {
	return func(s *StyleFormat) {
		s.Fill.Pattern.Type = pt
	}
}
