package types

import "encoding/xml"

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

func (e *CellFormulaType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *e {
	case CellFormulaTypeNormal:
		attr.Value = "normal"
	case CellFormulaTypeArray:
		attr.Value = "array"
	case CellFormulaTypeDataTable:
		attr.Value = "dataTable"
	case CellFormulaTypeShared:
		attr.Value = "shared"
	default:
		attr = xml.Attr{}
	}

	return attr, nil
}

func (e *CellFormulaType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "normal":
		*e = CellFormulaTypeNormal
	case "array":
		*e = CellFormulaTypeArray
	case "dataTable":
		*e = CellFormulaTypeDataTable
	case "shared":
		*e = CellFormulaTypeShared
	}

	return nil
}
