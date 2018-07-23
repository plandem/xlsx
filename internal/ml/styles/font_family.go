package styles

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FontFamilyType is a type to encode XSD ST_FontFamily
type FontFamilyType ml.PropertyInt

func (f *FontFamilyType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.PropertyInt)(f).MarshalXML(e, start)
}

func (f *FontFamilyType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.PropertyInt)(f).UnmarshalXML(d, start)
}
