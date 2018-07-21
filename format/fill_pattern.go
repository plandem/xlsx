package format

import (
	"encoding/xml"
)

//PatternType is a type to encode XSD ST_PatternType
type PatternType byte

//List of all possible values for PatternType
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

var (
	toPatternType   map[string]PatternType
	fromPatternType map[PatternType]string
)

func init() {
	fromPatternType = map[PatternType]string{
		PatternTypeNone:            "none",
		PatternTypeSolid:           "solid",
		PatternTypeMediumGray:      "mediumGray",
		PatternTypeDarkGray:        "darkGray",
		PatternTypeLightGray:       "lightGray",
		PatternTypeDarkHorizontal:  "darkHorizontal",
		PatternTypeDarkVertical:    "darkVertical",
		PatternTypeDarkDown:        "darkDown",
		PatternTypeDarkUp:          "darkUp",
		PatternTypeDarkGrid:        "darkGrid",
		PatternTypeDarkTrellis:     "darkTrellis",
		PatternTypeLightHorizontal: "lightHorizontal",
		PatternTypeLightVertical:   "lightVertical",
		PatternTypeLightDown:       "lightDown",
		PatternTypeLightUp:         "lightUp",
		PatternTypeLightGrid:       "lightGrid",
		PatternTypeLightTrellis:    "lightTrellis",
		PatternTypeGray125:         "gray125",
		PatternTypeGray0625:        "gray0625",
	}

	toPatternType = make(map[string]PatternType, len(fromPatternType))
	for k, v := range fromPatternType {
		toPatternType[v] = k
	}
}

func (e *PatternType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromPatternType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *PatternType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toPatternType[attr.Value]; ok {
		*e = v
	}

	return nil
}
