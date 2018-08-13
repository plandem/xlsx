package helpers

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"reflect"
	"unsafe"
)

//FromStyleFormat checks current style settings and returns copies of non-empty objects
func FromStyleFormat(f *format.StyleFormat) (font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, numFormat *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border, namedInfo *ml.NamedStyleInfo) {
	v := reflect.ValueOf(f).Elem()
	style := (*ml.DiffStyle)(unsafe.Pointer(v.FieldByName("styleInfo").Pointer()))
	named := (*ml.NamedStyleInfo)(unsafe.Pointer(v.FieldByName("namedInfo").Pointer()))

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

//ToStyleFormat creates a style object with copies of provided information
func ToStyleFormat(styleInfo *ml.DiffStyle, namedStyle *ml.DiffStyle, namedInfo *ml.NamedStyleInfo) *format.StyleFormat {
	f := format.New()
	v := reflect.ValueOf(f).Elem()

	si := (*ml.DiffStyle)(unsafe.Pointer(v.FieldByName("styleInfo").Pointer()))
	ni := (*ml.NamedStyleInfo)(unsafe.Pointer(v.FieldByName("namedInfo").Pointer()))

	//copy non-empty namedInfo
	if namedInfo != nil && namedInfo.BuiltinId != nil {
		*ni = *namedInfo
	}

	if styleInfo != nil {
		//TODO:
		panic("not implemented yet")
	}

	//set fields
	fs := v.FieldByName("styleInfo")
	fs = reflect.NewAt(fs.Type(), unsafe.Pointer(fs.UnsafeAddr())).Elem()
	fs.Set(reflect.ValueOf(si))

	fn := v.FieldByName("namedInfo")
	fn = reflect.NewAt(fn.Type(), unsafe.Pointer(fn.UnsafeAddr())).Elem()
	fn.Set(reflect.ValueOf(ni))

	return f
}
