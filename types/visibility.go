package types

import "encoding/xml"

//VisibilityType is a type to encode XSD ST_Visibility and ST_SheetState
type VisibilityType byte

//List of all possible values for VisibilityType
const (
	_ VisibilityType = iota
	VisibilityTypeVisible
	VisibilityTypeHidden
	VisibilityTypeVeryHidden
)

var (
	toVisibilityType   map[string]VisibilityType
	fromVisibilityType map[VisibilityType]string
)

func init() {
	fromVisibilityType = map[VisibilityType]string{
		VisibilityTypeVisible:    "visible",
		VisibilityTypeHidden:     "hidden",
		VisibilityTypeVeryHidden: "veryHidden",
	}

	toVisibilityType = make(map[string]VisibilityType, len(fromVisibilityType))
	for k, v := range fromVisibilityType {
		toVisibilityType[v] = k
	}
}

func (e VisibilityType) String() string {
	return fromVisibilityType[e]
}

func (e *VisibilityType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromVisibilityType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *VisibilityType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toVisibilityType[attr.Value]; ok {
		*e = v
	}

	return nil
}
