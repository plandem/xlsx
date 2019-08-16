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
	XMLName    xml.Name    `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing wsDr"`
	DMLName    ml.DMLName  `xml:",attr"`
	AnchorList *AnchorList `xml:",any"`
}

//ClientData is a direct mapping of XSD CT_AnchorClientData
type ClientData struct {
	LocksWithSheet  *bool `xml:"fLocksWithSheet,attr,omitempty"`
	PrintsWithSheet *bool `xml:"fPrintsWithSheet,attr,omitempty"`
}

type object struct {
	Shape        *ml.Reserved          `xml:"sp,omitempty"`
	Group        *ml.Reserved          `xml:"grpSp,omitempty"`
	GraphicFrame *GraphicalObjectFrame `xml:"graphicFrame,omitempty"`
	Connector    *ml.Reserved          `xml:"cxnSp,omitempty"`
	Picture      *ml.Reserved          `xml:"pic,omitempty"`
	Relation     *ml.Reserved          `xml:"contentPart,omitempty"`
}

//Marker is a direct mapping of XSD CT_Marker
type Marker struct {
	Col       int            `xml:"col"`
	OffsetCol dml.Coordinate `xml:"colOff"`
	Row       int            `xml:"row"`
	OffsetRow dml.Coordinate `xml:"rowOff"`
}

//AbsoluteAnchor is a direct mapping of XSD CT_AbsoluteAnchor
type AbsoluteAnchor struct {
	XMLName    xml.Name           `xml:"absoluteAnchor"`
	Point      dml.Point2D        `xml:"pos"`
	Size       dml.PositiveSize2D `xml:"ext"`
	ClientData ClientData         `xml:"clientData"`
	object
}

//OneCellAnchor is a direct mapping of XSD CT_OneCellAnchor
type OneCellAnchor struct {
	XMLName    xml.Name           `xml:"oneCellAnchor"`
	From       Marker             `xml:"from"`
	Size       dml.PositiveSize2D `xml:"ext"`
	ClientData ClientData         `xml:"clientData"`
	object
}

//TwoCellAnchor is a direct mapping of XSD CT_TwoCellAnchor
type TwoCellAnchor struct {
	XMLName    xml.Name   `xml:"twoCellAnchor"`
	From       Marker     `xml:"from"`
	To         Marker     `xml:"to"`
	EditAs     string     `xml:"editAs,attr,omitempty"` //enum
	ClientData ClientData `xml:"clientData"`
	object
}

//GraphicalObjectFrame is a direct mapping of XSD CT_GraphicalObjectFrame
type GraphicalObjectFrame struct {
	NonVisual *dml.GraphicalObjectFrameNonVisual `xml:"nvGraphicFramePr"`
	Graphic   *dml.GraphicalObject               `xml:"graphic"`
	Transform *dml.Transform2D                   `xml:"xfrm"`
	ml.ReservedAttributes
	ml.ReservedElements
}
