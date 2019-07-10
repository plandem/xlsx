// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type timePeriodRule struct {
	baseRule
}

var TimePeriod timePeriodRule

func (x timePeriodRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = TimePeriod
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeTimePeriod,
		}
	}
}

func (x timePeriodRule) setValue(r *Info, s *styles.Info, formula string, t primitives.TimePeriodType) {
	x.initIfRequired(r)
	r.rule.TimePeriod = t
	r.rule.Formula = []ml.Formula{ml.Formula(formula)}
	r.style = s
}

func (x timePeriodRule) Today(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `FLOOR(:cell:,1)=TODAY()`, primitives.TimePeriodToday)
	}
}

func (x timePeriodRule) Yesterday(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `FLOOR(:cell:,1)=TODAY()-1`, primitives.TimePeriodYesterday)
	}
}

func (x timePeriodRule) Tomorrow(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `FLOOR(:cell:,1)=TODAY()+1`, primitives.TimePeriodTomorrow)
	}
}

func (x timePeriodRule) Last7Days(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `AND(TODAY()-FLOOR(:cell:,1)<=6,FLOOR(:cell:,1)<=TODAY())`, primitives.TimePeriodLast7Days)
	}
}

func (x timePeriodRule) ThisWeek(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `AND(TODAY()-ROUNDDOWN(:cell:,0)<=WEEKDAY(TODAY())-1,ROUNDDOWN(:cell:,0)-TODAY()<=7-WEEKDAY(TODAY()))'`, primitives.TimePeriodThisWeek)
	}
}

func (x timePeriodRule) LastWeek(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `AND(TODAY()-ROUNDDOWN(:cell:,0)>=(WEEKDAY(TODAY())),TODAY()-ROUNDDOWN(:cell:,0)<(WEEKDAY(TODAY())+7))'`, primitives.TimePeriodLastWeek)
	}
}

func (x timePeriodRule) NextWeek(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `AND(ROUNDDOWN(:cell:,0)-TODAY()>(7-WEEKDAY(TODAY())),ROUNDDOWN(:cell:,0)-TODAY()<(15-WEEKDAY(TODAY())))'`, primitives.TimePeriodNextWeek)
	}
}

func (x timePeriodRule) ThisMonth(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `AND(MONTH(:cell:)=MONTH(TODAY()),YEAR(:cell:)=YEAR(TODAY()))'`, primitives.TimePeriodThisMonth)
	}
}

func (x timePeriodRule) LastMonth(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `AND(MONTH(:cell:)=MONTH(TODAY())-1,OR(YEAR(:cell:)=YEAR(TODAY()),AND(MONTH(:cell:)=1,YEAR(A1)=YEAR(TODAY())-1)))'`, primitives.TimePeriodLastMonth)
	}
}

func (x timePeriodRule) NextMonth(s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, s, `AND(MONTH(:cell:)=MONTH(TODAY())+1,OR(YEAR(:cell:)=YEAR(TODAY()),AND(MONTH(:cell:)=12,YEAR(:cell:)=YEAR(TODAY())+1)))'`, primitives.TimePeriodNextMonth)
	}
}
