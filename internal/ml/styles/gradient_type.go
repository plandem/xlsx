package styles

import (
	"encoding/xml"
)

//GradientType is a type to encode XSD ST_GradientType
type GradientType byte

var (
	ToGradientType  map[string]GradientType
	FromGradientType map[GradientType]string
)

func (t GradientType) String() string {
	return FromGradientType[t]
}

func (t *GradientType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromGradientType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (t *GradientType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToGradientType[attr.Value]; ok {
		*t = v
	}

	return nil
}
