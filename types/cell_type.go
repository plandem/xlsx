package types

import (
	"encoding/xml"
)

//CellType is a type to encode XSD ST_CellType
type CellType byte

//List of all possible values for CellType
const (
	CellTypeBool         CellType = 'b'
	CellTypeDate         CellType = 'd'
	CellTypeNumber       CellType = 'n'
	CellTypeError        CellType = 'e'
	CellTypeSharedString CellType = 's'

	CellTypeFormula      CellType = 'f' //meta type for 'str'
	CellTypeInlineString CellType = 'i' //meta type for 'inlineStr'
	CellTypeGeneral      CellType = 'g' //meta type for general type
)

func (e *CellType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *e {
	case CellTypeInlineString:
		attr.Value = "inlineStr"
	case CellTypeFormula:
		attr.Value = "str"
	case CellTypeGeneral:
		attr = xml.Attr{}
	default:
		attr.Value = string(*e)
	}

	return attr, nil
}

func (e *CellType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "inlineStr":
		*e = CellTypeInlineString
	case "str":
		*e = CellTypeFormula
	default:
		*e = CellType(attr.Value[0])
	}

	return nil
}
