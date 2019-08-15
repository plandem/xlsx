// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"encoding/xml"
	drawingML "github.com/plandem/ooxml/drawing/ml"
	"github.com/plandem/ooxml/ml"
)

//AnchorList is special type to hold all anchors with preserving order
type AnchorList []interface{}

//Drawing is a direct mapping of XSD CT_Drawing
type Drawing struct {
	XMLName    xml.Name    `xml:"wsDr"`
	AnchorList *AnchorList `xml:",any"`
}

//ClientData is a direct mapping of XSD ST_Coordinate
type Coordinate string //ST_CoordinateUnqualified s:ST_UniversalMeasure

//ClientData is a direct mapping of XSD CT_AnchorClientData
type ClientData struct {
	LocksWithSheet  *bool `xml:"fLocksWithSheet,attr,omitempty"`
	PrintsWithSheet *bool `xml:"fPrintsWithSheet,attr,omitempty"`
}

type object struct {
	Shape        *ml.Reserved  `xml:"sp,omitempty"`
	Group        *ml.Reserved  `xml:"grpSp,omitempty"`
	GraphicFrame *GraphicFrame `xml:"graphicFrame,omitempty"`
	Connector    *ml.Reserved  `xml:"cxnSp,omitempty"`
	Picture      *ml.Reserved  `xml:"pic,omitempty"`
	Relation     *ml.Reserved  `xml:"contentPart,omitempty"`
}

//dedicated type for custom marshaller to apply namespace
type id int

//Marker is a direct mapping of XSD CT_Marker
type Marker struct {
	Col       id         `xml:"col"`
	OffsetCol Coordinate `xml:"colOff"`
	Row       id         `xml:"row"`
	OffsetRow Coordinate `xml:"rowOff"`
}

//Point is a direct mapping of XSD CT_Point2D
type Point struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}

//Size is a direct mapping of XSD CT_PositiveSize2D
type Size struct {
	Width  uint `xml:"cx,attr"`
	Height uint `xml:"cy,attr"`
}

//AbsoluteAnchor is a direct mapping of XSD CT_AbsoluteAnchor
type AbsoluteAnchor struct {
	Point      Point      `xml:"pos"`
	Size       Size       `xml:"ext"`
	ClientData ClientData `xml:"clientData"`
	object
}

//OneCellAnchor is a direct mapping of XSD CT_OneCellAnchor
type OneCellAnchor struct {
	From       Marker     `xml:"from"`
	Size       Size       `xml:"ext"`
	ClientData ClientData `xml:"clientData"`
	object
}

//TwoCellAnchor is a direct mapping of XSD CT_TwoCellAnchor
type TwoCellAnchor struct {
	From       Marker     `xml:"from"`
	To         Marker     `xml:"to"`
	EditAs     string     `xml:"editAs,attr,omitempty"` //enum
	ClientData ClientData `xml:"clientData"`
	object
}

//GraphicFrame is a direct mapping of XSD CT_GraphicalObjectFrame
type GraphicFrame struct {
	Graphic *drawingML.Graphic `xml:"graphic"`
	ml.ReservedAttributes
	ml.ReservedElements
}

func (a *AnchorList) Add(anchor interface{}) {
	*a = append(*a, anchor)
}
