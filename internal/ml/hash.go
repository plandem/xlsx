// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml

import (
	"github.com/plandem/ooxml/index"
	"strconv"
	"strings"
)

//Hash builds hash code for all required values of CellAlignment to use as unique index
func (a *CellAlignment) Hash() index.Code {
	alignment := a
	if alignment == nil {
		alignment = &CellAlignment{}
	}

	return index.Hash(strings.Join([]string{
		strconv.FormatInt(int64(alignment.Horizontal), 10),
		strconv.FormatInt(int64(alignment.Vertical), 10),
		strconv.FormatInt(int64(alignment.TextRotation), 10),
		strconv.FormatBool(alignment.WrapText),
		strconv.FormatInt(int64(alignment.Indent), 10),
		strconv.FormatInt(int64(alignment.RelativeIndent), 10),
		strconv.FormatBool(alignment.JustifyLastLine),
		strconv.FormatBool(alignment.ShrinkToFit),
		strconv.FormatInt(int64(alignment.ReadingOrder), 10),
	}, ":"))
}

//Hash builds hash code for all required values of Color to use as unique index
func (c *Color) Hash() index.Code {
	color := c
	if color == nil {
		color = &Color{}
	}

	result := []string{
		strconv.FormatBool(color.Auto),
		color.RGB,
		strconv.FormatFloat(color.Tint, 'f', -1, 64),
	}

	if color.Indexed != nil {
		result = append(result, strconv.FormatInt(int64(*color.Indexed), 10))
	} else {
		result = append(result, "")
	}

	if color.Theme != nil {
		result = append(result, strconv.FormatInt(int64(*color.Theme), 10))
	} else {
		result = append(result, "")
	}

	return index.Hash(strings.Join(result, ":"))
}

//Hash builds hash code for all required values of Border to use as unique index
func (border *Border) Hash() index.Code {
	var b Border

	if border == nil {
		b = Border{}
	} else {
		//we don't want to mutate original border
		b = *border
	}

	if b.Left == nil {
		b.Left = &BorderSegment{}
	}

	if b.Right == nil {
		b.Right = &BorderSegment{}
	}

	if b.Top == nil {
		b.Top = &BorderSegment{}
	}

	if b.Bottom == nil {
		b.Bottom = &BorderSegment{}
	}

	if b.Diagonal == nil {
		b.Diagonal = &BorderSegment{}
	}

	if b.Vertical == nil {
		b.Vertical = &BorderSegment{}
	}

	if b.Horizontal == nil {
		b.Horizontal = &BorderSegment{}
	}

	return index.Hash(strings.Join([]string{
		b.Left.Color.Hash().String(),
		b.Left.Type.String(),

		b.Right.Color.Hash().String(),
		b.Right.Type.String(),

		b.Top.Color.Hash().String(),
		b.Top.Type.String(),

		b.Bottom.Color.Hash().String(),
		b.Bottom.Type.String(),

		b.Diagonal.Color.Hash().String(),
		b.Diagonal.Type.String(),

		b.Vertical.Color.Hash().String(),
		b.Vertical.Type.String(),

		b.Horizontal.Color.Hash().String(),
		b.Horizontal.Type.String(),

		strconv.FormatBool(bool(b.DiagonalUp)),
		strconv.FormatBool(bool(b.DiagonalDown)),
		strconv.FormatBool(bool(b.Outline)),
	}, ":"))
}

//Hash builds hash code for all required values of Fill to use as unique index
func (fill *Fill) Hash() index.Code {
	var f Fill

	if fill == nil {
		f = Fill{}
	} else {
		//we don't want to mutate original fill
		f = *fill
	}

	if f.Pattern == nil {
		f.Pattern = &PatternFill{}
	}

	if f.Gradient == nil {
		f.Gradient = &GradientFill{}
	}

	result := []string{
		strconv.FormatInt(int64(f.Pattern.Type), 10),
		f.Pattern.Color.Hash().String(),
		f.Pattern.Background.Hash().String(),
		strconv.FormatInt(int64(f.Gradient.Type), 10),
		strconv.FormatFloat(float64(f.Gradient.Degree), 'f', -1, 64),
		strconv.FormatFloat(float64(f.Gradient.Left), 'f', -1, 64),
		strconv.FormatFloat(float64(f.Gradient.Right), 'f', -1, 64),
		strconv.FormatFloat(float64(f.Gradient.Top), 'f', -1, 64),
		strconv.FormatFloat(float64(f.Gradient.Bottom), 'f', -1, 64),
	}

	for _, stop := range f.Gradient.Stop {
		result = append(result,
			strconv.FormatFloat(float64(stop.Position), 'f', -1, 64),
			stop.Color.Hash().String(),
		)
	}

	return index.Hash(strings.Join(result, ":"))
}

//Hash builds hash code for all required values of Font to use as unique index
func (f *Font) Hash() index.Code {
	font := f
	if font == nil {
		font = &Font{}
	}

	return index.Hash(strings.Join([]string{
		string(font.Name),
		strconv.FormatInt(int64(font.Charset), 10),
		strconv.FormatInt(int64(font.Family), 10),
		strconv.FormatBool(bool(font.Bold)),
		strconv.FormatBool(bool(font.Italic)),
		strconv.FormatBool(bool(font.Strike)),
		strconv.FormatBool(bool(font.Shadow)),
		strconv.FormatBool(bool(font.Condense)),
		strconv.FormatBool(bool(font.Extend)),
		font.Color.Hash().String(),
		strconv.FormatFloat(float64(font.Size), 'f', -1, 64),
		string(font.Underline),
		string(font.VAlign),
		string(font.Scheme),
	}, ":"))
}

//Hash builds hash code for all required values of NumberFormat to use as unique index
func (f *NumberFormat) Hash() index.Code {
	format := f
	if format == nil {
		format = &NumberFormat{}
	}

	return index.Hash(strings.Join([]string{
		strconv.FormatInt(int64(format.ID), 10),
		format.Code,
	}, ":"))
}

//Hash builds hash code for all required values of CellProtection to use as unique index
func (p *CellProtection) Hash() index.Code {
	protection := p
	if protection == nil {
		protection = &CellProtection{}
	}

	return index.Hash(strings.Join([]string{
		strconv.FormatBool(protection.Locked),
		strconv.FormatBool(protection.Hidden),
	}, ":"))
}

//Hash builds hash code for all required values of StringItem to use as unique index
func (s *StringItem) Hash() index.Code {
	si := s
	if si == nil {
		si = &StringItem{}
	}

	result := []string{
		string(si.Text),
		string(si.PhoneticPr.Hash()),
	}

	if si.RPh != nil {
		for _, r := range *si.RPh {
			result = append(result, r.Hash().String())
		}
	}

	if si.RichText != nil {
		for _, part := range *si.RichText {
			result = append(result, string(part.Text))

			if part.Font != nil {
				font := Font(*part.Font)
				result = append(result, font.Hash().String())
			}
		}
	}

	return index.Hash(strings.Join(result, ":"))
}

//Hash builds hash code for all required values of NamedStyle to use as unique index
func (ns *NamedStyle) Hash() index.Code {
	style := ns
	if style == nil {
		style = &NamedStyle{}
	}

	return index.Hash(strings.Join([]string{
		strconv.FormatInt(int64(style.NumFmtId), 10),
		strconv.FormatInt(int64(style.FontId), 10),
		strconv.FormatInt(int64(style.FillId), 10),
		strconv.FormatInt(int64(style.BorderId), 10),
		strconv.FormatBool(style.QuotePrefix),
		strconv.FormatBool(style.PivotButton),
		strconv.FormatBool(style.ApplyNumberFormat),
		strconv.FormatBool(style.ApplyFont),
		strconv.FormatBool(style.ApplyFill),
		strconv.FormatBool(style.ApplyBorder),
		strconv.FormatBool(style.ApplyAlignment),
		strconv.FormatBool(style.ApplyProtection),
		style.Alignment.Hash().String(),
		style.Protection.Hash().String(),
		style.ExtLst.Hash().String(),
	}, ":"))
}

//Hash builds hash code for all required values of DirectStyle to use as unique index
func (ds *DirectStyle) Hash() index.Code {
	style := ds
	if style == nil {
		style = &DirectStyle{}
	}

	s := NamedStyle(style.Style)
	return index.Hash(s.Hash().String() + ":" + strconv.FormatInt(int64(style.XfId), 10))
}

//Hash builds hash code for all required values of DiffStyle to use as unique index
func (ds *DiffStyle) Hash() index.Code {
	style := ds
	if style == nil {
		style = &DiffStyle{}
	}

	return index.Hash(strings.Join([]string{
		style.Border.Hash().String(),
		style.Fill.Hash().String(),
		style.Font.Hash().String(),
		style.NumberFormat.Hash().String(),
		style.Alignment.Hash().String(),
		style.Protection.Hash().String(),
		style.ExtLst.Hash().String(),
	}, ":"))
}