// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FontEffectType is a type to encode XSD ST_VerticalAlignRun
type FontEffectType ml.Property

//List of all possible values for FontEffectType
const (
	FontEffectBaseline    FontEffectType = "baseline"
	FontEffectSuperscript FontEffectType = "superscript"
	FontEffectSubscript   FontEffectType = "subscript"
)

//MarshalXML marshal FontEffectType
func (t *FontEffectType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.Property)(t).MarshalXML(e, start)
}

//UnmarshalXML unmarshal FontEffectType
func (t *FontEffectType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.Property)(t).UnmarshalXML(d, start)
}
