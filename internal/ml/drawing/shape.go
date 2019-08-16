// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
)

//Shape is a direct mapping of XSD CT_Shape
type Shape struct {
	NonVisual *ShapeNonVisual      `xml:"nvSpPr"`
	Shape     *dml.ShapeProperties `xml:"spPr"`
	Style     *dml.ShapeStyle      `xml:"style,omitempty"`
	Text      *dml.TextBody        `xml:"txBody,omitempty"`
	ml.ReservedAttributes
}

//ShapeNonVisual is a direct mapping of XSD CT_ShapeNonVisual
type ShapeNonVisual struct {
	DrawingProperties *dml.NonVisualCommonProperties `xml:"cNvPr"`
	ShapeProperties   *dml.NonVisualShapeProperties  `xml:"cNvSpPr"`
}
