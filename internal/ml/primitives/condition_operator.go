package primitives

import (
	"encoding/xml"
)

//ConditionOperatorType is a direct mapping of XSD ST_ConditionalFormattingOperator
type ConditionOperatorType byte

//ConditionOperatorType maps for marshal/unmarshal process
var (
	ToConditionOperatorType   map[string]ConditionOperatorType
	FromConditionOperatorType map[ConditionOperatorType]string
)

func (t ConditionOperatorType) String() string {
	return FromConditionOperatorType[t]
}

//MarshalXMLAttr marshal ConditionOperatorType
func (t *ConditionOperatorType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromConditionOperatorType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ConditionOperatorType
func (t *ConditionOperatorType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToConditionOperatorType[attr.Value]; ok {
		*t = v
	}

	return nil
}
