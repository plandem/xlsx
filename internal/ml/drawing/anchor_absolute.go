// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
)

//AbsoluteAnchor is a direct mapping of XSD CT_AbsoluteAnchor
type AbsoluteAnchor struct {
	XMLName    xml.Name           `xml:"absoluteAnchor"`
	Point      dml.Point2D        `xml:"pos"`
	Size       dml.PositiveSize2D `xml:"ext"`
	ClientData ClientData         `xml:"clientData"`
	object
}
