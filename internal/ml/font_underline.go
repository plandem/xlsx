package ml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//UnderlineType is a type to encode XSD CT_UnderlineProperty
type UnderlineType ml.Property

func (f *UnderlineType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.Property)(f).MarshalXML(e, start)
}

func (f *UnderlineType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.Property)(f).UnmarshalXML(d, start)
}
