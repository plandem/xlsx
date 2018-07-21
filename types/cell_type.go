package types

import (
	"encoding/xml"
)

//CellType is a type to encode XSD ST_CellType
type CellType byte

//List of all possible values for CellType
const (
	CellTypeGeneral CellType = iota
	CellTypeBool
	CellTypeDate
	CellTypeNumber
	CellTypeError
	CellTypeSharedString
	CellTypeFormula
	CellTypeInlineString
)

var (
	toCellType   map[string]CellType
	fromCellType map[CellType]string
)

func init() {
	fromCellType = map[CellType]string{
		CellTypeBool:         "b",
		CellTypeDate:         "d",
		CellTypeNumber:       "n",
		CellTypeError:        "e",
		CellTypeSharedString: "s",
		CellTypeFormula:      "str",
		CellTypeInlineString: "inlineStr",
	}

	toCellType = make(map[string]CellType, len(fromCellType))
	for k, v := range fromCellType {
		toCellType[v] = k
	}
}

func (e CellType) String() string {
	return fromCellType[e]
}

func (e *CellType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromCellType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *CellType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toCellType[attr.Value]; ok {
		*e = v
	}

	return nil
}
