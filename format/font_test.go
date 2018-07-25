package format

import (
	"encoding/xml"
	"github.com/plandem/xlsx/format/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFont(t *testing.T) {
	style := New(
		Font.Name("Calibri"),
		Font.Size(10),
		Font.Bold,
		Font.Italic,
		Font.Strikeout,
		Font.Shadow,
		Font.Condense,
		Font.Extend,
		Font.Family(FontFamilyDecorative),
		Font.Color("#FF00FF"),
		Font.Underline(UnderlineTypeSingle),
		Font.VAlign(FontVAlignBaseline),
		Font.Scheme(FontSchemeMinor),
	)

	require.IsType(t, &StyleFormat{}, style)
	require.Equal(t, &StyleFormat{
		key: "dec64c1f2177f8a1995cef78a107ef4e",
		font: ml.Font{
			Name:      "Calibri",
			Bold:      true,
			Italic:    true,
			Strike:    true,
			Shadow:    true,
			Condense:  true,
			Extend:    true,
			Size:      10.0,
			Color:     color.New("FFFF00FF"),
			Family:    FontFamilyDecorative,
			Underline: UnderlineTypeSingle,
			VAlign:    FontVAlignBaseline,
			Scheme:    FontSchemeMinor,
		},
	}, style)
}

func TestFontMarshal(t *testing.T) {
	//0 must be omitted
	style := New(
		Font.Size(0),
		Font.Family(0),
		Font.Charset(0),
	)
	encoded, err := xml.Marshal(&style.font)
	require.Empty(t, err)
	require.Equal(t, `<Font></Font>`, string(encoded))

	//simple version
	style = New(
		Font.Name("Calibri"),
	)
	encoded, _ = xml.Marshal(&style.font)
	require.Equal(t, `<Font><name val="Calibri"></name></Font>`, string(encoded))

	//full version
	style = New(
		Font.Name("Calibri"),
		Font.Size(10),
		Font.Bold,
		Font.Italic,
		Font.Strikeout,
		Font.Shadow,
		Font.Condense,
		Font.Extend,
		Font.Family(FontFamilyDecorative),
		Font.Color("#FF00FF"),
		Font.Underline(UnderlineTypeSingle),
		Font.VAlign(FontVAlignBaseline),
		Font.Scheme(FontSchemeMinor),
	)
	encoded, _ = xml.Marshal(&style.font)
	require.Equal(t, `<Font><name val="Calibri"></name><family val="5"></family><b val="true"></b><i val="true"></i><strike val="true"></strike><shadow val="true"></shadow><condense val="true"></condense><extend val="true"></extend><color indexed="6"></color><sz val="10"></sz><u val="single"></u><vertAlign val="baseline"></vertAlign><scheme val="minor"></scheme></Font>`, string(encoded))
}
