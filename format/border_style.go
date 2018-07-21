package format

import (
	"encoding/xml"
)

//BorderStyleType is a type to encode XSD BorderStyleType
type BorderStyleType byte

//List of all possible values for BorderStyleType
const (
	_ BorderStyleType = iota
	BorderStyleNone
	BorderStyleThin
	BorderStyleMedium
	BorderStyleDashed
	BorderStyleDotted
	BorderStyleThick
	BorderStyleDouble
	BorderStyleHair
	BorderStyleMediumDashed
	BorderStyleDashDot
	BorderStyleMediumDashDot
	BorderStyleDashDotDot
	BorderStyleMediumDashDotDot
	BorderStyleSlantDashDot
)

var (
	toBorderStyleType   map[string]BorderStyleType
	fromBorderStyleType map[BorderStyleType]string
)

func init() {
	fromBorderStyleType = map[BorderStyleType]string{
		BorderStyleNone:             "none",
		BorderStyleThin:             "thin",
		BorderStyleMedium:           "medium",
		BorderStyleDashed:           "dashed",
		BorderStyleDotted:           "dotted",
		BorderStyleThick:            "thick",
		BorderStyleDouble:           "double",
		BorderStyleHair:             "hair",
		BorderStyleMediumDashed:     "mediumDashed",
		BorderStyleDashDot:          "dashDot",
		BorderStyleMediumDashDot:    "mediumDashDot",
		BorderStyleDashDotDot:       "dashDotDot",
		BorderStyleMediumDashDotDot: "mediumDashDotDot",
		BorderStyleSlantDashDot:     "slantDashDot",
	}

	toBorderStyleType = make(map[string]BorderStyleType, len(fromBorderStyleType))
	for k, v := range fromBorderStyleType {
		toBorderStyleType[v] = k
	}
}

func (e BorderStyleType) String() string {
	return fromBorderStyleType[e]
}

func (e *BorderStyleType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromBorderStyleType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *BorderStyleType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toBorderStyleType[attr.Value]; ok {
		*e = v
	}

	return nil
}
