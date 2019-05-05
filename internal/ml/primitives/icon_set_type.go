package primitives

import (
	"encoding/xml"
)

//IconSetType is a direct mapping of XSD ST_CfvoType
type IconSetType byte

//IconSetType maps for marshal/unmarshal process
var (
	ToIconSetType   map[string]IconSetType
	FromIconSetType map[IconSetType]string
)

func (t IconSetType) String() string {
	return FromIconSetType[t]
}

//MarshalXMLAttr marshal IconSetType
func (t *IconSetType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromIconSetType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal IconSetType
func (t *IconSetType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToIconSetType[attr.Value]; ok {
		*t = v
	}

	return nil
}
