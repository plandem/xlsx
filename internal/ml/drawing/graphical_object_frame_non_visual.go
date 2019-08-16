// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"github.com/plandem/ooxml/drawing/dml"
)

//GraphicalObjectFrameNonVisual is a direct mapping of XSD CT_GraphicalObjectFrameNonVisual
type GraphicalObjectFrameNonVisual struct {
	DrawingProperties *dml.NonVisualDrawingProperties      `xml:"cNvPr"`
	FrameProperties   *dml.NonVisualGraphicFrameProperties `xml:"cNvGraphicFramePr"`
}
