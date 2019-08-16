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

func TestAnchorList(t *testing.T) {
	type Entity struct {
		XMLName xml.Name            `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing entity"`
		DMLName dml.Name            `xml:",attr"`
		Items   *drawing.AnchorList `xml:",any"`
	}

	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
	<xdr:entity xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
		<xdr:absoluteAnchor>
			<xdr:pos x="1" y="2"></xdr:pos>
			<xdr:ext cx="3" cy="4"></xdr:ext>
			<xdr:clientData></xdr:clientData>
		</xdr:absoluteAnchor>
		<xdr:twoCellAnchor>
			<xdr:from>
				<xdr:col>1</xdr:col>
				<xdr:colOff>2</xdr:colOff>
				<xdr:row>3</xdr:row>
				<xdr:rowOff>4</xdr:rowOff>
			</xdr:from>
			<xdr:to>
				<xdr:col>5</xdr:col>
				<xdr:colOff>6</xdr:colOff>
				<xdr:row>7</xdr:row>
				<xdr:rowOff>8</xdr:rowOff>
			</xdr:to>
			<xdr:clientData></xdr:clientData>
		</xdr:twoCellAnchor>
		<xdr:absoluteAnchor>
			<xdr:pos x="11" y="22"></xdr:pos>
			<xdr:ext cx="33" cy="44"></xdr:ext>
			<xdr:clientData></xdr:clientData>
		</xdr:absoluteAnchor>
		<xdr:oneCellAnchor>
			<xdr:from>
				<xdr:col>1</xdr:col>
				<xdr:colOff>2</xdr:colOff>
				<xdr:row>3</xdr:row>
				<xdr:rowOff>4</xdr:rowOff>
			</xdr:from>
			<xdr:ext cx="5" cy="6"></xdr:ext>
			<xdr:clientData></xdr:clientData>
		</xdr:oneCellAnchor>
		<xdr:absoluteAnchor>
			<xdr:pos x="111" y="222"></xdr:pos>
			<xdr:ext cx="333" cy="444"></xdr:ext>
			<xdr:clientData></xdr:clientData>
		</xdr:absoluteAnchor>
	</xdr:entity>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, &Entity{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "entity",
		},
		Items: &drawing.AnchorList{
			&drawing.AbsoluteAnchor{
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
			},
			&drawing.TwoCellAnchor{
				XMLName: xml.Name{
					Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
					Local: "twoCellAnchor",
				},
				From: drawing.Marker{
					Col:       1,
					Row:       3,
					OffsetCol: dml.Coordinate(2),
					OffsetRow: dml.Coordinate(4),
				},
				To: drawing.Marker{
					Col:       5,
					Row:       7,
					OffsetCol: dml.Coordinate(6),
					OffsetRow: dml.Coordinate(8),
				},
			},
			&drawing.AbsoluteAnchor{
				XMLName: xml.Name{
					Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
					Local: "absoluteAnchor",
				},
				Point: dml.Point2D{
					X: 11,
					Y: 22,
				},
				Size: dml.PositiveSize2D{
					Height: 33,
					Width:  44,
				},
			},
			&drawing.OneCellAnchor{
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
					Height: 5,
					Width:  6,
				},
			},
			&drawing.AbsoluteAnchor{
				XMLName: xml.Name{
					Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
					Local: "absoluteAnchor",
				},
				Point: dml.Point2D{
					X: 111,
					Y: 222,
				},
				Size: dml.PositiveSize2D{
					Height: 333,
					Width:  444,
				},
			},
		},
	}, entity)

	//encode data should be same as original
	encode, err := xml.Marshal(entity)
	require.Nil(t, err)
	require.Equal(t, strings.NewReplacer("xdr:", "", ":xdr", "").Replace(data), string(encode))
}

func TestAnchorList_Add(t *testing.T) {
	anchors := drawing.AnchorList{}
	anchors.Add(1)
	require.Equal(t, drawing.AnchorList{1}, anchors)
	anchors.Add(2)
	require.Equal(t, drawing.AnchorList{1, 2}, anchors)
	anchors.Add(3)
	require.Equal(t, drawing.AnchorList{1, 2, 3}, anchors)
	anchors.Add(1)
	require.Equal(t, drawing.AnchorList{1, 2, 3, 1}, anchors)
}
