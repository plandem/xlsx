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

func TestBorderStyle(t *testing.T) {
	type Entity struct {
		Attribute primitives.BorderStyleType `xml:"attribute,attr"`
	}

	list := map[primitives.BorderStyleType]string{
		primitives.BorderStyleType(0):      "",
		styles.BorderStyleNone:             styles.BorderStyleNone.String(),
		styles.BorderStyleThin:             styles.BorderStyleThin.String(),
		styles.BorderStyleMedium:           styles.BorderStyleMedium.String(),
		styles.BorderStyleDashed:           styles.BorderStyleDashed.String(),
		styles.BorderStyleDotted:           styles.BorderStyleDotted.String(),
		styles.BorderStyleThick:            styles.BorderStyleThick.String(),
		styles.BorderStyleDouble:           styles.BorderStyleDouble.String(),
		styles.BorderStyleHair:             styles.BorderStyleHair.String(),
		styles.BorderStyleMediumDashed:     styles.BorderStyleMediumDashed.String(),
		styles.BorderStyleDashDot:          styles.BorderStyleDashDot.String(),
		styles.BorderStyleMediumDashDot:    styles.BorderStyleMediumDashDot.String(),
		styles.BorderStyleDashDotDot:       styles.BorderStyleDashDotDot.String(),
		styles.BorderStyleMediumDashDotDot: styles.BorderStyleMediumDashDotDot.String(),
		styles.BorderStyleSlantDashDot:     styles.BorderStyleSlantDashDot.String(),
	}

	for v, s := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if v == 0 {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
			require.Equal(tt, s, decoded.Attribute.String())
		})
	}
}
