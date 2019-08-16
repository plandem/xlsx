// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"encoding/xml"
	"fmt"
)

//AnchorList is special type to hold all anchors with preserving order
type AnchorList []interface{}

const (
	anchorAbsolute      = "absoluteAnchor"
	anchorOneCellAnchor = "oneCellAnchor"
	anchorTwoCellAnchor = "twoCellAnchor"
	errorUnknownAnchor  = "unknown type of anchor: %s"
)

//Add another anchor to list
func (a *AnchorList) Add(anchor interface{}) {
	*a = append(*a, anchor)
}

//UnmarshalXML unmarshal AnchorList
func (a *AnchorList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var anchor interface{}

	switch start.Name.Local {
	case anchorAbsolute:
		anchor = &AbsoluteAnchor{}
	case anchorOneCellAnchor:
		anchor = &OneCellAnchor{}
	case anchorTwoCellAnchor:
		anchor = &TwoCellAnchor{}
	default:
		return fmt.Errorf(errorUnknownAnchor, start)
	}

	if err := d.DecodeElement(anchor, &start); err != nil {
		return err
	}

	a.Add(anchor)
	return nil
}
