package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToRichFont(t *testing.T) {
	style := format.New(
		format.Font.Size(8),
		format.Font.Color("#FF1122"),
	)

	font := toRichFont(style)
	require.IsType(t, &ml.RichFont{}, font)

	require.Equal(t, &ml.RichFont{
		Size:  8,
		Color: &ml.Color{RGB: "FFFF1122"},
	}, font)
}

func TestToRichText(t *testing.T) {
	text, err := toRichText("1", "2", "3", format.New(
		format.Font.Color("#FF3344"),
	))
	require.Nil(t, err)
	require.Equal(t, &ml.StringItem{
		RichText: &[]*ml.RichText{
			{
				Text: "1",
			},
			{
				Text: "2",
			},
			{
				Text: "3",
			},
		},
	}, text)

	text, err = toRichText(format.New(
		format.Font.Color("#FF1122"),
	), "1", format.New(
		format.Font.Size(8),
		format.Font.Color("#FF3344"),
	), "2")

	require.Nil(t, err)
	require.Equal(t, &ml.StringItem{
		RichText: &[]*ml.RichText{
			{
				Text: "1",
				Font: &ml.RichFont{
					Color: &ml.Color{RGB:"FFFF1122"},
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

	text, err = toRichText("1", "2", "3", format.New(
		format.Font.Color("#FF3344"),
	),format.New(
		format.Font.Color("#FF3344"),
	), "4")

	require.NotNil(t, err)
	require.Nil(t, text)
}

func TestFromRichText(t *testing.T) {
	text, err := toRichText("1", "2", "3", format.New(
		format.Font.Color("#FF3344"),
	))
	require.Nil(t, err)
	require.Equal(t, "123", fromRichText(text))

	text, err = toRichText(format.New(
		format.Font.Color("#FF1122"),
	), "1", format.New(
		format.Font.Size(8),
		format.Font.Color("#FF3344"),
	), "2")

	require.Nil(t, err)
	require.Equal(t, "12", fromRichText(text))

	require.Equal(t, "", fromRichText(nil))
}
