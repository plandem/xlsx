// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type duplicateRule struct {
	baseRule
}

var Duplicate duplicateRule

func (x duplicateRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Duplicate
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeDuplicateValues,
		}
	}
}

func (x duplicateRule) Styles(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.style = s
	}
}
