// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type topRule struct {
	baseRule
}

//Top is helper object to set specific options for rule
var Top topRule

func (x topRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Top
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeTop10,
			Rank: 10,
		}
	}
}

func (x topRule) Default(r *Info) {
	x.initIfRequired(r)
}

func (x topRule) Value(rank uint, settings ...interface{}) Option {
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

func (x topRule) Validate(r *Info) error {
	if r.rule.Percent {
		if r.rule.Rank < 1 || r.rule.Rank > 100 {
			return fmt.Errorf("top: value(%d) should be between (1 - 100)", r.rule.Rank)
		}
	} else {
		if r.rule.Rank < 1 || r.rule.Rank > 1000 {
			return fmt.Errorf("top: value(%d) should be between 1 and 1000", r.rule.Rank)
		}
	}

	return nil
}
