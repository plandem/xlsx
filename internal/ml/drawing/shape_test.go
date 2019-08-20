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

func TestShape(t *testing.T) {
	type Entity struct {
		XMLName xml.Name       `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing entity"`
		DMLName dml.Name       `xml:",attr"`
		Shape   *drawing.Shape `xml:"sp"`
	}

	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
	<xdr:entity xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
		<xdr:sp macro="" textlink="">
			<xdr:nvSpPr>
				<xdr:cNvPr id="2" name="Rectangle 1"></xdr:cNvPr>
				<xdr:cNvSpPr></xdr:cNvSpPr>
			</xdr:nvSpPr>
			<xdr:spPr>
				<a:xfrm>
					<a:off x="657225" y="1009650"></a:off>
					<a:ext cx="5322093" cy="561974"></a:ext>
				</a:xfrm>
				<a:prstGeom prst="rect">
					<a:avLst></a:avLst>
				</a:prstGeom>
			</xdr:spPr>
			<xdr:style>
				<a:lnRef idx="2">
					<a:schemeClr val="accent6"></a:schemeClr>
				</a:lnRef>
				<a:fillRef idx="1">
					<a:schemeClr val="lt1"></a:schemeClr>
				</a:fillRef>
				<a:effectRef idx="0">
					<a:schemeClr val="accent6"></a:schemeClr>
				</a:effectRef>
				<a:fontRef idx="minor">
					<a:schemeClr val="dk1"></a:schemeClr>
				</a:fontRef>
			</xdr:style>
			<xdr:txBody>
				<a:bodyPr rtlCol="0" anchor="t"></a:bodyPr>
				<a:lstStyle></a:lstStyle>
				<a:p>
					<a:pPr algn="l"></a:pPr>
					<a:r>
						<a:rPr lang="en-US" sz="1100" b="0" baseline="0">
							<a:solidFill>
								<a:sysClr val="windowText" lastClr="000000"></a:sysClr>
							</a:solidFill>
						</a:rPr>
						<a:t>All results within normal limit</a:t>
					</a:r>
					<a:endParaRPr lang="en-US" sz="1200" b="0">
						<a:solidFill>
							<a:sysClr val="windowText" lastClr="000000"></a:sysClr>
						</a:solidFill>
					</a:endParaRPr>
				</a:p>
			</xdr:txBody>
		</xdr:sp>
	</xdr:entity>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	object := &drawing.Shape{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "sp",
		},
		NonVisual: &drawing.ShapeNonVisual{
			CommonProperties: &dml.NonVisualCommonProperties{
				ID:   2,
				Name: "Rectangle 1",
			},
			ShapeProperties: &dml.NonVisualShapeProperties{},
		},
		ReservedAttributes: ml.ReservedAttributes{
			Attrs: []xml.Attr{
				{
					Name: xml.Name{
						Local: "macro",
					},
				},
				{
					Name: xml.Name{
						Local: "textlink",
					},
				},
			},
		},
	}

	//add shape
	object.Shape = &dml.Shape{
		Transform: &dml.Transform2D{
			Offset: &dml.Point2D{
				X: 657225,
				Y: 1009650,
			},
			Size: &dml.PositiveSize2D{
				Height: 5322093,
				Width:  561974,
			},
		},
		Geometry: dml.Geometry{
			Preset: &dml.PresetGeometry2D{
				Type: "rect",
				ReservedElements: ml.ReservedElements{
					Nodes: []ml.Reserved{
						{
							XMLName: xml.Name{
								Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
								Local: "avLst",
							},
						},
					},
				},
			},
		},
	}

	//add style
	object.Style = &dml.ShapeStyle{
		ReservedElements: ml.ReservedElements{
			Nodes: []ml.Reserved{
				{
					XMLName: xml.Name{
						Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
						Local: "lnRef",
					},
					InnerXML: `<a:schemeClr val="accent6"></a:schemeClr>`,
					ReservedAttributes: ml.ReservedAttributes{
						Attrs: []xml.Attr{
							{
								Name: xml.Name{
									Local: "idx",
								},
								Value: "2",
							},
						},
					},
				},
				{
					XMLName: xml.Name{
						Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
						Local: "fillRef",
					},
					InnerXML: `<a:schemeClr val="lt1"></a:schemeClr>`,
					ReservedAttributes: ml.ReservedAttributes{
						Attrs: []xml.Attr{
							{
								Name: xml.Name{
									Local: "idx",
								},
								Value: "1",
							},
						},
					},
				},
				{
					XMLName: xml.Name{
						Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
						Local: "effectRef",
					},
					InnerXML: `<a:schemeClr val="accent6"></a:schemeClr>`,
					ReservedAttributes: ml.ReservedAttributes{
						Attrs: []xml.Attr{
							{
								Name: xml.Name{
									Local: "idx",
								},
								Value: "0",
							},
						},
					},
				},
				{
					XMLName: xml.Name{
						Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
						Local: "fontRef",
					},
					InnerXML: `<a:schemeClr val="dk1"></a:schemeClr>`,
					ReservedAttributes: ml.ReservedAttributes{
						Attrs: []xml.Attr{
							{
								Name: xml.Name{
									Local: "idx",
								},
								Value: "minor",
							},
						},
					},
				},
			},
		},
	}

	//add text
	object.Text = &dml.TextBody{
		ReservedElements: ml.ReservedElements{
			Nodes: []ml.Reserved{
				{
					XMLName: xml.Name{
						Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
						Local: "bodyPr",
					},
					ReservedAttributes: ml.ReservedAttributes{
						Attrs: []xml.Attr{
							{
								Name: xml.Name{
									Local: "rtlCol",
								},
								Value: "0",
							},
							{
								Name: xml.Name{
									Local: "anchor",
								},
								Value: "t",
							},
						},
					},
				},
				{
					XMLName: xml.Name{
						Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
						Local: "lstStyle",
					},
				},
				{
					XMLName: xml.Name{
						Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
						Local: "p",
					},
					InnerXML: `<a:pPr algn="l"></a:pPr><a:r><a:rPr lang="en-US" sz="1100" b="0" baseline="0"><a:solidFill><a:sysClr val="windowText" lastClr="000000"></a:sysClr></a:solidFill></a:rPr><a:t>All results within normal limit</a:t></a:r><a:endParaRPr lang="en-US" sz="1200" b="0"><a:solidFill><a:sysClr val="windowText" lastClr="000000"></a:sysClr></a:solidFill></a:endParaRPr>`,
				},
			},
		},
	}

	require.Equal(t, &Entity{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "entity",
		},
		Shape: object,
	}, entity)

	//encode data should be same as original
	encode, err := xml.Marshal(entity)
	require.Nil(t, err)
	require.Equal(t, strings.NewReplacer("xdr:", "", ":xdr", "").Replace(data), string(encode))
}
