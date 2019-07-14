// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml_test

import (
	"bytes"
	"encoding/xml"
	//vml "github.com/plandem/ooxml/drawing/vml"
	//css "github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestComments(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
		<comments xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main">
			<authors>
				<author>Microsoft Office User</author>
			</authors>
			<commentList>
				<comment ref="C6" authorId="0" shapeId="0">
					<text>
						<r>
							<rPr>
								<rFont val="Tahoma"></rFont>
								<family val="2"></family>
								<color rgb="FF000000"></color>
								<sz val="10"></sz>
							</rPr>
							<t>My Comment1</t>
						</r>
					</text>
				</comment>
				<comment ref="C7" authorId="0">
					<text>
						<t>My Comment2</t>
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
					RichText: []*ml.RichText{
						{
							Text: "My Comment1",
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
			{
				Ref: primitives.Ref("C7").ToBounds(),
				AuthorID: 0,
				Text: &ml.StringItem{
					Text: "My Comment2",
				},
			},
		},
	}, comments)

	//encode data should be same as original
	encode, err := xml.Marshal(comments)
	require.Nil(t, err)
	require.Equal(t, data, string(encode))
}
