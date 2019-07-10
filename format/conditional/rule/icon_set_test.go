// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIconSet(t *testing.T) {
	r := New(
		IconSet.Type(IconSetType4Arrows),
		IconSet.ReverseIcons,
		IconSet.IconsOnly,
	)

	require.Equal(t, &Info{
		initialized: true,
		validator:   IconSet,
		rule: &ml.ConditionalRule{
			Type: primitives.ConditionTypeIconSet,
			IconSet: &ml.IconSet{
				Reverse: true,
				Type: IconSetType4Arrows,
				Values: []*ml.ConditionValue{
					{
						Type:  ValueTypePercent,
						Value: "0",
					},
					{
						Type:  ValueTypePercent,
						Value: "25",
					},
					{
						Type:  ValueTypePercent,
						Value: "50",
					},
					{
						Type:  ValueTypePercent,
						Value: "75",
					},
				},
				ShowValue: primitives.OptionalBool(false),
			},
		},
	}, r)
}
