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

func (t PatternType) String() string {
	return FromPatternType[t]
}

func (t *PatternType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromPatternType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (t *PatternType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToPatternType[attr.Value]; ok {
		*t = v
	}

	return nil
}
