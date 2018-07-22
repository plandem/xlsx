package ml

import (
	"encoding/xml"
)

//VAlignType is a type to encode XSD ST_VerticalAlignment
type VAlignType byte

var (
	ToVAlignType   map[string]VAlignType
	FromVAlignType map[VAlignType]string
)

func (e VAlignType) String() string {
	return FromVAlignType[e]
}

func (e *VAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromVAlignType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *VAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToVAlignType[attr.Value]; ok {
		*e = v
	}

	return nil
}
