package ml

import (
	"encoding/xml"
)

//HAlignType is a type to encode XSD ST_HorizontalAlignment
type HAlignType byte

var (
	ToHAlignType   map[string]HAlignType
	FromHAlignType map[HAlignType]string
)

func (e HAlignType) String() string {
	return FromHAlignType[e]
}

func (e *HAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromHAlignType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *HAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToHAlignType[attr.Value]; ok {
		*e = v
	}

	return nil
}
