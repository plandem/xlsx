package styles

import (
	"encoding/xml"
)

//BorderStyleType is a type to encode XSD BorderStyleType
type BorderStyleType byte

var (
	ToBorderStyleType   map[string]BorderStyleType
	FromBorderStyleType map[BorderStyleType]string
)

func (e BorderStyleType) String() string {
	return FromBorderStyleType[e]
}

func (e *BorderStyleType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromBorderStyleType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *BorderStyleType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToBorderStyleType[attr.Value]; ok {
		*e = v
	}

	return nil
}
