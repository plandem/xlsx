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

func TestDrawing(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
	<xdr:wsDr xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
		<xdr:twoCellAnchor>
			<xdr:from>
				<xdr:col>4</xdr:col>
				<xdr:colOff>19050</xdr:colOff>
				<xdr:row>1</xdr:row>
				<xdr:rowOff>0</xdr:rowOff>
			</xdr:from>
			<xdr:to>
				<xdr:col>9</xdr:col>
				<xdr:colOff>463550</xdr:colOff>
				<xdr:row>14</xdr:row>
				<xdr:rowOff>101600</xdr:rowOff>
			</xdr:to>
			<xdr:clientData></xdr:clientData>
		</xdr:twoCellAnchor>
		<xdr:oneCellAnchor>
			<xdr:from>
				<xdr:col>4</xdr:col>
				<xdr:colOff>19050</xdr:colOff>
				<xdr:row>1</xdr:row>
				<xdr:rowOff>0</xdr:rowOff>
			</xdr:from>
			<xdr:ext cx="8671719" cy="6290469"></xdr:ext>
			<xdr:clientData></xdr:clientData>
		</xdr:oneCellAnchor>
		<xdr:absoluteAnchor>
			<xdr:pos x="0" y="0"></xdr:pos>
			<xdr:ext cx="8671719" cy="6290469"></xdr:ext>
			<xdr:clientData></xdr:clientData>
			<xdr:graphicFrame macro="">
				<a:graphic>
					<a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/chart">
						<c:chart xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:id="rId1"></c:chart>
					</a:graphicData>
				</a:graphic>
				<xdr:xfrm>
					<a:off x="0" y="0"></a:off>
					<a:ext cx="0" cy="0"></a:ext>
				</xdr:xfrm>
			</xdr:graphicFrame>
		</xdr:absoluteAnchor>
	</xdr:wsDr>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	drw := &drawing.Drawing{}
	err := decoder.DecodeElement(drw, nil)
	require.Nil(t, err)

	//name
	require.Equal(t, xml.Name{
		Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
		Local: "wsDr",
	}, drw.XMLName)

	//first
	require.Equal(t, &drawing.TwoCellAnchor{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "twoCellAnchor",
		},
		From: drawing.Marker{
			Row:       1,
			Col:       4,
			OffsetRow: dml.Coordinate(0),
			OffsetCol: dml.Coordinate(19050),
		},
		To: drawing.Marker{
			Row:       14,
			Col:       9,
			OffsetRow: dml.Coordinate(101600),
			OffsetCol: dml.Coordinate(463550),
		},
	}, (*drw.AnchorList)[0])

	//second
	require.Equal(t, &drawing.OneCellAnchor{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "oneCellAnchor",
		},
		From: drawing.Marker{
			Row:       1,
			Col:       4,
			OffsetRow: dml.Coordinate(0),
			OffsetCol: dml.Coordinate(19050),
		},
		Size: dml.PositiveSize2D{
			Height: 8671719,
			Width:  6290469,
		},
	}, (*drw.AnchorList)[1])

	//third
	require.IsType(t, &drawing.AbsoluteAnchor{}, (*drw.AnchorList)[2])
	absAnchor := (*drw.AnchorList)[2].(*drawing.AbsoluteAnchor)
	require.Equal(t, dml.PositiveSize2D{
		Height: 8671719,
		Width:  6290469,
	}, absAnchor.Size)

	require.Equal(t, ml.ReservedAttributes{
		Attrs: []xml.Attr{
			{
				Name: xml.Name{Local: "macro"},
			},
		},
	}, absAnchor.GraphicFrame.ReservedAttributes)

	require.Equal(t, &dml.Transform2D{
		Offset: &dml.Point2D{
			X: 0,
			Y: 0,
		},
		Size: &dml.PositiveSize2D{
			Height: 0,
			Width:  0,
		},
	}, absAnchor.GraphicFrame.Transform)

	//encode data should be same as original
	encode, err := xml.Marshal(drw)
	require.Nil(t, err)
	require.Equal(t, strings.NewReplacer("xdr:", "", ":xdr", "").Replace(data), string(encode))
}
