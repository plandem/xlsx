// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

//IsAllowed check if type is in list of allowed types
func (t *ConditionValueType) IsAllowed(allowed ...ConditionValueType) bool {
	for _, a := range allowed {
		if *t == a {
			return true
		}
	}

	return false
}
