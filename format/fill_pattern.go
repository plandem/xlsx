package format

import "encoding/xml"

//PatternType is a type to encode XSD ST_PatternType
type PatternType byte

const (
	_ PatternType = iota
	PatternTypeNone
	PatternTypeSolid
	PatternTypeMediumGray
	PatternTypeDarkGray
	PatternTypeLightGray
	PatternTypeDarkHorizontal
	PatternTypeDarkVertical
	PatternTypeDarkDown
	PatternTypeDarkUp
	PatternTypeDarkGrid
	PatternTypeDarkTrellis
	PatternTypeLightHorizontal
	PatternTypeLightVertical
	PatternTypeLightDown
	PatternTypeLightUp
	PatternTypeLightGrid
	PatternTypeLightTrellis
	PatternTypeGray125
	PatternTypeGray0625
)

func (e *PatternType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *e {
	case PatternTypeNone:
		attr.Value = "none"
	case PatternTypeSolid:
		attr.Value = "solid"
	case PatternTypeMediumGray:
		attr.Value = "mediumGray"
	case PatternTypeDarkGray:
		attr.Value = "darkGray"
	case PatternTypeLightGray:
		attr.Value = "lightGray"
	case PatternTypeDarkHorizontal:
		attr.Value = "darkHorizontal"
	case PatternTypeDarkVertical:
		attr.Value = "darkVertical"
	case PatternTypeDarkDown:
		attr.Value = "darkDown"
	case PatternTypeDarkUp:
		attr.Value = "darkUp"
	case PatternTypeDarkGrid:
		attr.Value = "darkGrid"
	case PatternTypeDarkTrellis:
		attr.Value = "darkTrellis"
	case PatternTypeLightHorizontal:
		attr.Value = "lightHorizontal"
	case PatternTypeLightVertical:
		attr.Value = "lightVertical"
	case PatternTypeLightDown:
		attr.Value = "lightDown"
	case PatternTypeLightUp:
		attr.Value = "lightUp"
	case PatternTypeLightGrid:
		attr.Value = "lightGrid"
	case PatternTypeLightTrellis:
		attr.Value = "lightTrellis"
	case PatternTypeGray125:
		attr.Value = "gray125"
	case PatternTypeGray0625:
		attr.Value = "gray0625"
	default:
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *PatternType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "none":
		*e = PatternTypeNone
	case "solid":
		*e = PatternTypeSolid
	case "mediumGray":
		*e = PatternTypeMediumGray
	case "darkGray":
		*e = PatternTypeDarkGray
	case "lightGray":
		*e = PatternTypeLightGray
	case "darkHorizontal":
		*e = PatternTypeDarkHorizontal
	case "darkVertical":
		*e = PatternTypeDarkVertical
	case "darkDown":
		*e = PatternTypeDarkDown
	case "darkUp":
		*e = PatternTypeDarkUp
	case "darkGrid":
		*e = PatternTypeDarkGrid
	case "darkTrellis":
		*e = PatternTypeDarkTrellis
	case "lightHorizontal":
		*e = PatternTypeLightHorizontal
	case "lightVertical":
		*e = PatternTypeLightVertical
	case "lightDown":
		*e = PatternTypeLightDown
	case "lightUp":
		*e = PatternTypeLightUp
	case "lightGrid":
		*e = PatternTypeLightGrid
	case "lightTrellis":
		*e = PatternTypeLightTrellis
	case "gray125":
		*e = PatternTypeGray125
	case "gray0625":
		*e = PatternTypeGray0625
	}

	return nil
}
