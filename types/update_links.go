package types

import "encoding/xml"

//UpdateLinksType is a type to encode XSD ST_UpdateLinks
type UpdateLinksType byte

const (
	_ UpdateLinksType = iota
	UpdateLinksTypeUserSet
	UpdateLinksTypeNever
	UpdateLinksTypeAlways
)

func (e *UpdateLinksType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *e {
	case UpdateLinksTypeUserSet:
		attr.Value = "userSet"
	case UpdateLinksTypeNever:
		attr.Value = "never"
	case UpdateLinksTypeAlways:
		attr.Value = "always"
	default:
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *UpdateLinksType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "userSet":
		*e = UpdateLinksTypeUserSet
	case "never":
		*e = UpdateLinksTypeNever
	case "always":
		*e = UpdateLinksTypeAlways
	}

	return nil
}
