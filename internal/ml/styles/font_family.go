package styles

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FontFamilyType is a type to encode XSD ST_FontFamily
type FontFamilyType ml.PropertyInt

func (t *FontFamilyType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.PropertyInt)(t).MarshalXML(e, start)
}

func (t *FontFamilyType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.PropertyInt)(t).UnmarshalXML(d, start)
}
