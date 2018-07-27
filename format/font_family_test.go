package format_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml/styles"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFontFamily(t *testing.T) {
	type Element struct {
		Property styles.FontFamilyType `xml:"property,omitempty"`
	}

	list := map[string]styles.FontFamilyType{
		"":   styles.FontFamilyType(0),
		"1":  format.FontFamilyRoman,
		"2":  format.FontFamilySwiss,
		"3":  format.FontFamilyModern,
		"4":  format.FontFamilyScript,
		"5":  format.FontFamilyDecorative,
		"6":  styles.FontFamilyType(6), //officially 6-14 - reserved
		"25": styles.FontFamilyType(25),
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if s == "" {
				require.Equal(tt, `<Element></Element>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Element><property val="%s"></property></Element>`, s), string(encoded))
			}

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
