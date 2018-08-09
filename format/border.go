package format

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type borderTopSegmentOption byte
type borderBottomSegmentOption byte
type borderLeftSegmentOption byte
type borderRightSegmentOption byte
type borderDiagonalSegmentOption byte
type borderVerticalSegmentOption byte
type borderHorizontalSegmentOption byte

type borderOption struct {
	Top        borderTopSegmentOption
	Bottom     borderBottomSegmentOption
	Left       borderLeftSegmentOption
	Right      borderRightSegmentOption
	Diagonal   borderDiagonalSegmentOption
	Vertical   borderVerticalSegmentOption
	Horizontal borderHorizontalSegmentOption
}

//Border is a 'namespace' for all possible settings for border
var Border borderOption

func (b *borderOption) DiagonalUp(s *StyleFormat) {
	s.border.DiagonalUp = true
}

func (b *borderOption) DiagonalDown(s *StyleFormat) {
	s.border.DiagonalDown = true
}

func (b *borderOption) Outline(s *StyleFormat) {
	s.border.Outline = true
}

func (b *borderOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.border.Top.Type = t
		s.border.Bottom.Type = t
		s.border.Left.Type = t
		s.border.Right.Type = t
	}
}

func (b *borderOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		rgb := color.New(rgb)
		s.border.Top.Color = rgb
		s.border.Bottom.Color = rgb
		s.border.Left.Color = rgb
		s.border.Right.Color = rgb
	}
}

func (b *borderTopSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.border.Top.Type = t
	}
}

func (b *borderTopSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.border.Top.Color = color.New(rgb)
	}
}

func (b *borderBottomSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.border.Bottom.Type = t
	}
}

func (b *borderBottomSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.border.Bottom.Color = color.New(rgb)
	}
}

func (b *borderLeftSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.border.Left.Type = t
	}
}

func (b *borderLeftSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.border.Left.Color = color.New(rgb)
	}
}

func (b *borderRightSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.border.Right.Type = t
	}
}

func (b *borderRightSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.border.Right.Color = color.New(rgb)
	}
}

func (b *borderDiagonalSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.border.Diagonal.Type = t
	}
}

func (b *borderDiagonalSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.border.Diagonal.Color = color.New(rgb)
	}
}

func (b *borderVerticalSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.border.Vertical.Type = t
	}
}

func (b *borderVerticalSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.border.Vertical.Color = color.New(rgb)
	}
}

func (b *borderHorizontalSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.border.Horizontal.Type = t
	}
}

func (b *borderHorizontalSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.border.Horizontal.Color = color.New(rgb)
	}
}
