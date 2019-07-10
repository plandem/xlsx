// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/plandem/xlsx/types/options"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVisibility(t *testing.T) {
	type Entity struct {
		Attribute primitives.VisibilityType `xml:"attribute,attr"`
	}

	list := map[primitives.VisibilityType]string{
		primitives.VisibilityType(0): "",
		options.Visible:              options.Visible.String(),
		options.Hidden:               options.Hidden.String(),
		options.VeryHidden:           options.VeryHidden.String(),
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
