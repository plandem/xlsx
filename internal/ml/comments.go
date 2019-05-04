package ml

import (
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//AuthorID is helper alias type for authorId
type AuthorID uint

//GUID is a direct mapping of XSD ST_Guid
type GUID string

//Comments is a direct mapping of XSD CT_Comments
type Comments struct {
	XMLName     ml.Name           `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main comments"`
	Authors     []primitives.Text `xml:"authors>author"`
	CommentList []*Comment        `xml:"commentList>comment"`
	ExtLst      *ml.Reserved      `xml:"extLst,omitempty"`
}

//Comment is a direct mapping of XSD CT_Comment
type Comment struct {
	Text      *StringItem       `xml:"text"`
	CommentPr *ml.Reserved      `xml:"commentPr,omitempty"`
	Ref       primitives.Bounds `xml:"ref,attr"`
	AuthorID  AuthorID          `xml:"authorId,attr"`
	Guid      GUID              `xml:"guid,attr,omitempty"`
	ShapeID   ml.OptionalIndex  `xml:"shapeId,attr,omitempty"`
}
