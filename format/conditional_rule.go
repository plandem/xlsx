package format

import (
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
)

//conditionalRule is objects that holds combined information about conditional rule
type conditionalRule struct {
	rule  *ml.ConditionalRule
	style *StyleFormat
}

type conditionalRuleOption func(o *conditionalRule)

//Condition is a 'namespace' for all possible settings for conditional rule
var Condition conditionalRuleOption

//newConditionalRule creates and returns ConditionalRule object with requested options
func newConditionalRule(options ...conditionalRuleOption) *conditionalRule {
	r := &conditionalRule{
		rule: &ml.ConditionalRule{},
	}

	r.Set(options...)
	return r
}

//Set sets new options for conditional rule
func (r *conditionalRule) Set(options ...conditionalRuleOption) {
	for _, o := range options {
		o(r)
	}
}

func (co *conditionalRuleOption) AboveAverage(r *conditionalRule) {
	r.rule.AboveAverage = true
}

func (co *conditionalRuleOption) StopIfTrue(r *conditionalRule) {
	r.rule.StopIfTrue = true
}

func (co *conditionalRuleOption) Percent(r *conditionalRule) {
	r.rule.Percent = true
}

func (co *conditionalRuleOption) Bottom(r *conditionalRule) {
	r.rule.Bottom = true
}

func (co *conditionalRuleOption) EqualAverage(r *conditionalRule) {
	r.rule.EqualAverage = true
}

func (co *conditionalRuleOption) Priority(priority int) conditionalRuleOption {
	return func(r *conditionalRule) {
		r.rule.Priority = priority
	}
}

func (co *conditionalRuleOption) Style(style *StyleFormat) conditionalRuleOption {
	return func(r *conditionalRule) {
		r.style = style
	}
}

func (co *conditionalRuleOption) Type(t ConditionType) conditionalRuleOption {
	return func(r *conditionalRule) {
		r.rule.Type = t
	}
}

func (co *conditionalRuleOption) Operator(operator ConditionOperatorType) conditionalRuleOption {
	return func(r *conditionalRule) {
		r.rule.Operator = operator
	}
}

func (co *conditionalRuleOption) Text(text string) conditionalRuleOption {
	return func(r *conditionalRule) {
		r.rule.Text = text
	}
}

func (co *conditionalRuleOption) TimePeriod(period TimePeriodType) conditionalRuleOption {
	return func(r *conditionalRule) {
		r.rule.TimePeriod = period
	}
}

func (co *conditionalRuleOption) Rank(rank uint) conditionalRuleOption {
	return func(r *conditionalRule) {
		r.rule.Rank = rank
	}
}

func (co *conditionalRuleOption) Formula(formula Formula) conditionalRuleOption {
	return func(r *conditionalRule) {
		r.rule.Formula = formula
	}
}

func (co *conditionalRuleOption) ColorScale(pairs ...interface{}) conditionalRuleOption {
	return func(r *conditionalRule) {
		colorScale := &ml.ColorScale{}

		for _, p := range pairs {
			switch v := p.(type) {
			case string:
				colorScale.Colors = append(colorScale.Colors, color.New(v))
			case *conditionValue:
				colorScale.Values = append(colorScale.Values, &v.value)
			}
		}

		r.rule.ColorScale = colorScale
	}
}

func (co *conditionalRuleOption) IconSet(t IconSetType, percent bool, reverse bool, showValue bool, values ...*conditionValue) conditionalRuleOption {
	return func(r *conditionalRule) {
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

func (co *conditionalRuleOption) DataBar(min *conditionValue, minLength uint, max *conditionValue, maxLength uint, rgb string, showValue bool) conditionalRuleOption {
	return func(r *conditionalRule) {
		if min == nil {
			min = &conditionValue{}
		}

		if max == nil {
			max = &conditionValue{}
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
