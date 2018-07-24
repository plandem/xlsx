package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//Border returns string with all values of border
func Border(border *ml.Border) string {
	var b ml.Border

	if border == nil {
		b = ml.Border{}
	} else {
		//we don't want to mutate original border
		b = *border
	}

	if b.Left == nil {
		b.Left = &ml.BorderSegment{}
	}

	if b.Right == nil {
		b.Right = &ml.BorderSegment{}
	}

	if b.Top == nil {
		b.Top = &ml.BorderSegment{}
	}

	if b.Bottom == nil {
		b.Bottom = &ml.BorderSegment{}
	}

	if b.Diagonal == nil {
		b.Diagonal = &ml.BorderSegment{}
	}

	if b.Vertical == nil {
		b.Vertical = &ml.BorderSegment{}
	}

	if b.Horizontal == nil {
		b.Horizontal = &ml.BorderSegment{}
	}

	return strings.Join([]string{
		Color(b.Left.Color),
		b.Left.Type.String(),

		Color(b.Right.Color),
		b.Right.Type.String(),

		Color(b.Top.Color),
		b.Top.Type.String(),

		Color(b.Bottom.Color),
		b.Bottom.Type.String(),

		Color(b.Diagonal.Color),
		b.Diagonal.Type.String(),

		Color(b.Vertical.Color),
		b.Vertical.Type.String(),

		Color(b.Horizontal.Color),
		b.Horizontal.Type.String(),

		strconv.FormatBool(bool(b.DiagonalUp)),
		strconv.FormatBool(bool(b.DiagonalDown)),
		strconv.FormatBool(bool(b.Outline)),
	}, ":")
}
