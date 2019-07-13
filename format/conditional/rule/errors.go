// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type errorsRule struct {
	baseRule
}

//Errors is helper object to set specific options for rule
var Errors errorsRule

func (x errorsRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Errors
		r.rule = &ml.ConditionalRule{
			Type:    primitives.ConditionTypeContainsErrors,
			Formula: []ml.Formula{`ISERROR(:cell:)`},
		}
	}
}

func (x errorsRule) Styles(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.style = s
	}
}
