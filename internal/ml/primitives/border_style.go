// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
)

//BorderStyleType is a type to encode XSD BorderStyleType
type BorderStyleType byte

//BorderStyleType maps for marshal/unmarshal process
var (
	ToBorderStyleType   map[string]BorderStyleType
	FromBorderStyleType map[BorderStyleType]string
)

func (t BorderStyleType) String() string {
	return FromBorderStyleType[t]
}

//MarshalXMLAttr marshal BorderStyleType
func (t *BorderStyleType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromBorderStyleType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal BorderStyleType
func (t *BorderStyleType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToBorderStyleType[attr.Value]; ok {
		*t = v
	}

	return nil
}
