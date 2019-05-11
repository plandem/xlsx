package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type noErrorsRule struct {
	baseRule
}

var NoErrors noErrorsRule

func (x noErrorsRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = NoErrors
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeNotContainsErrors,
			Formula: []ml.Formula{`NOT(ISERROR(:cell:))`},
		}
	}
}

func (x noErrorsRule) Styles(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.style = s
	}
}
