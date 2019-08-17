// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"encoding/xml"
	"fmt"
)

const (
	groupShape        = "sp"
	groupGroup        = "grpSp"
	groupFrame        = "graphicFrame"
	groupConnector    = "cxnSp"
	groupPicture      = "pic"
	errorUnknownGroup = "unknown type of group's item: %s"
)

//GroupList is special type to hold all anchors with preserving order
type GroupList []interface{}

//Add another anchor to list
func (g *GroupList) Add(item interface{}) {
	*g = append(*g, item)
}

//UnmarshalXML unmarshal GroupList
func (g *GroupList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var item interface{}

	switch start.Name.Local {
	case groupShape:
		item = &Shape{}
	case groupGroup:
		item = &Group{}
	case groupFrame:
		item = &Frame{}
	case groupConnector:
		item = &Connector{}
	case groupPicture:
		item = &Picture{}
	default:
		return fmt.Errorf(errorUnknownGroup, start)
	}

	if err := d.DecodeElement(item, &start); err != nil {
		return err
	}

	g.Add(item)
	return nil
}
