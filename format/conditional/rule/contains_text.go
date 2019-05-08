package rule

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type containsTextRule struct {
	baseRule
}

var ContainsText containsTextRule

func (x containsTextRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = ContainsText
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeContainsText,
			Formula: `NOT(ISERROR(SEARCH("%s",%s)))`,
		}
	}
}

func (x containsTextRule) Default(r *Info) {
	x.initIfRequired(r)
}

func (x containsTextRule) Value(rank uint, settings ...interface{}) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		//r.rule.Rank = rank
		//for _, p := range settings {
		//	switch pv := p.(type) {
		//	case string:
		//		if pv == "%" {
		//			r.rule.Percent = true
		//		}
		//	case *styles.Info:
		//		r.style = pv
		//	}
		//}
	}
}

func (x containsTextRule) Validate(r *Info) error {
	//if r.rule.Percent {
	//	if r.rule.Rank < 1 || r.rule.Rank > 100 {
	//		return errors.New(fmt.Sprintf("bottom: value(%d) should be between (1 - 100)", r.rule.Rank))
	//	}
	//} else {
	//	if r.rule.Rank < 1 || r.rule.Rank > 1000 {
	//		return errors.New(fmt.Sprintf("bottom: value(%d) should be between 1 and 1000", r.rule.Rank))
	//	}
	//}

	return nil
}
