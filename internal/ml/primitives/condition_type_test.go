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

func TestConditionType(t *testing.T) {
	type Entity struct {
		Attribute primitives.ConditionType `xml:"attribute,attr"`
	}

	list := map[primitives.ConditionType]string{
		primitives.ConditionType(0):               "",
		primitives.ConditionTypeExpression:        primitives.ConditionTypeExpression.String(),
		primitives.ConditionTypeCellIs:            primitives.ConditionTypeCellIs.String(),
		primitives.ConditionTypeColorScale:        primitives.ConditionTypeColorScale.String(),
		primitives.ConditionTypeDataBar:           primitives.ConditionTypeDataBar.String(),
		primitives.ConditionTypeIconSet:           primitives.ConditionTypeIconSet.String(),
		primitives.ConditionTypeTop10:             primitives.ConditionTypeTop10.String(),
		primitives.ConditionTypeUniqueValues:      primitives.ConditionTypeUniqueValues.String(),
		primitives.ConditionTypeDuplicateValues:   primitives.ConditionTypeDuplicateValues.String(),
		primitives.ConditionTypeContainsText:      primitives.ConditionTypeContainsText.String(),
		primitives.ConditionTypeNotContainsText:   primitives.ConditionTypeNotContainsText.String(),
		primitives.ConditionTypeBeginsWith:        primitives.ConditionTypeBeginsWith.String(),
		primitives.ConditionTypeEndsWith:          primitives.ConditionTypeEndsWith.String(),
		primitives.ConditionTypeContainsBlanks:    primitives.ConditionTypeContainsBlanks.String(),
		primitives.ConditionTypeNotContainsBlanks: primitives.ConditionTypeNotContainsBlanks.String(),
		primitives.ConditionTypeContainsErrors:    primitives.ConditionTypeContainsErrors.String(),
		primitives.ConditionTypeNotContainsErrors: primitives.ConditionTypeNotContainsErrors.String(),
		primitives.ConditionTypeTimePeriod:        primitives.ConditionTypeTimePeriod.String(),
		primitives.ConditionTypeAboveAverage:      primitives.ConditionTypeAboveAverage.String(),
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
