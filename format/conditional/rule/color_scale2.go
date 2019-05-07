package rule

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type colorScale2Rule byte
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

func (x colorScale2Rule) Min(min string, rgb string) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.ColorScale.Colors[0] = color.New(rgb)
		r.rule.ColorScale.Values[0].Value = min
	}
}

func (x colorScale2Rule) Max(max string, rgb string) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.ColorScale.Colors[1] = color.New(rgb)
		r.rule.ColorScale.Values[1].Value = max
	}
}

func (x colorScale2Rule) Validate(r *Info) error {
	return nil
}
