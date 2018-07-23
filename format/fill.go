package format

import "github.com/plandem/xlsx/internal/ml/styles"

type fillOption byte

//Fill is a 'namespace' for all possible settings for fill
var Fill fillOption

func (p *fillOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Fill.Pattern.Color = newColor(color)
	}
}

func (p *fillOption) Background(color string) option {
	return func(s *StyleFormat) {
		s.Fill.Pattern.Background = newColor(color)
	}
}

func (p *fillOption) Type(pt styles.PatternType) option {
	return func(s *StyleFormat) {
		s.Fill.Pattern.Type = pt
	}
}
