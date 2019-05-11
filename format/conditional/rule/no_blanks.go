package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type noBlanksRule struct {
	baseRule
}

var NoBlanks noBlanksRule

func (x noBlanksRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = NoBlanks
		r.rule = &ml.ConditionalRule{
			Type:    primitives.ConditionTypeNotContainsBlanks,
			Formula: []ml.Formula{`LEN(TRIM(:cell:))>0`},
		}
	}
}

func (x noBlanksRule) Styles(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.style = s
	}
}
