package types

import "encoding/xml"

//ObjectsType is a type to encode XSD ST_Objects
type ObjectsType byte

//List of all possible values for ObjectsType
const (
	_ ObjectsType = iota
	ObjectsTypeAll
	ObjectsTypePlaceholders
	ObjectsTypeNone
)

func (e *ObjectsType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *e {
	case ObjectsTypeAll:
		attr.Value = "all"
	case ObjectsTypePlaceholders:
		attr.Value = "placeholders"
	case ObjectsTypeNone:
		attr.Value = "none"
	default:
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *ObjectsType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "all":
		*e = ObjectsTypeAll
	case "placeholders":
		*e = ObjectsTypePlaceholders
	case "none":
		*e = ObjectsTypeNone
	}

	return nil
}
