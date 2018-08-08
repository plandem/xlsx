package primitives

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

var (
	toObjectsType   map[string]ObjectsType
	fromObjectsType map[ObjectsType]string
)

func init() {
	fromObjectsType = map[ObjectsType]string{
		ObjectsTypeAll:          "all",
		ObjectsTypePlaceholders: "placeholders",
		ObjectsTypeNone:         "none",
	}

	toObjectsType = make(map[string]ObjectsType, len(fromObjectsType))
	for k, v := range fromObjectsType {
		toObjectsType[v] = k
	}
}

func (e ObjectsType) String() string {
	return fromObjectsType[e]
}

//MarshalXMLAttr marshal ObjectsType
func (e *ObjectsType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromObjectsType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ObjectsType
func (e *ObjectsType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toObjectsType[attr.Value]; ok {
		*e = v
	}

	return nil
}
