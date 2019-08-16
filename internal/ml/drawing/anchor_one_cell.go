// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
)

//OneCellAnchor is a direct mapping of XSD CT_OneCellAnchor
type OneCellAnchor struct {
	XMLName    xml.Name           `xml:"oneCellAnchor"`
	From       Marker             `xml:"from"`
	Size       dml.PositiveSize2D `xml:"ext"`
	ClientData ClientData         `xml:"clientData"`
	object
}
