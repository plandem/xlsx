package format_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/stretchr/testify/require"
	"testing"
	"github.com/plandem/xlsx/internal/ml/styles"
)

func TestFontUnderline(t *testing.T) {
	type Element struct {
		Property styles.UnderlineType `xml:"property,omitempty"`
	}

	list := map[string]styles.UnderlineType{
		"single":           format.UnderlineTypeSingle,
		"double":           format.UnderlineTypeDouble,
		"singleAccounting": format.UnderlineTypeSingleAccounting,
		"doubleAccounting": format.UnderlineTypeDoubleAccounting,
		"none":             format.UnderlineTypeNone,
		"underline-a":      styles.UnderlineType("underline-a"),
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: v}
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
