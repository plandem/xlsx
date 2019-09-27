// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToRichFont(t *testing.T) {
	style := styles.New(
		styles.Font.Size(8),
		styles.Font.Color("#FF1122"),
	)

	font := styles.ToRichFont(style)
	require.IsType(t, &ml.RichFont{}, font)

	require.Equal(t, &ml.RichFont{
		Size:  8,
		Color: &ml.Color{RGB: "FFFF1122"},
	}, font)
}

func TestToRichText(t *testing.T) {
	s := styles.New(
		styles.Alignment.HAlign(styles.HAlignCenter),
	)

	text, cellStyles, err := toRichText(
		//normal strings
		"1", "2", "3",
		//fmt.Stringer
		types.BoundsFromIndexes(0, 0, 1, 1),
		//custom type with underlying type as string
		types.CellRefFromIndexes(2, 2),
		//cell styles
		s,
	)
	require.Nil(t, err)
	require.Equal(t, &ml.StringItem{
		RichText: []*ml.RichText{
			{
				Text: "1",
			},
			{
				Text: "2",
			},
			{
				Text: "3",
			},
			{
				Text: "A1:B2",
			},
			{
				Text: "C3",
			},
		},
	}, text)
	require.Equal(t, s, cellStyles)

	text, cellStyles, err = toRichText(styles.New(
		styles.Font.Color("#FF1122"),
	), "1", styles.New(
		styles.Font.Size(8),
		styles.Font.Color("#FF3344"),
	), "2")

	require.Nil(t, err)
	require.Equal(t, &ml.StringItem{
		RichText: []*ml.RichText{
			{
				Text: "1",
				Font: &ml.RichFont{
					Color: &ml.Color{RGB: "FFFF1122"},
				},
			},
			{
				Text: "2",
				Font: &ml.RichFont{
					Size:  8,
					Color: &ml.Color{RGB: "FFFF3344"},
				},
			},
		},
	}, text)
	require.Nil(t, cellStyles)

	text, cellStyles, err = toRichText("1", "2", "3", styles.New(
		styles.Font.Color("#FF3344"),
	), styles.New(
		styles.Font.Color("#FF3344"),
	), "4")

	require.NotNil(t, err)
	require.Nil(t, text)
	require.Nil(t, cellStyles)
}

func TestFromRichText(t *testing.T) {
	s := styles.New(
		styles.Alignment.HAlign(styles.HAlignCenter),
	)

	text, cellStyles, err := toRichText("1", "2", "3", s)
	require.Nil(t, err)
	require.Equal(t, "123", fromRichText(text))
	require.Equal(t, s, cellStyles)

	text, cellStyles, err = toRichText(styles.New(
		styles.Font.Color("#FF1122"),
	), "1", styles.New(
		styles.Font.Size(8),
		styles.Font.Color("#FF3344"),
	), "2")

	require.Nil(t, err)
	require.Equal(t, "12", fromRichText(text))
	require.Equal(t, "", fromRichText(nil))
	require.Nil(t, cellStyles)
}
