// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//GUID is a direct mapping of XSD ST_Guid
type GUID string

//Comments is a direct mapping of XSD CT_Comments
type Comments struct {
	XMLName     xml.Name          `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main comments"`
	Authors     []primitives.Text `xml:"authors>author"`
	CommentList []*Comment        `xml:"commentList>comment"`
	ExtLst      *ml.Reserved      `xml:"extLst,omitempty"`
}

//Comment is a direct mapping of XSD CT_Comment
type Comment struct {
	Text     *StringItem       `xml:"text"`
	Options  *CommentOptions   `xml:"commentPr,omitempty"`
	Ref      primitives.Bounds `xml:"ref,attr"`
	AuthorID int               `xml:"authorId,attr"`
	Guid     GUID              `xml:"guid,attr,omitempty"`
	ShapeID  ml.OptionalIndex  `xml:"shapeId,attr,omitempty"`
}

//CommentOptions is a direct mapping of CT_CommentPr
type CommentOptions struct {
	Anchor      *ml.Reserved `xml:"anchor"`
	Locked      *bool        `xml:"locked,attr,omitempty"`
	DefaultSize *bool        `xml:"defaultSize,attr,omitempty"`
	Print       *bool        `xml:"print,attr,omitempty"`
	Disabled    bool         `xml:"disabled,attr,omitempty"`
	JustLastX   bool         `xml:"justLastX,attr,omitempty"`
	AutoScale   bool         `xml:"autoScale,attr,omitempty"`
	AutoFill    *bool        `xml:"autoFill,attr,omitempty"`
	AutoLine    *bool        `xml:"autoLine,attr,omitempty"`
	AltText     string       `xml:"altText,attr,omitempty"`
	TextHAlign  *ml.Reserved `xml:"textHAlign,attr,omitempty"`
	TextVAlign  *ml.Reserved `xml:"textVAlign,attr,omitempty"`
	LockText    *bool        `xml:"lockText,attr,omitempty"`
}

//ClientData is direct mapping for CT_ClientData, that used for Excel specific settings of Shape
type ClientData struct {
	XMLName       xml.Name `xml:"x:ClientData"`
	MoveWithCells bool     `xml:"x:MoveWithCells,omitempty"`
	SizeWithCells bool     `xml:"x:SizeWithCells,omitempty"`
	AutoFill      bool     `xml:"x:AutoFill,omitempty"`
	Type          string   `xml:"ObjectType,attr"`
	Anchor        string   `xml:"x:Anchor,omitempty"`
	Row           int      `xml:"x:Row"`
	Column        int      `xml:"x:Column"`
}
