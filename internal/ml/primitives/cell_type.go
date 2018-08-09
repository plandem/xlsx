package primitives

import (
	"encoding/xml"
)

//CellType is a type to encode XSD ST_CellType
type CellType byte

var (
	ToCellType   map[string]CellType
	FromCellType map[CellType]string
)

func (e CellType) String() string {
	return FromCellType[e]
}

//MarshalXMLAttr marshal CellType
func (e *CellType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromCellType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal CellType
func (e *CellType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToCellType[attr.Value]; ok {
		*e = v
	}

	return nil
}
