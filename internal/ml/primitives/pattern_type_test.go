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

func TestFillPattern(t *testing.T) {
	type Entity struct {
		Attribute primitives.PatternType `xml:"attribute,attr"`
	}

	list := map[primitives.PatternType]string{
		primitives.PatternType(0):         "",
		styles.PatternTypeNone:            styles.PatternTypeNone.String(),
		styles.PatternTypeSolid:           styles.PatternTypeSolid.String(),
		styles.PatternTypeMediumGray:      styles.PatternTypeMediumGray.String(),
		styles.PatternTypeDarkGray:        styles.PatternTypeDarkGray.String(),
		styles.PatternTypeLightGray:       styles.PatternTypeLightGray.String(),
		styles.PatternTypeDarkHorizontal:  styles.PatternTypeDarkHorizontal.String(),
		styles.PatternTypeDarkVertical:    styles.PatternTypeDarkVertical.String(),
		styles.PatternTypeDarkDown:        styles.PatternTypeDarkDown.String(),
		styles.PatternTypeDarkUp:          styles.PatternTypeDarkUp.String(),
		styles.PatternTypeDarkGrid:        styles.PatternTypeDarkGrid.String(),
		styles.PatternTypeDarkTrellis:     styles.PatternTypeDarkTrellis.String(),
		styles.PatternTypeLightHorizontal: styles.PatternTypeLightHorizontal.String(),
		styles.PatternTypeLightVertical:   styles.PatternTypeLightVertical.String(),
		styles.PatternTypeLightDown:       styles.PatternTypeLightDown.String(),
		styles.PatternTypeLightUp:         styles.PatternTypeLightUp.String(),
		styles.PatternTypeLightGrid:       styles.PatternTypeLightGrid.String(),
		styles.PatternTypeLightTrellis:    styles.PatternTypeLightTrellis.String(),
		styles.PatternTypeGray125:         styles.PatternTypeGray125.String(),
		styles.PatternTypeGray0625:        styles.PatternTypeGray0625.String(),
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
