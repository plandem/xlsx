// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/ml"
	"reflect"
)

//DefaultDirectStyle is ID for any default direct style than depends on context:
// E.g. for cell it will be equal to NamedStyle 'Normal', for hyperlink - NamedStyle 'Hyperlink'
const DefaultDirectStyle = DirectStyleID(0)

//Info is objects that holds combined information about cell styling
type Info struct {
	styleInfo *ml.DiffStyle
	namedInfo *ml.NamedStyleInfo
}

type Option func(o *Info)

//New creates and returns Info object with requested options
func New(options ...Option) *Info {
	s := &Info{
		&ml.DiffStyle{
			NumberFormat: &ml.NumberFormat{},
			Font:         &ml.Font{},
			Fill: &ml.Fill{
				Pattern:  &ml.PatternFill{},
				Gradient: &ml.GradientFill{},
			},
			Border: &ml.Border{
				Left:       &ml.BorderSegment{},
				Right:      &ml.BorderSegment{},
				Top:        &ml.BorderSegment{},
				Bottom:     &ml.BorderSegment{},
				Diagonal:   &ml.BorderSegment{},
				Vertical:   &ml.BorderSegment{},
				Horizontal: &ml.BorderSegment{},
			},
			Alignment:  &ml.CellAlignment{},
			Protection: &ml.CellProtection{},
		},
		&ml.NamedStyleInfo{},
	}
	s.Set(options...)
	return s
}

//Set sets new options for style
func (s *Info) Set(options ...Option) {
	for _, o := range options {
		o(s)
	}
}

//private method used by stylesheet manager to unpack Info
func from(f *Info) (font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, numFormat *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border, namedInfo *ml.NamedStyleInfo) {
	style := f.styleInfo
	named := f.namedInfo

	//copy non-empty namedInfo
	if *named != (ml.NamedStyleInfo{}) {
		namedInfo = &ml.NamedStyleInfo{}
		*namedInfo = *named
	}

	//copy non-empty alignment
	if *style.Alignment != (ml.CellAlignment{}) {
		alignment = &ml.CellAlignment{}
		*alignment = *style.Alignment
	}

	//copy non-empty font
	if (*style.Font != ml.Font{} && *style.Font != ml.Font{Size: 0, Family: 0, Charset: 0}) {
		font = &ml.Font{}
		*font = *style.Font
	}

	//copy non-empty numFormat
	if *style.NumberFormat != (ml.NumberFormat{}) {
		numFormat = &ml.NumberFormat{}
		*numFormat = *style.NumberFormat
	}

	//copy non-empty protection
	if *style.Protection != (ml.CellProtection{}) {
		protection = &ml.CellProtection{}
		*protection = *style.Protection
	}

	//copy non-empty border
	border = &ml.Border{}
	*border = *style.Border

	if reflect.DeepEqual(border.Left, &ml.BorderSegment{}) {
		border.Left = nil
	} else {
		border.Left = &ml.BorderSegment{}
		*border.Left = *style.Border.Left
	}

	if reflect.DeepEqual(border.Right, &ml.BorderSegment{}) {
		border.Right = nil
	} else {
		border.Right = &ml.BorderSegment{}
		*border.Right = *style.Border.Right
	}

	if reflect.DeepEqual(border.Top, &ml.BorderSegment{}) {
		border.Top = nil
	} else {
		border.Top = &ml.BorderSegment{}
		*border.Top = *style.Border.Top
	}

	if reflect.DeepEqual(border.Bottom, &ml.BorderSegment{}) {
		border.Bottom = nil
	} else {
		border.Bottom = &ml.BorderSegment{}
		*border.Bottom = *style.Border.Bottom
	}

	if reflect.DeepEqual(border.Diagonal, &ml.BorderSegment{}) {
		border.Diagonal = nil
	} else {
		border.Diagonal = &ml.BorderSegment{}
		*border.Diagonal = *style.Border.Diagonal
	}

	if reflect.DeepEqual(border.Vertical, &ml.BorderSegment{}) {
		border.Vertical = nil
	} else {
		border.Vertical = &ml.BorderSegment{}
		*border.Vertical = *style.Border.Vertical
	}

	if reflect.DeepEqual(border.Horizontal, &ml.BorderSegment{}) {
		border.Horizontal = nil
	} else {
		border.Horizontal = &ml.BorderSegment{}
		*border.Horizontal = *style.Border.Horizontal
	}

	//if border is actually empty, then nil it
	if *border == (ml.Border{}) {
		border = nil
	}

	//copy non-empty fill
	fill = &ml.Fill{}

	//copy pattern
	if !reflect.DeepEqual(style.Fill.Pattern, &ml.PatternFill{}) {
		fill.Pattern = &ml.PatternFill{}
		*fill.Pattern = *style.Fill.Pattern
	}

	//copy gradient
	if !reflect.DeepEqual(style.Fill.Gradient, &ml.GradientFill{}) {
		fill.Gradient = &ml.GradientFill{}
		*fill.Gradient = *style.Fill.Gradient
		copy(fill.Gradient.Stop, style.Fill.Gradient.Stop)
	}

	//if fill is actually empty, then nil it
	if *fill == (ml.Fill{}) {
		fill = nil
	}

	return
}

//private method used by to convert Info to ml.RichFont
func toRichFont(f *Info) *ml.RichFont {
	style := f.styleInfo

	//copy non-empty font
	if (*style.Font != ml.Font{} && *style.Font != ml.Font{Size: 0, Family: 0, Charset: 0}) {
		font := ml.RichFont(*style.Font)
		return &font
	}

	return nil
}
