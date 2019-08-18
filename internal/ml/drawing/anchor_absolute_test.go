// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package drawing_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/ml/drawing"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestAbsoluteAnchor(t *testing.T) {
	type Entity struct {
		XMLName xml.Name                `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing entity"`
		DMLName dml.Name                `xml:",attr"`
		Anchor  *drawing.AbsoluteAnchor `xml:"absoluteAnchor"`
	}

	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
	<xdr:entity xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
		<xdr:absoluteAnchor>
			<xdr:pos x="1" y="2"></xdr:pos>
			<xdr:ext cx="3" cy="4"></xdr:ext>
			<xdr:clientData></xdr:clientData>
			<xdr:graphicFrame macro="">
				<xdr:xfrm>
					<a:off x="11" y="22"></a:off>
					<a:ext cx="33" cy="44"></a:ext>
				</xdr:xfrm>
				<a:graphic>
					<a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/chart">
						<c:chart xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:id="rId1"></c:chart>
					</a:graphicData>
				</a:graphic>
			</xdr:graphicFrame>
		</xdr:absoluteAnchor>
	</xdr:entity>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	anchor := &drawing.AbsoluteAnchor{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "absoluteAnchor",
		},
		Point: dml.Point2D{
			X: 1,
			Y: 2,
		},
		Size: dml.PositiveSize2D{
			Height: 3,
			Width:  4,
		},
	}

	anchor.Frame = &drawing.Frame{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "graphicFrame",
		},
		Graphic: &dml.GraphicFrame{
			Data: &dml.GraphicFrameData{
				Uri: "http://schemas.openxmlformats.org/drawingml/2006/chart",
				Chart: &dml.ChartRef{
					RID: "rId1",
				},
			},
		},
		Transform: &dml.Transform2D{
			Offset: &dml.Point2D{
				X: 11,
				Y: 22,
			},
			Size: &dml.PositiveSize2D{
				Height: 33,
				Width:  44,
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
