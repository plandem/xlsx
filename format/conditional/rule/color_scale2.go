package rule

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type colorScale2Rule struct {
	baseRule
}

var ColorScale2 colorScale2Rule

func (x colorScale2Rule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = ColorScale2
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeColorScale,
			ColorScale: &ml.ColorScale{
				Values: []*ml.ConditionValue{
					{
						Type: ValueTypeLowest,
					},
					{
						Type: ValueTypeHighest,
					},
				},
				Colors: []*ml.Color{
					color.New("#FF7128"),
					color.New("#FFEF9C"),
				},
			},
		}
	}
}

func (x colorScale2Rule) Default(r *Info) {
	x.initIfRequired(r)
}

func (x colorScale2Rule) setValue(r *Info, idx int, value string, settings []interface{}) {
	x.initIfRequired(r)
	r.rule.ColorScale.Values[idx].Value = value

	for _, p := range settings {
		switch pv := p.(type) {
		case string:
			r.rule.ColorScale.Colors[idx] = color.New(pv)
		case primitives.ConditionValueType:
			r.rule.ColorScale.Values[idx].Type = pv
		}
	}
}

func (x colorScale2Rule) Min(min string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, 0, min, settings)
	}
}

func (x colorScale2Rule) Max(max string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, 1, max, settings)
	}
}

func (x colorScale2Rule) Validate(r *Info) error {
	if !r.rule.ColorScale.Values[0].Type.IsAllowed(
		ValueTypeLowest,
		ValueTypeNumber,
		ValueTypePercent,
		ValueTypeFormula,
		ValueTypePercentile,
	) {
		return errors.New(fmt.Sprintf("colorScale2: Not allowed type '%s' for min value", r.rule.ColorScale.Values[0].Type))
	}

	if !r.rule.ColorScale.Values[1].Type.IsAllowed(
		ValueTypeNumber,
		ValueTypePercent,
		ValueTypeFormula,
		ValueTypePercentile,
		ValueTypeHighest,
	) {
		return errors.New(fmt.Sprintf("colorScale2: Not allowed type '%s' for max value", r.rule.ColorScale.Values[1].Type))
	}

	return nil
}
