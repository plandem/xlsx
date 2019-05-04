package ml_test

import (
	"bytes"
	"encoding/xml"
	ooxml "github.com/plandem/ooxml/ml"
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

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	comments := &ml.Comments{}
	err := decoder.DecodeElement(comments, nil)
	require.Nil(t, err)
	require.Equal(t, &ml.Comments{
		XMLName:ooxml.Name{
			Space:"",
			Local:"",
		},
		Authors:[]primitives.Text{
			"Microsoft Office User",
		},
		CommentList:[]*ml.Comment{
			{
				Ref: primitives.Ref("C6").ToBounds(),
				AuthorID: 0,
				ShapeID: 0,
				Text: &ml.StringItem{
					Text: "",
					RichText: &[]*ml.RichText{
						{
							Text: "My Comment",
							Font: &ml.RichFont{
								Size: 10,
								Color: &ml.Color{RGB:"FF000000"},
								Name: "Tahoma",
								Family: 2,
							},
						},
					},
				},
			},
		},
	}, comments)
}
