package primitives

import (
	"encoding/xml"
)

//CellFormulaType is a type to encode XSD ST_CellFormulaType
type CellFormulaType byte

//List of all possible values for CellFormulaType
const (
	_ CellFormulaType = iota
	CellFormulaTypeNormal
	CellFormulaTypeArray
	CellFormulaTypeDataTable
	CellFormulaTypeShared
)

var (
	toCellFormulaType   map[string]CellFormulaType
	fromCellFormulaType map[CellFormulaType]string
)

func init() {
	fromCellFormulaType = map[CellFormulaType]string{
		CellFormulaTypeNormal:    "normal",
		CellFormulaTypeArray:     "array",
		CellFormulaTypeDataTable: "dataTable",
		CellFormulaTypeShared:    "shared",
	}

	toCellFormulaType = make(map[string]CellFormulaType, len(fromCellFormulaType))
	for k, v := range fromCellFormulaType {
		toCellFormulaType[v] = k
	}
}

func (e CellFormulaType) String() string {
	return fromCellFormulaType[e]
}

//MarshalXMLAttr marshal CellFormulaType
func (e *CellFormulaType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromCellFormulaType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal CellFormulaType
func (e *CellFormulaType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toCellFormulaType[attr.Value]; ok {
		*e = v
	}

	return nil
}
