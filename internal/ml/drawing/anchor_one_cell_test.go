// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/drawing/dml/chart"
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/ml/drawing"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestOneCellAnchor(t *testing.T) {
	type Entity struct {
		XMLName xml.Name               `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing entity"`
		DMLName dml.Name               `xml:",attr"`
		Anchor  *drawing.OneCellAnchor `xml:"oneCellAnchor"`
	}

	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
	<xdr:entity xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
		<xdr:oneCellAnchor>
			<xdr:from>
				<xdr:col>1</xdr:col>
				<xdr:colOff>2</xdr:colOff>
				<xdr:row>3</xdr:row>
				<xdr:rowOff>4</xdr:rowOff>
			</xdr:from>
			<xdr:ext cx="11" cy="22"></xdr:ext>
			<xdr:clientData></xdr:clientData>
			<xdr:graphicFrame macro="">
				<xdr:xfrm>
					<a:off x="111" y="222"></a:off>
					<a:ext cx="333" cy="444"></a:ext>
				</xdr:xfrm>
				<a:graphic>
					<a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/chart">
						<c:chart xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:id="rId1"></c:chart>
					</a:graphicData>
				</a:graphic>
			</xdr:graphicFrame>
		</xdr:oneCellAnchor>
	</xdr:entity>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	anchor := &drawing.OneCellAnchor{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "oneCellAnchor",
		},
		From: drawing.Marker{
			Col:       1,
			OffsetCol: dml.Coordinate(2),
			Row:       3,
			OffsetRow: dml.Coordinate(4),
		},
		Size: dml.PositiveSize2D{
			Height: 11,
			Width:  22,
		},
	}

	anchor.Frame = &drawing.Frame{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "graphicFrame",
		},
		Graphic: &dml.GraphicalObject{
			Data: &dml.GraphicalObjectData{
				Uri: "http://schemas.openxmlformats.org/drawingml/2006/chart",
				Chart: &chart.Ref{
					RID: "rId1",
				},
			},
		},
		Transform: &dml.Transform2D{
			Offset: &dml.Point2D{
				X: 111,
				Y: 222,
			},
			Size: &dml.PositiveSize2D{
				Height: 333,
				Width:  444,
			},
		},
		ReservedAttributes: ml.ReservedAttributes{
			Attrs: []xml.Attr{
				{
					Name: xml.Name{
						Local: "macro",
					},
				},
			},
		},
	}
	require.Equal(t, &Entity{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "entity",
		},
		Anchor: anchor,
	}, entity)

	//encode data should be same as original
	encode, err := xml.Marshal(entity)
	require.Nil(t, err)
	require.Equal(t, strings.NewReplacer("xdr:", "", ":xdr", "").Replace(data), string(encode))
}
