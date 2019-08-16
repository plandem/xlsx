// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import "encoding/xml"

//TwoCellAnchor is a direct mapping of XSD CT_TwoCellAnchor
type TwoCellAnchor struct {
	XMLName    xml.Name   `xml:"twoCellAnchor"`
	From       Marker     `xml:"from"`
	To         Marker     `xml:"to"`
	EditAs     string     `xml:"editAs,attr,omitempty"` //enum
	ClientData ClientData `xml:"clientData"`
	object
}
