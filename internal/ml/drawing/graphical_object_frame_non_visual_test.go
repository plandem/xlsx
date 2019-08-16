// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/xlsx/internal/ml/drawing"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGraphicalObjectFrameNonVisual(t *testing.T) {
	type Entity struct {
		XMLName   xml.Name                               `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing entity"`
		DMLName   dml.Name                               `xml:",attr"`
		NonVisual *drawing.GraphicalObjectFrameNonVisual `xml:"nvGraphicFramePr"`
	}

	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
	<xdr:entity xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
		<xdr:nvGraphicFramePr>
			<xdr:cNvPr id="2" name="Chart 1"></xdr:cNvPr>
			<xdr:cNvGraphicFramePr></xdr:cNvGraphicFramePr>
		</xdr:nvGraphicFramePr>
	</xdr:entity>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	object := &drawing.GraphicalObjectFrameNonVisual{
		DrawingProperties: &dml.NonVisualDrawingProperties{
			ID: 2,
			Name: "Chart 1",
		},
		FrameProperties: &dml.NonVisualGraphicFrameProperties{},
	}

	require.Equal(t, &Entity{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "entity",
		},
		NonVisual: object,
	}, entity)

	//encode data should be same as original
	encode, err := xml.Marshal(entity)
	require.Nil(t, err)
	require.Equal(t, strings.NewReplacer("xdr:", "", ":xdr", "").Replace(data), string(encode))
}
