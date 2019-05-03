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
	s.styleInfo.Border.DiagonalUp = true
}

func (b *borderOption) DiagonalDown(s *StyleFormat) {
	s.styleInfo.Border.DiagonalDown = true
}

func (b *borderOption) Outline(s *StyleFormat) {
	s.styleInfo.Border.Outline = true
}

func (b *borderOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Top.Type = t
		s.styleInfo.Border.Bottom.Type = t
		s.styleInfo.Border.Left.Type = t
		s.styleInfo.Border.Right.Type = t
	}
}

func (b *borderOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		rgb := color.New(rgb)
		s.styleInfo.Border.Top.Color = rgb
		s.styleInfo.Border.Bottom.Color = rgb
		s.styleInfo.Border.Left.Color = rgb
		s.styleInfo.Border.Right.Color = rgb
	}
}

func (b *borderTopSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Top.Type = t
	}
}

func (b *borderTopSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Top.Color = color.New(rgb)
	}
}

func (b *borderBottomSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Bottom.Type = t
	}
}

func (b *borderBottomSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Bottom.Color = color.New(rgb)
	}
}

func (b *borderLeftSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Left.Type = t
	}
}

func (b *borderLeftSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Left.Color = color.New(rgb)
	}
}

func (b *borderRightSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Right.Type = t
	}
}

func (b *borderRightSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Right.Color = color.New(rgb)
	}
}

func (b *borderDiagonalSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Diagonal.Type = t
	}
}

func (b *borderDiagonalSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Diagonal.Color = color.New(rgb)
	}
}

func (b *borderVerticalSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Vertical.Type = t
	}
}

func (b *borderVerticalSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Vertical.Color = color.New(rgb)
	}
}

func (b *borderHorizontalSegmentOption) Type(t primitives.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Horizontal.Type = t
	}
}

func (b *borderHorizontalSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.styleInfo.Border.Horizontal.Color = color.New(rgb)
	}
}
