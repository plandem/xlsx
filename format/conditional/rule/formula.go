package rule

import (
	"errors"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type formulaRule struct {
	baseRule
}

var Formula formulaRule

func (x formulaRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Formula
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeExpression,
		}
	}
}

func (x formulaRule) Expression(formula string, settings ...interface{}) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.Formula = []ml.Formula{ml.Formula(x.escape(formula))}

		for _, p := range settings {
			switch pv := p.(type) {
			case *styles.Info:
				r.style = pv
			}
		}
	}
}

func (x formulaRule) Validate(r *Info) error {
	if len(r.rule.Formula) == 0 || len(r.rule.Formula[0]) == 0 {
		return errors.New("formula: no expression to use as rule")
	}

	return nil
}
