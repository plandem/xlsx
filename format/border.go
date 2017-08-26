package format

type border struct {
	Top    borderSegment
	Bottom borderSegment
	Left   borderSegment
	Right  borderSegment
}

type borderSegment struct {
	Color ARGB
	Type  BorderStyleType
}

type borderTopSegmentOption byte
type borderBottomSegmentOption byte
type borderLeftSegmentOption byte
type borderRightSegmentOption byte
type borderOption struct {
	Top    borderTopSegmentOption
	Bottom borderBottomSegmentOption
	Left   borderLeftSegmentOption
	Right  borderRightSegmentOption
}

//Border is a 'namespace' for all possible settings for border
var Border borderOption

func (b *borderOption) Type(t BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Top.Type = t
		s.Border.Bottom.Type = t
		s.Border.Left.Type = t
		s.Border.Right.Type = t
	}
}

func (b *borderOption) Color(color string) option {
	return func(s *StyleFormat) {
		rgb := ColorToARGB(color)
		s.Border.Top.Color = rgb
		s.Border.Bottom.Color = rgb
		s.Border.Left.Color = rgb
		s.Border.Right.Color = rgb
	}
}

func (b *borderTopSegmentOption) Type(t BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Top.Type = t
	}
}

func (b *borderTopSegmentOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Border.Top.Color = ColorToARGB(color)
	}
}

func (b *borderBottomSegmentOption) Type(t BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Bottom.Type = t
	}
}

func (b *borderBottomSegmentOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Border.Bottom.Color = ColorToARGB(color)
	}
}

func (b *borderLeftSegmentOption) Type(t BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Left.Type = t
	}
}

func (b *borderLeftSegmentOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Border.Left.Color = ColorToARGB(color)
	}
}

func (b *borderRightSegmentOption) Type(t BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Right.Type = t
	}
}

func (b *borderRightSegmentOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Border.Right.Color = ColorToARGB(color)
	}
}
