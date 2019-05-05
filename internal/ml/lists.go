package ml

import (
	"encoding/xml"
)

//DiffStyleList is a direct mapping of XSD CT_Dxfs
type DiffStyleList struct {
	Count int          `xml:"count,attr"`
	Items []*DiffStyle `xml:"dxf,omitempty"`
}

//NamedStyleInfoList is a direct mapping of XSD CT_CellStyles
type NamedStyleInfoList struct {
	Count int               `xml:"count,attr"`
	Items []*NamedStyleInfo `xml:"cellStyle,omitempty"`
}

//DirectStyleList is a direct mapping of XSD CT_CellXfs
type DirectStyleList struct {
	Count int            `xml:"count,attr"`
	Items []*DirectStyle `xml:"xf,omitempty"`
}

//NamedStyleList is a direct mapping of XSD cellStyleXfs
type NamedStyleList struct {
	Count int           `xml:"count,attr"`
	Items []*NamedStyle `xml:"xf,omitempty"`
}

//BorderList is a direct mapping of XSD CT_Borders
type BorderList struct {
	Count int       `xml:"count,attr"`
	Items []*Border `xml:"border,omitempty"`
}

//FontList is a direct mapping of XSD CT_Fonts
type FontList struct {
	Count int     `xml:"count,attr"`
	Items []*Font `xml:"font,omitempty"`
}

//FillList is a direct mapping of XSD CT_Fills
type FillList struct {
	Count int     `xml:"count,attr"`
	Items []*Fill `xml:"fill,omitempty"`
}

//NumberFormatList is a direct mapping of XSD CT_NumFmts
type NumberFormatList struct {
	Count int             `xml:"count,attr"`
	Items []*NumberFormat `xml:"numFmt,omitempty"`
}

func (r *DiffStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

func (r *NamedStyleInfoList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

func (r *DirectStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

func (r *NamedStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

func (r *BorderList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

func (r *FontList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

func (r *FillList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

func (r *NumberFormatList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}
