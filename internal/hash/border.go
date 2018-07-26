package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//Border returns string with all values of border
func Border(border *ml.Border) Key {
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

	return Key(strings.Join([]string{
		string(Color(b.Left.Color)),
		b.Left.Type.String(),

		string(Color(b.Right.Color)),
		b.Right.Type.String(),

		string(Color(b.Top.Color)),
		b.Top.Type.String(),

		string(Color(b.Bottom.Color)),
		b.Bottom.Type.String(),

		string(Color(b.Diagonal.Color)),
		b.Diagonal.Type.String(),

		string(Color(b.Vertical.Color)),
		b.Vertical.Type.String(),

		string(Color(b.Horizontal.Color)),
		b.Horizontal.Type.String(),

		strconv.FormatBool(bool(b.DiagonalUp)),
		strconv.FormatBool(bool(b.DiagonalDown)),
		strconv.FormatBool(bool(b.Outline)),
	}, ":"))
}
