package format

type font struct {
	Name      string
	Family    FontFamilyType
	Bold      bool
	Italic    bool
	Strike    bool
	Shadow    bool
	Condense  bool
	Extend    bool
	Color     ARGB
	Size      float64
	Underline UnderlineType
	VAlign    FontVAlignType
	Scheme    FontSchemeType
}

type fontOption byte

//Font is a 'namespace' for all possible settings for font
var Font fontOption

func (f *fontOption) Name(name string) option {
	return func(s *StyleFormat) {
		s.Font.Name = name
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

func (f *fontOption) Family(family FontFamilyType) option {
	return func(s *StyleFormat) {
		s.Font.Family = family
	}
}

func (f *fontOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Font.Color = ColorToARGB(color)
	}
}

func (f *fontOption) Size(size float64) option {
	return func(s *StyleFormat) {
		s.Font.Size = size
	}
}

func (f *fontOption) Underline(ut UnderlineType) option {
	return func(s *StyleFormat) {
		s.Font.Underline = ut
	}
}

func (f *fontOption) VAlign(va FontVAlignType) option {
	return func(s *StyleFormat) {
		s.Font.VAlign = va
	}
}

func (f *fontOption) Scheme(sn FontSchemeType) option {
	return func(s *StyleFormat) {
		s.Font.Scheme = sn
	}
}
