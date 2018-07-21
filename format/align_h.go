package format

import (
	"encoding/xml"
)

//HAlignType is a type to encode XSD ST_HorizontalAlignment
type HAlignType byte

//List of all possible values for HAlignType
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

var (
	toHAlignType   map[string]HAlignType
	fromHAlignType map[HAlignType]string
)

func init() {
	fromHAlignType = map[HAlignType]string{
		HAlignGeneral:          "general",
		HAlignLeft:             "left",
		HAlignCenter:           "center",
		HAlignRight:            "right",
		HAlignFill:             "fill",
		HAlignJustify:          "justify",
		HAlignCenterContinuous: "centerContinuous",
		HAlignDistributed:      "distributed",
	}

	toHAlignType = make(map[string]HAlignType, len(fromHAlignType))
	for k, v := range fromHAlignType {
		toHAlignType[v] = k
	}
}

func (e *HAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromHAlignType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *HAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toHAlignType[attr.Value]; ok {
		*e = v
	}

	return nil
}
