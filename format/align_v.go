package format

import (
	"encoding/xml"
)

//VAlignType is a type to encode XSD ST_VerticalAlignment
type VAlignType byte

//List of all possible values for VAlignType
const (
	_ VAlignType = iota
	VAlignTop
	VAlignCenter
	VAlignBottom
	VAlignJustify
	VAlignDistributed
)

var (
	toVAlignType   map[string]VAlignType
	fromVAlignType map[VAlignType]string
)

func init() {
	fromVAlignType = map[VAlignType]string{
		VAlignTop:         "top",
		VAlignCenter:      "center",
		VAlignBottom:      "bottom",
		VAlignJustify:     "justify",
		VAlignDistributed: "distributed",
	}

	toVAlignType = make(map[string]VAlignType, len(fromVAlignType))
	for k, v := range fromVAlignType {
		toVAlignType[v] = k
	}
}

func (e VAlignType) String() string {
	return fromVAlignType[e]
}

func (e *VAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromVAlignType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *VAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toVAlignType[attr.Value]; ok {
		*e = v
	}

	return nil
}
