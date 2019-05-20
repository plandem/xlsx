package ml_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/vml"
	"github.com/plandem/ooxml/vml/css"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestComments(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<comments xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main">
			<authors>
				<author>Microsoft Office User</author>
			</authors>
			<commentList>
				<comment ref="C6" authorId="0" shapeId="0">
            		<text>
                		<r>
							<rPr>
								<sz val="10"/>
								<color rgb="FF000000"/>
								<rFont val="Tahoma"/>
								<family val="2"/>
							</rPr>
							<t>My Comment</t>
						</r>
					</text>
				</comment>
			</commentList>
		</comments>
	`)

	shapeID := 0
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	comments := &ml.Comments{}
	err := decoder.DecodeElement(comments, nil)
	require.Nil(t, err)
	require.Equal(t, &ml.Comments{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main",
			Local: "comments",
		},
		Authors: []primitives.Text{
			"Microsoft Office User",
		},
		CommentList: []*ml.Comment{
			{
				Ref:      primitives.Ref("C6").ToBounds(),
				AuthorID: 0,
				ShapeID:  &shapeID,
				Text: &ml.StringItem{
					Text: "",
					RichText: &[]*ml.RichText{
						{
							Text: "My Comment",
							Font: &ml.RichFont{
								Size:   10,
								Color:  &ml.Color{RGB: "FF000000"},
								Name:   "Tahoma",
								Family: 2,
							},
						},
					},
				},
			},
		},
	}, comments)
}

func TestDrawings(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
		<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel">
    		<o:shapelayout v:ext="edit">
        		<o:idmap v:ext="edit" data="1"/>
    		</o:shapelayout>
    		<v:shapetype id="_x0000_t202" coordsize="21600,21600" o:spt="202" path="m,l,21600r21600,l21600,xe">
	        	<v:stroke joinstyle="miter"/>
	   	     	<v:path gradientshapeok="t" o:connecttype="rect"/>
	    	</v:shapetype>
	    	<v:shape id="_x0000_s1025" type="#_x0000_t202" style="position:absolute;margin-left:59.25pt;margin-top:1.5pt;width:96pt;height:55.5pt;z-index:1;visibility:hidden" fillcolor="#ffffe1" o:insetmode="auto" filled="true">
	        	<v:fill color="red" color2="red"/>
        		<v:shadow on="t" color="black" obscured="t"/>
        		<v:path o:connecttype="none"/>
    	    	<v:textbox style="mso-direction-alt:auto">
	            	<div style="text-align:right"></div>
        		</v:textbox>
        		<x:ClientData ObjectType="Note">
            		<x:MoveWithCells/>
            		<x:SizeWithCells/>
            		<x:Anchor>1, 15, 0, 2, 3, 15, 3, 16</x:Anchor>
            		<x:AutoFill>False</x:AutoFill>
            		<x:Row>0</x:Row>
            		<x:Column>0</x:Column>
        		</x:ClientData>
    		</v:shape>
		</xml>
	`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	drawings := &vml.Excel{}
	err := decoder.DecodeElement(drawings, nil)
	require.Nil(t, err)

	//check decoded data
	require.Equal(t, "_x0000_s1025", drawings.Shape[0].Attrs["id"])
	require.Equal(t, "#_x0000_t202", drawings.Shape[0].Attrs["type"])
	require.Equal(t, "#ffffe1", drawings.Shape[0].Attrs["fillcolor"])
	require.Equal(t, "auto", drawings.Shape[0].Attrs["o:insetmode"])
	require.Equal(t, "true", drawings.Shape[0].Attrs["filled"])

	//update shape with custom types
	drawings.Shape[0].Attrs["style"] = &css.Style{
		Width:    100,
		Height:   200,
		Position: css.PositionAbsolute,
		Visible:  css.VisibilityHidden,
		ZIndex:   1,
	}

	/*
	ClientData: &xClientData{
		ObjectType: "Note",
		Anchor: fmt.Sprintf(
			"%d, 23, %d, 0, %d, %d, %d, 5",
			1+yAxis, 1+xAxis, 2+yAxis+lineCount, colCount+yAxis, 2+xAxis+lineCount),
		AutoFill: "True",
		Row:      xAxis,
		Column:   yAxis,
	},

	[col_start, x1, row_start, y1, col_end, x2, row_end, y2]
	*/
	drawings.Shape[0].Attrs["fillcolor"] = "#ff00ff"
	drawings.Shape[0].Nested[4] = ml.ClientData{
		MoveWithCells: true,
		SizeWithCells: true,
		Anchor:        "1, 15, 0, 2, 3, 15, 3, 16",
		AutoFill:      false, //'false' value will be omitted
		Row:           10,
		Column:        1,
		Type:          "Note",
	}

	encoded, err := xml.Marshal(&drawings)
	require.Nil(t, err)

	drawings2 := &vml.Excel{}
	decoder = xml.NewDecoder(bytes.NewReader([]byte(encoded)))
	err = decoder.DecodeElement(drawings2, nil)
	require.Nil(t, err)

	//check decoded data
	require.Equal(t, "#ff00ff", drawings2.Shape[0].Attrs["fillcolor"])
	require.Equal(t, "position:absolute;width:100px;height:200px;z-index:1;visibility:hidden", drawings2.Shape[0].Attrs["style"])
	require.Equal(t, &vml.Reserved{
		Name: xml.Name{
			Local: "x:ClientData",
		},
		Attrs: map[string]interface{}{
			"ObjectType": "Note",
		},
		Nested: []interface{}{
			&vml.Reserved{
				Name: xml.Name{
					Local: "x:MoveWithCells",
				},
				InnerXML: "true",
				Attrs:    map[string]interface{}{},
			},
			&vml.Reserved{
				Name: xml.Name{
					Local: "x:SizeWithCells",
				},
				InnerXML: "true",
				Attrs:    map[string]interface{}{},
			},
			&vml.Reserved{
				Name: xml.Name{
					Local: "x:Anchor",
				},
				InnerXML: "1, 15, 0, 2, 3, 15, 3, 16",
				Attrs:    map[string]interface{}{},
			},
			&vml.Reserved{
				Name: xml.Name{
					Local: "x:Row",
				},
				InnerXML: "10",
				Attrs:    map[string]interface{}{},
			},
			&vml.Reserved{
				Name: xml.Name{
					Local: "x:Column",
				},
				InnerXML: "1",
				Attrs:    map[string]interface{}{},
				//Nested:[]interface {}{},
			},
		},
	}, drawings2.Shape[0].Nested[4])
}
