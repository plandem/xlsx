package primitives

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FontCharsetType is a type to encode charset of font
type FontCharsetType ml.PropertyInt

//MarshalXML marshal FontCharsetType
func (t *FontCharsetType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.PropertyInt)(t).MarshalXML(e, start)
}

//UnmarshalXML unmarshal FontCharsetType
func (t *FontCharsetType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.PropertyInt)(t).UnmarshalXML(d, start)
}
