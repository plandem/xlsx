package rule

import (
	"errors"
	"fmt"
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
				//Percent:   true //default is true, but never used?!
				Type: IconSetType3TrafficLights1,
			},
		}

		x.initValues(r)
	}
}

func (x iconSetRule) Default(r *Info) {
	x.initIfRequired(r)
}

func (x iconSetRule) initValues(r *Info) {
	var total int
	if r.rule.IconSet.Type > __iconSetType3Icons && r.rule.IconSet.Type < __iconSetType4Icons {
		total = 3
	} else if r.rule.IconSet.Type > __iconSetType4Icons && r.rule.IconSet.Type < __iconSetType5Icons {
		total = 4
	} else if r.rule.IconSet.Type > __iconSetType5Icons {
		total = 5
	}

	// The values should start with the highest value and with each subsequent one being lower.
	// The lowest number value in an icon set has properties defined by Excel.
	// Therefore in a n icon set, there are n-1 values.
	// The default value is: (n * 100) / total
	values := make([]*ml.ConditionValue, total)
	for i := 0; i < total; i++ {
		v := ml.ConditionValue{}
		v.Type = ValueTypePercent
		v.Value = fmt.Sprintf("%d", (i*100)/total)
		values[i] = &v
	}

	r.rule.IconSet.Values = values
}

func (x iconSetRule) Type(t primitives.IconSetType) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.rule.IconSet.Type = t
		x.initValues(r)
	}
}

func (x iconSetRule) ReverseIcons(r *Info) {
	x.initIfRequired(r)
	r.rule.IconSet.Reverse = true
}

func (x iconSetRule) IconsOnly(r *Info) {
	x.initIfRequired(r)
	r.rule.IconSet.ShowValue = ml.OptionalBool(false)
}

func (x iconSetRule) setValue(r *Info, idx int, value string, settings []interface{}) {
	x.initIfRequired(r)

	total := len(r.rule.IconSet.Values) - 1

	if idx < 0 || idx >= total {
		return
	}

	idx = total - idx
	for _, p := range settings {
		switch pv := p.(type) {
		case string:
			if pv == ">" {
				r.rule.IconSet.Values[idx].GreaterThanEqual = ml.OptionalBool(false)
			}
		case primitives.ConditionValueType:
			r.rule.IconSet.Values[idx].Type = pv
		}
	}

	r.rule.IconSet.Values[idx].Value = value
}

func (x iconSetRule) Value(index int, value string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, index, value, settings)
	}
}

func (x iconSetRule) Validate(r *Info) error {
	for i, v := range r.rule.IconSet.Values {
		if !v.Type.IsAllowed(
			ValueTypeNumber,
			ValueTypePercent,
			ValueTypeFormula,
			ValueTypePercentile,
		) {
			return errors.New(fmt.Sprintf("iconSet: Not allowed type '%s' for value at index %d", v.Type, i))
		}
	}

	return nil
}
