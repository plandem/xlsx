package primitives

import "encoding/xml"

//ConditionValueType is a direct mapping of XSD ST_CfvoType
type ConditionValueType byte

//ConditionType maps for marshal/unmarshal process
var (
	ToConditionValueType   map[string]ConditionValueType
	FromConditionValueType map[ConditionValueType]string
)

func (t ConditionValueType) String() string {
	return FromConditionValueType[t]
}

//MarshalXMLAttr marshal ConditionValueType
func (t *ConditionValueType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromConditionValueType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ConditionValueType
func (t *ConditionValueType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToConditionValueType[attr.Value]; ok {
		*t = v
	}

	return nil
}
