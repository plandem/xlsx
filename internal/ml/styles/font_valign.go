package styles

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FontVAlignType is a type to encode XSD ST_VerticalAlignRun
type FontVAlignType ml.Property

//List of all possible values for FontVAlignType
const (
	FontVAlignBaseline    FontVAlignType = "baseline"
	FontVAlignSuperscript FontVAlignType = "superscript"
	FontVAlignSubscript   FontVAlignType = "subscript"
)

func (t *FontVAlignType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.Property)(t).MarshalXML(e, start)
}

func (t *FontVAlignType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.Property)(t).UnmarshalXML(d, start)
}
