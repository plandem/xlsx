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
	"strconv"
	"testing"
)

func TestFontFamily(t *testing.T) {
	type Element struct {
		Property primitives.FontFamilyType `xml:"property,omitempty"`
	}

	list := map[primitives.FontFamilyType]string{
		primitives.FontFamilyType(0):  "",
		styles.FontFamilyRoman:        strconv.Itoa(int(styles.FontFamilyRoman)),
		styles.FontFamilySwiss:        strconv.Itoa(int(styles.FontFamilySwiss)),
		styles.FontFamilyModern:       strconv.Itoa(int(styles.FontFamilyModern)),
		styles.FontFamilyScript:       strconv.Itoa(int(styles.FontFamilyScript)),
		styles.FontFamilyDecorative:   strconv.Itoa(int(styles.FontFamilyDecorative)),
		primitives.FontFamilyType(6):  strconv.Itoa(6), //officially 6-14 - reserved
		primitives.FontFamilyType(25): strconv.Itoa(25),
	}

	for v, s := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if v == 0 {
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
