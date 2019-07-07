package rule

import (
	"fmt"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type colorScale3Rule struct {
	baseRule
}

var ColorScale3 colorScale3Rule

func (x colorScale3Rule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = ColorScale3
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeColorScale,
			ColorScale: &ml.ColorScale{
				Values: []*ml.ConditionValue{
					{
						Type: ValueTypeLowest,
					},
					{
						Type:  ValueTypePercentile,
						Value: "50",
					},
					{
						Type: ValueTypeHighest,
					},
				},
				Colors: []*ml.Color{
					color.New("#F8696B"),
					color.New("#FFEB84"),
					color.New("#63BE7B"),
				},
			},
		}
	}
}

func (x colorScale3Rule) Default(r *Info) {
	x.initIfRequired(r)
}

func (x colorScale3Rule) setValue(r *Info, idx int, value string, settings []interface{}) {
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

func (x colorScale3Rule) Min(min string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, 0, min, settings)
	}
}

func (x colorScale3Rule) Mid(mid string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, 1, mid, settings)
	}
}

func (x colorScale3Rule) Max(max string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, 2, max, settings)
	}
}

func (x colorScale3Rule) Validate(r *Info) error {
	if !r.rule.ColorScale.Values[0].Type.IsAllowed(
		ValueTypeLowest,
		ValueTypeNumber,
		ValueTypePercent,
		ValueTypeFormula,
		ValueTypePercentile,
	) {
		return fmt.Errorf("colorScale3: Not allowed type '%s' for min value", r.rule.ColorScale.Values[0].Type)
	}

	if !r.rule.ColorScale.Values[2].Type.IsAllowed(
		ValueTypeNumber,
		ValueTypePercent,
		ValueTypeFormula,
		ValueTypePercentile,
		ValueTypeHighest,
	) {
		return fmt.Errorf("colorScale3: Not allowed type '%s' for max value", r.rule.ColorScale.Values[2].Type)
	}
	return nil
}
