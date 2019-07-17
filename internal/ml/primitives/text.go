// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
	"github.com/plandem/ooxml/index"
	"github.com/plandem/ooxml/ml"
	"strings"
	"unicode"
)

//Text is textual type that can have leading/trailing whitespace or newlines that must be preserved
type Text string

//Hash builds hash code for all required values of Text to use as unique index
func (t Text) Hash() index.Code {
	return index.Hash(string(t))
}

//MarshalXML marshal Text
func (t *Text) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	value := string(*t)

	//need to preserve space?
	if len(value) > 0 && (unicode.IsSpace(rune(value[0])) || unicode.IsSpace(rune(value[len(value)-1])) || strings.IndexByte(value, '\n') != -1) {
		start.Attr = append(start.Attr, ml.AttrPreserveSpace)
	}

	return e.EncodeElement(value, start)
}
