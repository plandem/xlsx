// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditionOperator(t *testing.T) {
	type Entity struct {
		Attribute primitives.ConditionOperatorType `xml:"attribute,attr"`
	}

	list := map[primitives.ConditionOperatorType]string{
		primitives.ConditionOperatorType(0):            "",
		primitives.ConditionOperatorLessThan:           primitives.ConditionOperatorLessThan.String(),
		primitives.ConditionOperatorLessThanOrEqual:    primitives.ConditionOperatorLessThanOrEqual.String(),
		primitives.ConditionOperatorEqual:              primitives.ConditionOperatorEqual.String(),
		primitives.ConditionOperatorNotEqual:           primitives.ConditionOperatorNotEqual.String(),
		primitives.ConditionOperatorGreaterThanOrEqual: primitives.ConditionOperatorGreaterThanOrEqual.String(),
		primitives.ConditionOperatorGreaterThan:        primitives.ConditionOperatorGreaterThan.String(),
		primitives.ConditionOperatorBetween:            primitives.ConditionOperatorBetween.String(),
		primitives.ConditionOperatorNotBetween:         primitives.ConditionOperatorNotBetween.String(),
		primitives.ConditionOperatorContainsText:       primitives.ConditionOperatorContainsText.String(),
		primitives.ConditionOperatorNotContains:        primitives.ConditionOperatorNotContains.String(),
		primitives.ConditionOperatorBeginsWith:         primitives.ConditionOperatorBeginsWith.String(),
		primitives.ConditionOperatorEndsWith:           primitives.ConditionOperatorEndsWith.String(),
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
