// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
)

//Connector is a direct mapping of XSD CT_Connector
type Connector struct {
	NonVisual *ConnectorNonVisual  `xml:"nvCxnSpPr"`
	Shape     *dml.ShapeProperties `xml:"spPr"`
	Style     *dml.ShapeStyle      `xml:"style"`
	ml.ReservedAttributes
}
