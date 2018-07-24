package format

import (
	"github.com/plandem/xlsx/format/internal/color"
	"github.com/plandem/xlsx/internal/ml/styles"
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

func (b *borderOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Top.Type = t
		s.Border.Bottom.Type = t
		s.Border.Left.Type = t
		s.Border.Right.Type = t
	}
}

func (b *borderOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		rgb := color.New(rgb)
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

func (b *borderTopSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.Border.Top.Color = color.New(rgb)
	}
}

func (b *borderBottomSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Bottom.Type = t
	}
}

func (b *borderBottomSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.Border.Bottom.Color = color.New(rgb)
	}
}

func (b *borderLeftSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Left.Type = t
	}
}

func (b *borderLeftSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.Border.Left.Color = color.New(rgb)
	}
}

func (b *borderRightSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Right.Type = t
	}
}

func (b *borderRightSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.Border.Right.Color = color.New(rgb)
	}
}

func (b *borderDiagonalSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Diagonal.Type = t
	}
}

func (b *borderDiagonalSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.Border.Diagonal.Color = color.New(rgb)
	}
}

func (b *borderVerticalSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Vertical.Type = t
	}
}

func (b *borderVerticalSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.Border.Vertical.Color = color.New(rgb)
	}
}

func (b *borderHorizontalSegmentOption) Type(t styles.BorderStyleType) option {
	return func(s *StyleFormat) {
		s.Border.Horizontal.Type = t
	}
}

func (b *borderHorizontalSegmentOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.Border.Horizontal.Color = color.New(rgb)
	}
}
