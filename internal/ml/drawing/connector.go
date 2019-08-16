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
	Style     *dml.ShapeStyle      `xml:"style,omitempty"`
	ml.ReservedAttributes
}

//ConnectorNonVisual is a direct mapping of XSD CT_ConnectorNonVisual
type ConnectorNonVisual struct {
	DrawingProperties   *dml.NonVisualCommonProperties    `xml:"cNvPr"`
	ConnectorProperties *dml.NonVisualConnectorProperties `xml:"cNvCxnSpPr"`
}
