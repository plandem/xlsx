package rule

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type bottomRule struct {
	baseRule
}

var Bottom bottomRule

func (x bottomRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Bottom
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeTop10,
			Rank: 10,
			Bottom: true,
		}
	}
}

func (x bottomRule) Default(r *Info) {
	x.initIfRequired(r)
}

func (x bottomRule) Value(rank uint, settings ...interface{}) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.Rank = rank
		for _, p := range settings {
			switch pv := p.(type) {
			case string:
				if pv == "%" {
					r.rule.Percent = true
				}
			case *styles.Info:
				r.style = pv
			}
		}
	}
}

func (x bottomRule) Validate(r *Info) error {
	if r.rule.Percent {
		if r.rule.Rank < 1 || r.rule.Rank > 100 {
			return errors.New(fmt.Sprintf("bottom: value(%d) should be between (1 - 100)", r.rule.Rank))
		}
	} else {
		if r.rule.Rank < 1 || r.rule.Rank > 1000 {
			return errors.New(fmt.Sprintf("bottom: value(%d) should be between 1 and 1000", r.rule.Rank))
		}
	}

	return nil
}
