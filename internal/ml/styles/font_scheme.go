package styles

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FontSchemeType is a type to encode XSD ST_FontScheme
type FontSchemeType ml.Property

func (t *FontSchemeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.Property)(t).MarshalXML(e, start)
}

func (t *FontSchemeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.Property)(t).UnmarshalXML(d, start)
}
