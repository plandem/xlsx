package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type errorsRule struct {
	baseRule
}

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
