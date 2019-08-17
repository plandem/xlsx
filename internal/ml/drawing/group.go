// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
)

//Group is a direct mapping of XSD CT_GroupShape
type Group struct {
	XMLName   xml.Name             `xml:"grpSp"`
	NonVisual *GroupNonVisual      `xml:"nvGrpSpPr"`
	Group     *dml.GroupProperties `xml:"grpSpPr"`
	Items     *GroupList           `xml:",any"`
}

//GroupNonVisual is a direct mapping of XSD CT_GroupShapeNonVisual
type GroupNonVisual struct {
	CommonProperties *dml.NonVisualCommonProperties `xml:"cNvPr"`
	GroupProperties  *dml.NonVisualGroupProperties  `xml:"cNvGrpSpPr"`
}
