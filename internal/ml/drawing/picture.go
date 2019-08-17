// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
)

//Picture is a direct mapping of XSD CT_Picture
type Picture struct {
	NonVisual *PictureNonVisual       `xml:"nvPicPr"`
	BlipFill  *dml.BlipFillProperties `xml:"blipFill"`
	Shape     *dml.ShapeProperties    `xml:"spPr"`
	Style     *dml.ShapeStyle         `xml:"style,omitempty"`
	ml.ReservedAttributes
}

//PictureNonVisual is a direct mapping of XSD CT_PictureNonVisual
type PictureNonVisual struct {
	CommonProperties  *dml.NonVisualCommonProperties  `xml:"cNvPr"`
	PictureProperties *dml.NonVisualPictureProperties `xml:"cNvPicPr"`
}
