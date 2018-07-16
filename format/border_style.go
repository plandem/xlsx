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

func (e *BorderStyleType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *e {
	case BorderStyleNone:
		attr.Value = "none"
	case BorderStyleThin:
		attr.Value = "thin"
	case BorderStyleMedium:
		attr.Value = "medium"
	case BorderStyleDashed:
		attr.Value = "dashed"
	case BorderStyleDotted:
		attr.Value = "dotted"
	case BorderStyleThick:
		attr.Value = "thick"
	case BorderStyleDouble:
		attr.Value = "double"
	case BorderStyleHair:
		attr.Value = "hair"
	case BorderStyleMediumDashed:
		attr.Value = "mediumDashed"
	case BorderStyleDashDot:
		attr.Value = "dashDot"
	case BorderStyleMediumDashDot:
		attr.Value = "mediumDashDot"
	case BorderStyleDashDotDot:
		attr.Value = "dashDotDot"
	case BorderStyleMediumDashDotDot:
		attr.Value = "mediumDashDotDot"
	case BorderStyleSlantDashDot:
		attr.Value = "slantDashDot"
	default:
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *BorderStyleType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "none":
		*e = BorderStyleNone
	case "thin":
		*e = BorderStyleThin
	case "medium":
		*e = BorderStyleMedium
	case "dashed":
		*e = BorderStyleDashed
	case "dotted":
		*e = BorderStyleDotted
	case "thick":
		*e = BorderStyleThick
	case "double":
		*e = BorderStyleDouble
	case "hair":
		*e = BorderStyleHair
	case "mediumDashed":
		*e = BorderStyleMediumDashed
	case "dashDot":
		*e = BorderStyleDashDot
	case "mediumDashDot":
		*e = BorderStyleMediumDashDot
	case "dashDotDot":
		*e = BorderStyleDashDotDot
	case "mediumDashDotDot":
		*e = BorderStyleMediumDashDotDot
	case "slantDashDot":
		*e = BorderStyleSlantDashDot
	}

	return nil
}
