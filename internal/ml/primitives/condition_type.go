package primitives

import (
	"encoding/xml"
)

//ConditionType is a direct mapping of XSD ST_CfType
type ConditionType byte

//ConditionType maps for marshal/unmarshal process
var (
	toConditionType   map[string]ConditionType
	fromConditionType map[ConditionType]string
)

//List of all possible values for Type
const (
	_ ConditionType = iota
	ConditionTypeExpression
	ConditionTypeCellIs

	ConditionTypeColorScale      //
	ConditionTypeDataBar         //
	ConditionTypeIconSet         //
	ConditionTypeTop10           //
	ConditionTypeUniqueValues    //
	ConditionTypeDuplicateValues //

	ConditionTypeContainsText
	ConditionTypeNotContainsText
	ConditionTypeBeginsWith
	ConditionTypeEndsWith
	ConditionTypeContainsBlanks
	ConditionTypeNotContainsBlanks
	ConditionTypeContainsErrors
	ConditionTypeNotContainsErrors

	ConditionTypeTimePeriod
	ConditionTypeAboveAverage
)

func init() {
	fromConditionType = map[ConditionType]string{
		ConditionTypeExpression:        "expression",
		ConditionTypeCellIs:            "cellIs",
		ConditionTypeColorScale:        "colorScale",
		ConditionTypeDataBar:           "dataBar",
		ConditionTypeIconSet:           "iconSet",
		ConditionTypeTop10:             "top10",
		ConditionTypeUniqueValues:      "uniqueValues",
		ConditionTypeDuplicateValues:   "duplicateValues",
		ConditionTypeContainsText:      "containsText",
		ConditionTypeNotContainsText:   "notContainsText",
		ConditionTypeBeginsWith:        "beginsWith",
		ConditionTypeEndsWith:          "endsWith",
		ConditionTypeContainsBlanks:    "containsBlanks",
		ConditionTypeNotContainsBlanks: "notContainsBlanks",
		ConditionTypeContainsErrors:    "containsErrors",
		ConditionTypeNotContainsErrors: "notContainsErrors",
		ConditionTypeTimePeriod:        "timePeriod",
		ConditionTypeAboveAverage:      "aboveAverage",
	}

	toConditionType = make(map[string]ConditionType, len(fromConditionType))
	for k, v := range fromConditionType {
		toConditionType[v] = k
	}
}

func (t ConditionType) String() string {
	return fromConditionType[t]
}

//MarshalXMLAttr marshal ConditionType
func (t *ConditionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromConditionType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ConditionType
func (t *ConditionType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toConditionType[attr.Value]; ok {
		*t = v
	}

	return nil
}
