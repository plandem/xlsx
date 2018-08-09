package primitives

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//UnderlineType is a type to encode XSD CT_UnderlineProperty
type UnderlineType ml.Property

//MarshalXML marshal UnderlineType
func (t *UnderlineType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.Property)(t).MarshalXML(e, start)
}

//UnmarshalXML unmarshal UnderlineType
func (t *UnderlineType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.Property)(t).UnmarshalXML(d, start)
}
