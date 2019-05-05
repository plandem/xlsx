package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFontCharset(t *testing.T) {
	type Element struct {
		Property primitives.FontCharsetType `xml:"property,omitempty"`
	}

	list := map[string]styles.FontCharsetType{
		//"0":   format.FontCharsetANSI,
		"1":   styles.FontCharsetDEFAULT,
		"2":   styles.FontCharsetSYMBOL,
		"77":  styles.FontCharsetMAC,
		"128": styles.FontCharsetSHIFTJIS,
		"129": styles.FontCharsetHANGUL,
		"130": styles.FontCharsetJOHAB,
		"134": styles.FontCharsetGB2312,
		"136": styles.FontCharsetCHINESEBIG5,
		"161": styles.FontCharsetGREEK,
		"162": styles.FontCharsetTURKISH,
		"163": styles.FontCharsetVIETNAMESE,
		"177": styles.FontCharsetHEBREW,
		"178": styles.FontCharsetARABIC,
		"186": styles.FontCharsetBALTIC,
		"204": styles.FontCharsetRUSSIAN,
		"222": styles.FontCharsetTHAI,
		"238": styles.FontCharsetEASTEUROPE,
		"255": styles.FontCharsetOEM,
		"25":  styles.FontCharsetType(25),
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
