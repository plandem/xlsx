// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

//List of all possible values for OperatorType
const (
	_ ConditionOperatorType = iota
	ConditionOperatorLessThan
	ConditionOperatorLessThanOrEqual
	ConditionOperatorEqual
	ConditionOperatorNotEqual
	ConditionOperatorGreaterThanOrEqual
	ConditionOperatorGreaterThan
	ConditionOperatorBetween
	ConditionOperatorNotBetween
	ConditionOperatorContainsText
	ConditionOperatorNotContains
	ConditionOperatorBeginsWith
	ConditionOperatorEndsWith
)

func init() {
	FromConditionOperatorType = map[ConditionOperatorType]string{
		ConditionOperatorLessThan:           "lessThan",
		ConditionOperatorLessThanOrEqual:    "lessThanOrEqual",
		ConditionOperatorEqual:              "equal",
		ConditionOperatorNotEqual:           "notEqual",
		ConditionOperatorGreaterThanOrEqual: "greaterThanOrEqual",
		ConditionOperatorGreaterThan:        "greaterThan",
		ConditionOperatorBetween:            "between",
		ConditionOperatorNotBetween:         "notBetween",
		ConditionOperatorContainsText:       "containsText",
		ConditionOperatorNotContains:        "notContains",
		ConditionOperatorBeginsWith:         "beginsWith",
		ConditionOperatorEndsWith:           "endsWith",
	}

	ToConditionOperatorType = make(map[string]ConditionOperatorType, len(FromConditionOperatorType))
	for k, v := range FromConditionOperatorType {
		ToConditionOperatorType[v] = k
	}
}

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
