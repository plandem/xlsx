// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type averageRule struct {
	baseRule
}

//Average is helper object to set specific options for rule
var Average averageRule

func (x averageRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Average
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeAboveAverage,
		}
	}
}

func (x averageRule) Above(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.style = s
	}
}

func (x averageRule) EqualOrAbove(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.EqualAverage = true
		r.style = s
	}
}

func (x averageRule) Below(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.AboveAverage = primitives.OptionalBool(false)
		r.style = s
	}
}

func (x averageRule) EqualOrBelow(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.AboveAverage = primitives.OptionalBool(false)
		r.rule.EqualAverage = true
		r.style = s
	}
}

func (x averageRule) StdDevAbove(n int, s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.style = s
		r.rule.StdDev = n
	}
}

func (x averageRule) StdDevBelow(n int, s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.AboveAverage = primitives.OptionalBool(false)
		r.style = s
		r.rule.StdDev = n
	}
}
