package primitives

import (
	"encoding/xml"
)

//TimePeriodType is a direct mapping of XSD ST_TimePeriod
type TimePeriodType byte

//TimePeriodType maps for marshal/unmarshal process
var (
	ToTimePeriodType   map[string]TimePeriodType
	FromTimePeriodType map[TimePeriodType]string
)

func (t TimePeriodType) String() string {
	return FromTimePeriodType[t]
}

//MarshalXMLAttr marshal TimePeriodType
func (t *TimePeriodType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromTimePeriodType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal TimePeriodType
func (t *TimePeriodType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToTimePeriodType[attr.Value]; ok {
		*t = v
	}

	return nil
}
