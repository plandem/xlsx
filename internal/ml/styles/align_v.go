package styles

import (
	"encoding/xml"
)

//VAlignType is a type to encode XSD ST_VerticalAlignment
type VAlignType byte

var (
	ToVAlignType   map[string]VAlignType
	FromVAlignType map[VAlignType]string
)

func (t VAlignType) String() string {
	return FromVAlignType[t]
}

func (t *VAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromVAlignType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (t *VAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToVAlignType[attr.Value]; ok {
		*t = v
	}

	return nil
}
