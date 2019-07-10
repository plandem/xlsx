// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
)

//VAlignType is a type to encode XSD ST_VerticalAlignment
type VAlignType byte

//VAlignType maps for marshal/unmarshal process
var (
	ToVAlignType   map[string]VAlignType
	FromVAlignType map[VAlignType]string
)

func (t VAlignType) String() string {
	return FromVAlignType[t]
}

//MarshalXMLAttr marshal VAlignType
func (t *VAlignType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromVAlignType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal VAlignType
func (t *VAlignType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToVAlignType[attr.Value]; ok {
		*t = v
	}

	return nil
}
