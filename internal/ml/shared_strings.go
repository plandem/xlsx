// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//SharedStrings is a direct mapping of XSD CT_Sst
type SharedStrings struct {
	XMLName     xml.Name      `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main sst"`
	Count       uint          `xml:"count,attr,omitempty"`
	UniqueCount uint          `xml:"uniqueCount,attr,omitempty"`
	StringItem  []*StringItem `xml:"si"`
	ExtLst      *ml.Reserved  `xml:"extLst,omitempty"`
}

//StringItem is a direct mapping of XSD CT_Rst
type StringItem struct {
	Text       primitives.Text `xml:"t,omitempty"`   //optional
	RichText   []*RichText     `xml:"r,omitempty"`   //optional
	RPh        []*ml.Reserved  `xml:"rPh,omitempty"` //optional
	PhoneticPr *ml.Reserved    `xml:"phoneticPr,omitempty"`
}

//RichText is a direct mapping of XSD CT_RElt
type RichText struct {
	Font *RichFont       `xml:"rPr,omitempty"`
	Text primitives.Text `xml:"t"` //required
}

//RichFont is a direct mapping of XSD CT_RPrElt
//N.B.: it's weird, but CT_RPrElt is clone CT_Font, except different tag for 'Name' field
type RichFont struct {
	Name      ml.Property                `xml:"rFont,omitempty"`
	Charset   primitives.FontCharsetType `xml:"charset,omitempty"`
	Family    primitives.FontFamilyType  `xml:"family,omitempty"`
	Bold      ml.PropertyBool            `xml:"b,omitempty"`
	Italic    ml.PropertyBool            `xml:"i,omitempty"`
	Strike    ml.PropertyBool            `xml:"strike,omitempty"`
	Shadow    ml.PropertyBool            `xml:"shadow,omitempty"`
	Condense  ml.PropertyBool            `xml:"condense,omitempty"`
	Extend    ml.PropertyBool            `xml:"extend,omitempty"`
	Color     *Color                     `xml:"color,omitempty"`
	Size      ml.PropertyDouble          `xml:"sz,omitempty"`
	Underline primitives.UnderlineType   `xml:"u,omitempty"`
	VAlign    primitives.FontVAlignType  `xml:"vertAlign,omitempty"`
	Scheme    primitives.FontSchemeType  `xml:"scheme,omitempty"`
}
