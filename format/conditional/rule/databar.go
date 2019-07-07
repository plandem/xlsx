package rule

import (
	"fmt"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type dataBarRule struct {
	baseRule
}

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
				MinLength: 10,
				MaxLength: 90,
			},
		}
	}
}

func (x dataBarRule) Default(r *Info) {
	x.initIfRequired(r)
}

func (x dataBarRule) setValue(r *Info, idx int, value string, settings []interface{}) {
	x.initIfRequired(r)
	r.rule.DataBar.Values[idx].Value = value

	for _, p := range settings {
		switch pv := p.(type) {
		case primitives.ConditionValueType:
			r.rule.DataBar.Values[idx].Type = pv
		}
	}
}

func (x dataBarRule) Min(value string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, 0, value, settings)
	}
}

func (x dataBarRule) Max(value string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, 1, value, settings)
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
	r.rule.DataBar.ShowValue = primitives.OptionalBool(false)
}

func (x dataBarRule) Validate(r *Info) error {
	if !r.rule.DataBar.Values[0].Type.IsAllowed(
		ValueTypeLowest,
		ValueTypeNumber,
		ValueTypePercent,
		ValueTypeFormula,
		ValueTypePercentile,
	) {
		return fmt.Errorf("dataBar: Not allowed type '%s' for min value", r.rule.DataBar.Values[0].Type)
	}

	if !r.rule.DataBar.Values[1].Type.IsAllowed(
		ValueTypeNumber,
		ValueTypePercent,
		ValueTypeFormula,
		ValueTypePercentile,
		ValueTypeHighest,
	) {
		return fmt.Errorf("dataBar: Not allowed type '%s' for max value", r.rule.DataBar.Values[1].Type)
	}

	return nil
}
