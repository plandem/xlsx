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

func (e *VAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *e {
	case VAlignTop:
		attr.Value = "top"
	case VAlignCenter:
		attr.Value = "center"
	case VAlignBottom:
		attr.Value = "bottom"
	case VAlignJustify:
		attr.Value = "justify"
	case VAlignDistributed:
		attr.Value = "distributed"
	default:
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *VAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "top":
		*e = VAlignTop
	case "center":
		*e = VAlignCenter
	case "bottom":
		*e = VAlignBottom
	case "justify":
		*e = VAlignJustify
	case "distributed":
		*e = VAlignDistributed
	}

	return nil
}
