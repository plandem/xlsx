package primitives

import (
	"encoding/xml"
)

//ConditionType is a direct mapping of XSD ST_CfType
type ConditionType byte

//ConditionType maps for marshal/unmarshal process
var (
	ToConditionType   map[string]ConditionType
	FromConditionType map[ConditionType]string
)

func (t ConditionType) String() string {
	return FromConditionType[t]
}

//MarshalXMLAttr marshal ConditionType
func (t *ConditionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromConditionType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ConditionType
func (t *ConditionType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToConditionType[attr.Value]; ok {
		*t = v
	}

	return nil
}
