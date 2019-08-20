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

func TestPicture(t *testing.T) {
	type Entity struct {
		XMLName xml.Name         `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing entity"`
		DMLName dml.Name         `xml:",attr"`
		Picture *drawing.Picture `xml:"pic"`
	}

	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
	<xdr:entity xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
		<xdr:pic>
			<xdr:nvPicPr>
				<xdr:cNvPr id="2" name="Picture 1" descr="python.png"></xdr:cNvPr>
				<xdr:cNvPicPr>
					<a:picLocks noChangeAspect="true"></a:picLocks>
				</xdr:cNvPicPr>
			</xdr:nvPicPr>
			<xdr:blipFill>
				<a:blip xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:embed="rId1"></a:blip>
				<a:stretch>
					<a:fillRect></a:fillRect>
				</a:stretch>
			</xdr:blipFill>
			<xdr:spPr>
				<a:xfrm>
					<a:off x="2047875" y="190500"></a:off>
					<a:ext cx="1333333" cy="1733333"></a:ext>
				</a:xfrm>
				<a:prstGeom prst="rect">
					<a:avLst></a:avLst>
				</a:prstGeom>
			</xdr:spPr>
		</xdr:pic>
	</xdr:entity>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	object := &drawing.Picture{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "pic",
		},
		NonVisual: &drawing.PictureNonVisual{
			CommonProperties: &dml.NonVisualCommonProperties{
				ID:          2,
				Name:        "Picture 1",
				Description: "python.png",
			},
			PictureProperties: &dml.NonVisualPictureProperties{
				Locking: &dml.PictureLocking{},
			},
		},
	}

	object.NonVisual.PictureProperties.Locking.NoChangeAspect = true

	//add blip
	object.BlipFill = &dml.BlipFill{
		Blip: &dml.Blip{
			Embed: "rId1",
		},
		ReservedElements: ml.ReservedElements{
			Nodes: []ml.Reserved{
				{
					XMLName: xml.Name{
						Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
						Local: "stretch",
					},
					InnerXML: "<a:fillRect></a:fillRect>",
				},
			},
		},
	}

	//add shape
	object.Shape = &dml.Shape{
		Transform: &dml.Transform2D{
			Offset: &dml.Point2D{
				X: 2047875,
				Y: 190500,
			},
			Size: &dml.PositiveSize2D{
				Height: 1333333,
				Width:  1733333,
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

	//compare data
	require.Equal(t, &Entity{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing",
			Local: "entity",
		},
		Picture: object,
	}, entity)

	//encode data should be same as original
	encode, err := xml.Marshal(entity)
	require.Nil(t, err)
	require.Equal(t, strings.NewReplacer("xdr:", "", ":xdr", "").Replace(data), string(encode))
}
