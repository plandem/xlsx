package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFontCharset(t *testing.T) {
	type Element struct {
		Property primitives.FontCharsetType `xml:"property,omitempty"`
	}

	list := map[string]format.FontCharsetType{
		//"0":   format.FontCharsetANSI,
		"1":   format.FontCharsetDEFAULT,
		"2":   format.FontCharsetSYMBOL,
		"77":  format.FontCharsetMAC,
		"128": format.FontCharsetSHIFTJIS,
		"129": format.FontCharsetHANGUL,
		"130": format.FontCharsetJOHAB,
		"134": format.FontCharsetGB2312,
		"136": format.FontCharsetCHINESEBIG5,
		"161": format.FontCharsetGREEK,
		"162": format.FontCharsetTURKISH,
		"163": format.FontCharsetVIETNAMESE,
		"177": format.FontCharsetHEBREW,
		"178": format.FontCharsetARABIC,
		"186": format.FontCharsetBALTIC,
		"204": format.FontCharsetRUSSIAN,
		"222": format.FontCharsetTHAI,
		"238": format.FontCharsetEASTEUROPE,
		"255": format.FontCharsetOEM,
		"25":  format.FontCharsetType(25),
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: primitives.FontCharsetType(v)}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Element><property val="%s"></property></Element>`, s), string(encoded))

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
