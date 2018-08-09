package primitives

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Text is textual type that can have leading/trailing spaces that must be preserved
type Text string

//MarshalXML marshal Text
func (t *Text) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	value := string(*t)

	//need to preserve space?
	if len(value) > 0 && (value[0] == 32 || value[len(value)-1] == 32) {
		start.Attr = append(start.Attr, ml.AttrPreserveSpace)
	}

	return e.EncodeElement(value, start)
}
