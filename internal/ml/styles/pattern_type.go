package styles

import (
	"encoding/xml"
)

//PatternType is a type to encode XSD ST_PatternType
type PatternType byte

var (
	ToPatternType   map[string]PatternType
	FromPatternType map[PatternType]string
)

func (e PatternType) String() string {
	return FromPatternType[e]
}

func (e *PatternType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromPatternType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *PatternType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToPatternType[attr.Value]; ok {
		*e = v
	}

	return nil
}
