package format

type fill struct {
	Color      ARGB
	Background ARGB
	Type       PatternType
}

type fillOption byte

//Fill is a 'namespace' for all possible settings for fill
var Fill fillOption

func (p *fillOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Fill.Color = ColorToARGB(color)
	}
}

func (p *fillOption) Background(color string) option {
	return func(s *StyleFormat) {
		s.Fill.Background = ColorToARGB(color)
	}
}

func (p *fillOption) Type(pt PatternType) option {
	return func(s *StyleFormat) {
		s.Fill.Type = pt
	}
}
