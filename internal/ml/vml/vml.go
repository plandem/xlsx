package vml

import (
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/ooxml/vml/css"
)

//ShapeType is identical to the Shape element except it cannot reference another ShapeType element
type ShapeType = Shape

//Shape is direct mapping of CT_Shape, that used for VML drawing
type Shape struct {
	ID        string     `xml:"id,attr"`
	Type      string     `xml:"type,attr,omitempty"`
	Style     css.Style  `xml:"style,attr,omitempty"`
	FillColor string     `xml:"fillcolor,attr,omitempty"`
	Filled    bool       `xml:"filled,attr,omitempty"`
	Data      ClientData `xml:"ClientData,omitempty"`
	ml.Reserved
}

//ClientData is direct mapping for CT_ClientData, that used for Excel specific settings of Shape
type ClientData struct {
	Type          string `xml:"ObjectType,attr"`
	MoveWithCells bool   `xml:"MoveWithCells,omitempty"`
	SizeWithCells bool   `xml:"SizeWithCells,omitempty"`
	Anchor        string `xml:"Anchor,omitempty"`
	AutoFill      bool   `xml:"AutoFill,omitempty"`
	Row           int    `xml:"Row"`
	Column        int    `xml:"Column"`
}
