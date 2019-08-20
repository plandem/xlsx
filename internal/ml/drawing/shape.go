// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
)

//Shape is a direct mapping of XSD CT_Shape
type Shape struct {
	XMLName   xml.Name        `xml:"sp"`
	NonVisual *ShapeNonVisual `xml:"nvSpPr"`
	Shape     *dml.Shape      `xml:"spPr"`
	Style     *dml.ShapeStyle `xml:"style,omitempty"`
	Text      *dml.TextBody   `xml:"txBody,omitempty"`
	ml.ReservedAttributes
}

//ShapeNonVisual is a direct mapping of XSD CT_ShapeNonVisual
type ShapeNonVisual struct {
	CommonProperties *dml.NonVisualCommonProperties `xml:"cNvPr"`
	ShapeProperties  *dml.NonVisualShapeProperties  `xml:"cNvSpPr"`
}
