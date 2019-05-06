package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/color"
	"github.com/plandem/xlsx/internal/ml"
)

//Info is objects that holds combined information about conditional rule
type Info struct {
	rule  *ml.ConditionalRule
	style *styles.Info
}

type Option func(o *Info)

//New creates and returns Info object with requested options
func New(options ...Option) *Info {
	r := &Info{
		rule: &ml.ConditionalRule{},
	}

	r.Set(options...)
	return r
}

//Set sets new options for conditional rule
func (r *Info) Set(options ...Option) {
	for _, o := range options {
		o(r)
	}
}

func AboveAverage(r *Info) {
	r.rule.AboveAverage = true
}

func StopIfTrue(r *Info) {
	r.rule.StopIfTrue = true
}

func Percent(r *Info) {
	r.rule.Percent = true
}

func Bottom(r *Info) {
	r.rule.Bottom = true
}

func EqualAverage(r *Info) {
	r.rule.EqualAverage = true
}

func Priority(priority int) Option {
	return func(r *Info) {
		r.rule.Priority = priority
	}
}

func Style(style *styles.Info) Option {
	return func(r *Info) {
		r.style = style
	}
}

func Type(t ConditionType) Option {
	return func(r *Info) {
		r.rule.Type = t
	}
}

func Operator(operator OperatorType) Option {
	return func(r *Info) {
		r.rule.Operator = operator
	}
}

func Text(text string) Option {
	return func(r *Info) {
		r.rule.Text = text
	}
}

func TimePeriod(period TimePeriodType) Option {
	return func(r *Info) {
		r.rule.TimePeriod = period
	}
}

func Rank(rank uint) Option {
	return func(r *Info) {
		r.rule.Rank = rank
	}
}

func Formula(formula FormulaType) Option {
	return func(r *Info) {
		r.rule.Formula = formula
	}
}

func ColorScale(pairs ...interface{}) Option {
	return func(r *Info) {
		colorScale := &ml.ColorScale{}

		for _, p := range pairs {
			switch v := p.(type) {
			case string:
				colorScale.Colors = append(colorScale.Colors, color.New(v))
			case *value:
				colorScale.Values = append(colorScale.Values, &v.value)
			}
		}

		r.rule.Type = TypeColorScale
		r.rule.ColorScale = colorScale
	}
}

func IconSet(t IconSetType, percent bool, reverse bool, showValue bool, values ...*value) Option {
	return func(r *Info) {
		iconSet := &ml.IconSet{
			Type:      t,
			Percent:   percent,
			Reverse:   reverse,
			ShowValue: showValue,
		}

		for _, v := range values {
			iconSet.Values = append(iconSet.Values, &v.value)
		}

		r.rule.Type = TypeIconSet
		r.rule.IconSet = iconSet
	}
}

func DataBar(min *value, minLength uint, max *value, maxLength uint, rgb string, showValue bool) Option {
	return func(r *Info) {
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

		r.rule.Type = TypeDataBar
		r.rule.DataBar = dataBar
	}
}

func from(r *Info) (*ml.ConditionalRule, *styles.Info) {
	return r.rule, r.style
}
