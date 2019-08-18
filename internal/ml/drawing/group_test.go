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

func TestGroup(t *testing.T) {
	type Entity struct {
		XMLName xml.Name       `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing entity"`
		DMLName dml.Name       `xml:",attr"`
		Group   *drawing.Group `xml:"grpSp"`
	}

	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
	<xdr:entity xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
		<xdr:grpSp>
			<xdr:nvGrpSpPr>
				<xdr:cNvPr id="2" name="Group 1"></xdr:cNvPr>
				<xdr:cNvGrpSpPr></xdr:cNvGrpSpPr>
			</xdr:nvGrpSpPr>
			<xdr:grpSpPr>
				<a:xfrm>
					<a:off x="2451100" y="889000"></a:off>
					<a:ext cx="2171700" cy="914400"></a:ext>
					<a:chOff x="2451100" y="889000"></a:chOff>
					<a:chExt cx="2171700" cy="914400"></a:chExt>
				</a:xfrm>
			</xdr:grpSpPr>
			<xdr:graphicFrame>
				<xdr:nvGraphicFramePr>
					<xdr:cNvPr id="2" name="Chart 1"></xdr:cNvPr>
					<xdr:cNvGraphicFramePr></xdr:cNvGraphicFramePr>
				</xdr:nvGraphicFramePr>
				<xdr:xfrm>
					<a:off x="1" y="2"></a:off>
					<a:ext cx="3" cy="4"></a:ext>
				</xdr:xfrm>
				<a:graphic>
					<a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/chart">
						<c:chart xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:id="rId1"></c:chart>
					</a:graphicData>
				</a:graphic>
			</xdr:graphicFrame>
			<xdr:graphicFrame>
				<xdr:nvGraphicFramePr>
					<xdr:cNvPr id="3" name="Chart 2"></xdr:cNvPr>
					<xdr:cNvGraphicFramePr></xdr:cNvGraphicFramePr>
				</xdr:nvGraphicFramePr>
				<xdr:xfrm>
					<a:off x="1" y="2"></a:off>
					<a:ext cx="3" cy="4"></a:ext>
				</xdr:xfrm>
				<a:graphic>
					<a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/chart">
						<c:chart xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:id="rId1"></c:chart>
					</a:graphicData>
				</a:graphic>
			</xdr:graphicFrame>
		</xdr:grpSp>
	</xdr:entity>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	object := &drawing.Group{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "grpSp",
		},
		NonVisual: &drawing.GroupNonVisual{
			CommonProperties: &dml.NonVisualCommonProperties{
				ID:   2,
				Name: "Group 1",
			},
			GroupProperties: &dml.NonVisualGroupProperties{},
		},
		Group: &dml.Group{
			Transform: &dml.GroupTransform2D{
				Offset: &dml.Point2D{
					X: 2451100,
					Y: 889000,
				},
				Size: &dml.PositiveSize2D{
					Height: 2171700,
					Width:  914400,
				},
				ChildOffset: &dml.Point2D{
					X: 2451100,
					Y: 889000,
				},
				ChildSize: &dml.PositiveSize2D{
					Height: 2171700,
					Width:  914400,
				},
			},
		},
		Items: &drawing.GroupList{
			&drawing.Frame{
				XMLName: xml.Name{
					Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
					Local: "graphicFrame",
				},
				NonVisual: &drawing.FrameNonVisual{
					CommonProperties: &dml.NonVisualCommonProperties{
						ID:   2,
						Name: "Chart 1",
					},
					FrameProperties: &dml.NonVisualGraphicFrameProperties{},
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
						X: 1,
						Y: 2,
					},
					Size: &dml.PositiveSize2D{
						Height: 3,
						Width:  4,
					},
				},
			},
			&drawing.Frame{
				XMLName: xml.Name{
					Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
					Local: "graphicFrame",
				},
				NonVisual: &drawing.FrameNonVisual{
					CommonProperties: &dml.NonVisualCommonProperties{
						ID:   3,
						Name: "Chart 2",
					},
					FrameProperties: &dml.NonVisualGraphicFrameProperties{},
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
						X: 1,
						Y: 2,
					},
					Size: &dml.PositiveSize2D{
						Height: 3,
						Width:  4,
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
		Group: object,
	}, entity)

	//encode data should be same as original
	encode, err := xml.Marshal(entity)
	require.Nil(t, err)
	require.Equal(t, strings.NewReplacer("xdr:", "", ":xdr", "").Replace(data), string(encode))
}
