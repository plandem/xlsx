package ml

import (
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//SharedStrings is a direct mapping of XSD CT_Sst
type SharedStrings struct {
	XMLName     ml.Name       `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main sst"`
	Count       uint          `xml:"count,attr,omitempty"`
	UniqueCount uint          `xml:"uniqueCount,attr,omitempty"`
	StringItem  []*StringItem `xml:"si"`
	ExtLst      *ml.Reserved  `xml:"extLst,omitempty"`
}

//StringItem is a direct mapping of XSD CT_Rst
type StringItem struct {
	Text       primitives.Text `xml:"t,omitempty"`
	R          *ml.Reserved    `xml:"r,omitempty"`
	RPh        *ml.Reserved    `xml:"rPh,omitempty"`
	PhoneticPr *ml.Reserved    `xml:"phoneticPr,omitempty"`
}
