package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
)

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

func (b *borderOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Top.Type = t
		s.Border.Bottom.Type = t
		s.Border.Left.Type = t
		s.Border.Right.Type = t
	}
}

func (b *borderOption) Color(color string) option {
	return func(s *StyleFormat) {
		rgb := newColor(color)
		s.Border.Top.Color = rgb
		s.Border.Bottom.Color = rgb
		s.Border.Left.Color = rgb
		s.Border.Right.Color = rgb
	}
}

func (b *borderTopSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Top.Type = t
	}
}

func (b *borderTopSegmentOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Border.Top.Color = newColor(color)
	}
}

func (b *borderBottomSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Bottom.Type = t
	}
}

func (b *borderBottomSegmentOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Border.Bottom.Color = newColor(color)
	}
}

func (b *borderLeftSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Left.Type = t
	}
}

func (b *borderLeftSegmentOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Border.Left.Color = newColor(color)
	}
}

func (b *borderRightSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Right.Type = t
	}
}

func (b *borderRightSegmentOption) Color(color string) option {
	return func(s *StyleFormat) {
		s.Border.Right.Color = newColor(color)
	}
}
