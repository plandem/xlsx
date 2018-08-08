package primitives

import (
	"encoding/xml"
)

//GradientType is a type to encode XSD ST_GradientType
type GradientType byte

//GradientType maps for marshal/unmarshal process
var (
	ToGradientType   map[string]GradientType
	FromGradientType map[GradientType]string
)

func (t GradientType) String() string {
	return FromGradientType[t]
}

//MarshalXMLAttr marshal GradientType
func (t *GradientType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromGradientType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal GradientType
func (t *GradientType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToGradientType[attr.Value]; ok {
		*t = v
	}

	return nil
}
