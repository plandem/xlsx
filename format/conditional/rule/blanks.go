// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type blanksRule struct {
	baseRule
}

//Blanks is helper object to set specific options for rule
var Blanks blanksRule

func (x blanksRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Blanks
		r.rule = &ml.ConditionalRule{
			Type:    primitives.ConditionTypeContainsBlanks,
			Formula: []ml.Formula{`LEN(TRIM(:cell:))=0`},
		}
	}
}

func (x blanksRule) Styles(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.style = s
	}
}
