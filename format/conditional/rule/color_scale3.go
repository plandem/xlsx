package rule

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type colorScale3Rule byte
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
						Type: ValueTypePercentile,
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

func (x colorScale3Rule) Min(min string, rgb string) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.ColorScale.Colors[0] = color.New(rgb)
		r.rule.ColorScale.Values[0].Value = min
	}
}

func (x colorScale3Rule) Mid(percentile string, rgb string) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.ColorScale.Colors[1] = color.New(rgb)
		r.rule.ColorScale.Values[1].Value = percentile
	}
}

func (x colorScale3Rule) Max(max string, rgb string) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.ColorScale.Colors[2] = color.New(rgb)
		r.rule.ColorScale.Values[2].Value = max
	}
}

func (x colorScale3Rule) Validate(r *Info) error {
	return nil
}
