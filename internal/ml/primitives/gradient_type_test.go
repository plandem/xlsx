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

func TestFillGradient(t *testing.T) {
	type Entity struct {
		Attribute primitives.GradientType `xml:"attribute,attr"`
	}

	list := map[primitives.GradientType]string{
		styles.GradientTypeLinear: styles.GradientTypeLinear.String(),
		styles.GradientTypePath:   styles.GradientTypePath.String(),
	}

	for v, s := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if v == 0 {
				require.Equal(tt, `<Entity attribute="linear"></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)

			if s == "" {
				require.Equal(tt, "linear", decoded.Attribute.String())
			} else {
				require.Equal(tt, s, decoded.Attribute.String())
			}
		})
	}
}
