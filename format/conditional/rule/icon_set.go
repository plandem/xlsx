package rule

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type iconSetRule byte
var IconSet iconSetRule

func (x iconSetRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = IconSet
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeIconSet,
			IconSet: &ml.IconSet{
				ShowValue: true,
				Percent:   true,
				Type:      IconSetType3TrafficLights1,
			},
		}
	}
}

func (x iconSetRule) Default(r *Info) {
	x.initIfRequired(r)
}

func (x iconSetRule) Type(t primitives.IconSetType) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.IconSet.Type = t
	}
}

func (x iconSetRule) ReverseIcons(r *Info) {
	x.initIfRequired(r)
	r.rule.IconSet.Reverse = true
}

func (x iconSetRule) IconsOnly(r *Info) {
	x.initIfRequired(r)
	r.rule.IconSet.ShowValue = false
}

func (x iconSetRule) Value(value string, valueType primitives.ConditionValueType, criteria primitives.ConditionOperatorType) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		//TODO: create!!!
	}
}

func (x iconSetRule) Validate(r *Info) error {
	return nil
}
