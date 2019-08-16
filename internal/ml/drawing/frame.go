// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
)

//Frame is a direct mapping of XSD CT_GraphicalObjectFrame
type Frame struct {
	NonVisual *FrameNonVisual      `xml:"nvGraphicFramePr"`
	Transform *dml.Transform2D     `xml:"xfrm"`
	Graphic   *dml.GraphicalObject `xml:"graphic"`
	ml.ReservedAttributes
}

//FrameNonVisual is a direct mapping of XSD CT_GraphicalObjectFrameNonVisual
type FrameNonVisual struct {
	DrawingProperties *dml.NonVisualCommonProperties `xml:"cNvPr"`
	FrameProperties   *dml.NonVisualFrameProperties  `xml:"cNvGraphicFramePr"`
}
