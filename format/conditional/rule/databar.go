package rule

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type dataBarRule byte
var DataBar dataBarRule

func (x dataBarRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = DataBar
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeDataBar,
			DataBar: &ml.DataBar{
				Values: []*ml.ConditionValue{
					{
						Type: ValueTypeLowest,
					},
					{
						Type: ValueTypeHighest,
					},
				},
				Color:     color.New("#638EC6"),
				ShowValue: true,
				MinLength: 10,
				MaxLength: 90,
			},
		}
	}
}

func (x dataBarRule) Default(r *Info) {
	x.initIfRequired(r)
}

func (x dataBarRule) Min(value string, t primitives.ConditionValueType) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.DataBar.Values[0].Type = t
		r.rule.DataBar.Values[0].Value = value
	}
}

func (x dataBarRule) Max(value string, t primitives.ConditionValueType) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.DataBar.Values[1].Type = t
		r.rule.DataBar.Values[1].Value = value
	}
}

func (x dataBarRule) Color(rgb string) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.DataBar.Color = color.New(rgb)
	}
}

func (x dataBarRule) BarOnly(r *Info) {
	x.initIfRequired(r)
	r.rule.DataBar.ShowValue = false
}

func (x dataBarRule) Validate(r *Info) error {
	return nil
}
