// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FontFamilyType is a type to encode XSD ST_FontFamily
type FontFamilyType ml.PropertyInt

//MarshalXML marshal FontFamilyType
func (t *FontFamilyType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.PropertyInt)(t).MarshalXML(e, start)
}

//UnmarshalXML unmarshal FontFamilyType
func (t *FontFamilyType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.PropertyInt)(t).UnmarshalXML(d, start)
}
