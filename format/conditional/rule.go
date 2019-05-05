package conditional

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"

)

//ruleInfo is objects that holds combined information about conditional rule
type ruleInfo struct {
	rule  *ml.ConditionalRule
	style *styles.Info
}

type ruleOption func(o *ruleInfo)

//Rule is a 'namespace' for all possible settings for conditional rule
var Rule ruleOption

//newRule creates and returns ruleInfo object with requested options
func newRule(options ...ruleOption) *ruleInfo {
	r := &ruleInfo{
		rule: &ml.ConditionalRule{},
	}

	r.Set(options...)
	return r
}

//Set sets new options for conditional rule
func (r *ruleInfo) Set(options ...ruleOption) {
	for _, o := range options {
		o(r)
	}
}

func (co *ruleOption) AboveAverage(r *ruleInfo) {
	r.rule.AboveAverage = true
}

func (co *ruleOption) StopIfTrue(r *ruleInfo) {
	r.rule.StopIfTrue = true
}

func (co *ruleOption) Percent(r *ruleInfo) {
	r.rule.Percent = true
}

func (co *ruleOption) Bottom(r *ruleInfo) {
	r.rule.Bottom = true
}

func (co *ruleOption) EqualAverage(r *ruleInfo) {
	r.rule.EqualAverage = true
}

func (co *ruleOption) Priority(priority int) ruleOption {
	return func(r *ruleInfo) {
		r.rule.Priority = priority
	}
}

func (co *ruleOption) Style(style *styles.Info) ruleOption {
	return func(r *ruleInfo) {
		r.style = style
	}
}

func (co *ruleOption) Type(t Type) ruleOption {
	return func(r *ruleInfo) {
		r.rule.Type = t
	}
}

func (co *ruleOption) Operator(operator OperatorType) ruleOption {
	return func(r *ruleInfo) {
		r.rule.Operator = operator
	}
}

func (co *ruleOption) Text(text string) ruleOption {
	return func(r *ruleInfo) {
		r.rule.Text = text
	}
}

func (co *ruleOption) TimePeriod(period TimePeriodType) ruleOption {
	return func(r *ruleInfo) {
		r.rule.TimePeriod = period
	}
}

func (co *ruleOption) Rank(rank uint) ruleOption {
	return func(r *ruleInfo) {
		r.rule.Rank = rank
	}
}

func (co *ruleOption) Formula(formula Formula) ruleOption {
	return func(r *ruleInfo) {
		r.rule.Formula = formula
	}
}

func (co *ruleOption) ColorScale(pairs ...interface{}) ruleOption {
	return func(r *ruleInfo) {
		colorScale := &ml.ColorScale{}

		for _, p := range pairs {
			switch v := p.(type) {
			case string:
				colorScale.Colors = append(colorScale.Colors, color.New(v))
			case *value:
				colorScale.Values = append(colorScale.Values, &v.value)
			}
		}

		r.rule.ColorScale = colorScale
	}
}

func (co *ruleOption) IconSet(t IconSetType, percent bool, reverse bool, showValue bool, values ...*value) ruleOption {
	return func(r *ruleInfo) {
		iconSet := &ml.IconSet{
			Type:      t,
			Percent:   percent,
			Reverse:   reverse,
			ShowValue: showValue,
		}

		for _, v := range values {
			iconSet.Values = append(iconSet.Values, &v.value)
		}

		r.rule.IconSet = iconSet
	}
}

func (co *ruleOption) DataBar(min *value, minLength uint, max *value, maxLength uint, rgb string, showValue bool) ruleOption {
	return func(r *ruleInfo) {
		if min == nil {
			min = &value{}
		}

		if max == nil {
			max = &value{}
		}

		dataBar := &ml.DataBar{
			Values: []*ml.ConditionValue{
				&min.value,
				&max.value,
			},
			MinLength: minLength,
			MaxLength: maxLength,
			Color:     color.New(rgb),
			ShowValue: showValue,
		}

		r.rule.DataBar = dataBar
	}
}
