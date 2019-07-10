// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFontFamily(t *testing.T) {
	type Element struct {
		Property primitives.FontFamilyType `xml:"property,omitempty"`
	}

	list := map[string]primitives.FontFamilyType{
		"":   primitives.FontFamilyType(0),
		"1":  styles.FontFamilyRoman,
		"2":  styles.FontFamilySwiss,
		"3":  styles.FontFamilyModern,
		"4":  styles.FontFamilyScript,
		"5":  styles.FontFamilyDecorative,
		"6":  primitives.FontFamilyType(6), //officially 6-14 - reserved
		"25": primitives.FontFamilyType(25),
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
