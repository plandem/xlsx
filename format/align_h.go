package format

import "encoding/xml"

//HAlignType is a type to encode XSD ST_HorizontalAlignment
type HAlignType byte

const (
	_ HAlignType = iota
	HAlignGeneral
	HAlignLeft
	HAlignCenter
	HAlignRight
	HAlignFill
	HAlignJustify
	HAlignCenterContinuous
	HAlignDistributed
)

func (e *HAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *e {
	case HAlignGeneral:
		attr.Value = "general"
	case HAlignLeft:
		attr.Value = "left"
	case HAlignCenter:
		attr.Value = "center"
	case HAlignRight:
		attr.Value = "right"
	case HAlignFill:
		attr.Value = "fill"
	case HAlignJustify:
		attr.Value = "justify"
	case HAlignCenterContinuous:
		attr.Value = "centerContinuous"
	case HAlignDistributed:
		attr.Value = "distributed"
	default:
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *HAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "general":
		*e = HAlignGeneral
	case "left":
		*e = HAlignLeft
	case "center":
		*e = HAlignCenter
	case "right":
		*e = HAlignRight
	case "fill":
		*e = HAlignFill
	case "justify":
		*e = HAlignJustify
	case "centerContinuous":
		*e = HAlignCenterContinuous
	case "distributed":
		*e = HAlignDistributed
	}

	return nil
}
