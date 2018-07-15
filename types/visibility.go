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

func (e *VisibilityType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *e {
	case VisibilityTypeVisible:
		attr.Value = "visible"
	case VisibilityTypeHidden:
		attr.Value = "hidden"
	case VisibilityTypeVeryHidden:
		attr.Value = "veryHidden"
	default:
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *VisibilityType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "visible":
		*e = VisibilityTypeVisible
	case "hidden":
		*e = VisibilityTypeHidden
	case "veryHidden":
		*e = VisibilityTypeVeryHidden
	}

	return nil
}
