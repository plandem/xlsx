// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
)

//GraphicalObjectFrame is a direct mapping of XSD CT_GraphicalObjectFrame
type GraphicalObjectFrame struct {
	NonVisual *GraphicalObjectFrameNonVisual `xml:"nvGraphicFramePr"`
	Transform *dml.Transform2D               `xml:"xfrm"`
	Graphic   *dml.GraphicalObject           `xml:"graphic"`
	ml.ReservedAttributes
	ml.ReservedElements
}
