// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing

import (
	"github.com/plandem/ooxml/drawing/dml"
)

//Group is a direct mapping of XSD CT_GroupShape
type Group struct {
	NonVisual *GroupShapeNonVisual `xml:"nvGrpSpPr"`
	Group     *dml.GroupProperties `xml:"grpSpPr"`
	Items     *GroupList           `xml:",any"`
}
