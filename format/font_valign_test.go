package format_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFontVAlign(t *testing.T) {
	type Element struct {
		Property ml.FontVAlignType `xml:"property,omitempty"`
	}

	list := map[string]ml.FontVAlignType{
		"baseline":    format.FontVAlignBaseline,
		"superscript": format.FontVAlignSuperscript,
		"subscript":   format.FontVAlignSubscript,
		"align-a":     ml.FontVAlignType("align-a"),
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
