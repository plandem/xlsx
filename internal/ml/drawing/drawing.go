// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
)

//Drawing is a direct mapping of XSD CT_Drawing
type Drawing struct {
	XMLName xml.Name    `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing wsDr"`
	DMLName dml.Name    `xml:",attr"`
	Items   *AnchorList `xml:",any"`
}

//ClientData is a direct mapping of XSD CT_AnchorClientData
type ClientData struct {
	LocksWithSheet  *bool `xml:"fLocksWithSheet,attr,omitempty"`
	PrintsWithSheet *bool `xml:"fPrintsWithSheet,attr,omitempty"`
}

//Ref is a direct mapping of XSD CT_Rel
type Ref struct {
	RIDName ml.RIDName `xml:",attr"`
	RID     ml.RID     `xml:"id,attr"`
}

//Marker is a direct mapping of XSD CT_Marker
type Marker struct {
	Col       int            `xml:"col"`
	OffsetCol dml.Coordinate `xml:"colOff"`
	Row       int            `xml:"row"`
	OffsetRow dml.Coordinate `xml:"rowOff"`
}

type object struct {
	Shape     *Shape     `xml:"sp,omitempty"`
	Group     *Group     `xml:"grpSp,omitempty"`
	Frame     *Frame     `xml:"graphicFrame,omitempty"`
	Connector *Connector `xml:"cxnSp,omitempty"`
	Picture   *Picture   `xml:"pic,omitempty"`
	Relation  *Ref       `xml:"contentPart,omitempty"`
}
