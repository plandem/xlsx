package ml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FontSchemeType is a type to encode XSD ST_FontScheme
type FontSchemeType ml.Property

func (f *FontSchemeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.Property)(f).MarshalXML(e, start)
}

func (f *FontSchemeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.Property)(f).UnmarshalXML(d, start)
}
