package primitives

import "encoding/xml"

//UpdateLinksType is a type to encode XSD ST_UpdateLinks
type UpdateLinksType byte

//List of all possible values for UpdateLinksType
const (
	_ UpdateLinksType = iota
	UpdateLinksTypeUserSet
	UpdateLinksTypeNever
	UpdateLinksTypeAlways
)

var (
	toUpdateLinksType   map[string]UpdateLinksType
	fromUpdateLinksType map[UpdateLinksType]string
)

func init() {
	fromUpdateLinksType = map[UpdateLinksType]string{
		UpdateLinksTypeUserSet: "userSet",
		UpdateLinksTypeNever:   "never",
		UpdateLinksTypeAlways:  "always",
	}

	toUpdateLinksType = make(map[string]UpdateLinksType, len(fromUpdateLinksType))
	for k, v := range fromUpdateLinksType {
		toUpdateLinksType[v] = k
	}
}

func (e UpdateLinksType) String() string {
	return fromUpdateLinksType[e]
}

//MarshalXMLAttr marshal UpdateLinksType
func (e *UpdateLinksType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromUpdateLinksType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal UpdateLinksType
func (e *UpdateLinksType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toUpdateLinksType[attr.Value]; ok {
		*e = v
	}

	return nil
}
