package internal

import (
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"reflect"
	"strconv"
)

//formatConverter is helper type to hide some methods and don't pollute globals
type formatConverter byte

//FromFormat is helper function to convert format into ml objects
func FromFormat(f *format.StyleFormat) (*ml.Font, *ml.Fill, *ml.CellAlignment, *ml.NumberFormat, *ml.CellProtection, *ml.Border) {
	var fc formatConverter

	//convert font
	font := &ml.Font{
		Name:      sharedML.Property(f.Font.Name),
		Family:    sharedML.Property(f.Font.Family.String()),
		Bold:      sharedML.PropertyBool(f.Font.Bold),
		Italic:    sharedML.PropertyBool(f.Font.Italic),
		Strike:    sharedML.PropertyBool(f.Font.Strike),
		Shadow:    sharedML.PropertyBool(f.Font.Shadow),
		Condense:  sharedML.PropertyBool(f.Font.Condense),
		Extend:    sharedML.PropertyBool(f.Font.Extend),
		Color:     fc.convertColor(f.Font.Color),
		Underline: sharedML.Property(f.Font.Underline.String()),
		VAlign:    sharedML.Property(f.Font.VAlign.String()),
		Scheme:    sharedML.Property(f.Font.Scheme.String()),
	}

	if f.Font.Size > 0 {
		font.Size = sharedML.Property(strconv.FormatFloat(f.Font.Size, 'f', -1, 64))
	}

	//convert fill
	fill := &ml.Fill{
		Pattern: &ml.PatternFill{
			Type:       f.Fill.Type,
			Color:      fc.convertColor(f.Fill.Color),
			Background: fc.convertColor(f.Fill.Background),
		},
	}

	//convert border
	border := fc.convertBorder(f)

	//convert number format
	numFormat := &ml.NumberFormat{
		ID:   f.NumFormat.ID,
		Code: f.NumFormat.Code,
	}

	//convert protection
	protection := &ml.CellProtection{
		Hidden: f.Protection.Hidden,
		Locked: f.Protection.Locked,
	}

	if reflect.DeepEqual(protection, &ml.CellProtection{}) {
		protection = nil
	}

	//convert alignment
	alignment := &ml.CellAlignment{
		Horizontal:      f.Alignment.Horizontal,
		Vertical:        f.Alignment.Vertical,
		TextRotation:    f.Alignment.TextRotation,
		WrapText:        f.Alignment.WrapText,
		Indent:          f.Alignment.Indent,
		RelativeIndent:  f.Alignment.RelativeIndent,
		JustifyLastLine: f.Alignment.JustifyLastLine,
		ShrinkToFit:     f.Alignment.ShrinkToFit,
		ReadingOrder:    f.Alignment.ReadingOrder,
	}

	if reflect.DeepEqual(alignment, &ml.CellAlignment{}) {
		alignment = nil
	}

	return font, fill, alignment, numFormat, protection, border
}

func (fc *formatConverter) convertColor(rgb format.ARGB) *ml.Color {
	var color *ml.Color

	if len(rgb) > 0 {
		if index := rgb.ToIndex(); index != -1 {
			color = &ml.Color{Indexed: &index}
		} else {
			color = &ml.Color{RGB: rgb}
		}
	}

	return color
}

func (fc *formatConverter) convertBorder(f *format.StyleFormat) *ml.Border {
	var (
		borderTop    *ml.BorderSegment
		borderBottom *ml.BorderSegment
		borderLeft   *ml.BorderSegment
		borderRight  *ml.BorderSegment
	)

	borderColor := fc.convertColor(f.Border.Top.Color)
	borderType := f.Border.Top.Type
	if borderColor != nil || borderType > 0 {
		borderTop = &ml.BorderSegment{
			Color: borderColor,
			Type:  borderType,
		}
	}

	borderColor = fc.convertColor(f.Border.Bottom.Color)
	borderType = f.Border.Bottom.Type
	if borderColor != nil || borderType > 0 {
		borderBottom = &ml.BorderSegment{
			Color: borderColor,
			Type:  borderType,
		}
	}

	borderColor = fc.convertColor(f.Border.Left.Color)
	borderType = f.Border.Left.Type
	if borderColor != nil || borderType > 0 {
		borderLeft = &ml.BorderSegment{
			Color: borderColor,
			Type:  borderType,
		}
	}

	borderColor = fc.convertColor(f.Border.Right.Color)
	borderType = f.Border.Right.Type
	if borderColor != nil || borderType > 0 {
		borderRight = &ml.BorderSegment{
			Color: borderColor,
			Type:  borderType,
		}
	}

	return &ml.Border{
		Top:    borderTop,
		Bottom: borderBottom,
		Left:   borderLeft,
		Right:  borderRight,
	}
}
