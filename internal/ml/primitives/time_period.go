package primitives

import (
	"encoding/xml"
)

//TimePeriodType is a direct mapping of XSD ST_TimePeriod
type TimePeriodType byte

//TimePeriodType maps for marshal/unmarshal process
var (
	toTimePeriodType   map[string]TimePeriodType
	fromTimePeriodType map[TimePeriodType]string
)

//List of all possible values for TimePeriodType
const (
	_ TimePeriodType = iota
	TimePeriodToday
	TimePeriodYesterday
	TimePeriodTomorrow
	TimePeriodLast7Days
	TimePeriodThisMonth
	TimePeriodLastMonth
	TimePeriodNextMonth
	TimePeriodThisWeek
	TimePeriodLastWeek
	TimePeriodNextWeek
)

func init() {
	fromTimePeriodType = map[TimePeriodType]string{
		TimePeriodToday:     "today",
		TimePeriodYesterday: "yesterday",
		TimePeriodTomorrow:  "tomorrow",
		TimePeriodLast7Days: "last7Days",
		TimePeriodThisMonth: "thisMonth",
		TimePeriodLastMonth: "lastMonth",
		TimePeriodNextMonth: "nextMonth",
		TimePeriodThisWeek:  "thisWeek",
		TimePeriodLastWeek:  "lastWeek",
		TimePeriodNextWeek:  "nextWeek",
	}

	toTimePeriodType = make(map[string]TimePeriodType, len(fromTimePeriodType))
	for k, v := range fromTimePeriodType {
		toTimePeriodType[v] = k
	}
}

func (t TimePeriodType) String() string {
	return fromTimePeriodType[t]
}

//MarshalXMLAttr marshal TimePeriodType
func (t *TimePeriodType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromTimePeriodType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal TimePeriodType
func (t *TimePeriodType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toTimePeriodType[attr.Value]; ok {
		*t = v
	}

	return nil
}
