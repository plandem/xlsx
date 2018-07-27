package styles

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//UnderlineType is a type to encode XSD CT_UnderlineProperty
type UnderlineType ml.Property

func (t *UnderlineType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.Property)(t).MarshalXML(e, start)
}

func (t *UnderlineType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.Property)(t).UnmarshalXML(d, start)
}
