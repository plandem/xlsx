package format

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//ConditionalFormat is objects that holds combined information about cell conditional format
type ConditionalFormat struct {
	info  *ml.ConditionalFormatting
	rules []*conditionalRule
}

type conditionalOption func(o *ConditionalFormat)

//Conditions is a 'namespace' for all possible settings for ConditionalFormat
var Conditions conditionalOption

//NewConditions creates and returns ConditionalFormat object with requested options
func NewConditions(options ...conditionalOption) *ConditionalFormat {
	f := &ConditionalFormat{
		info:  &ml.ConditionalFormatting{},
		rules: []*conditionalRule{},
	}

	f.Set(options...)
	return f
}

//Set sets new options for conditional
func (f *ConditionalFormat) Set(options ...conditionalOption) {
	for _, o := range options {
		o(f)
	}
}

func (f *ConditionalFormat) Validate() error {
	if len(f.info.Bounds) == 0 {
		return errors.New("no any refs for conditional formatting")
	}

	if len(f.rules) == 0 {
		return errors.New("no any rules for conditional formatting")
	}

	for i, r := range f.rules {
		if r.rule.Type == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no type", i))
		}

		if r.rule.Priority < 1 {
			return errors.New(fmt.Sprintf("conditional rule#%d: priority(%d) can't be higher thatn 1", i, r.rule.Priority))
		}

		if r.rule.Type == ConditionTypeCellIs && r.rule.Operator == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no operator", i))
		}

		if r.rule.Type == ConditionTypeTop10 && r.rule.Rank == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: wrong rank", i))
		}

		if r.rule.Type == ConditionTypeContainsText && len(r.rule.Text) == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no text", i))
		}

		if r.rule.Type == ConditionTypeTimePeriod && r.rule.TimePeriod == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no time period", i))
		}
	}

	return nil
}

func (co *conditionalOption) Pivot(cf *ConditionalFormat) {
	cf.info.Pivot = true
}

func (co *conditionalOption) Refs(refs ...primitives.Ref) conditionalOption {
	return func(cf *ConditionalFormat) {
		for _, ref := range refs {
			cf.info.Bounds.Add(ref)
		}
	}
}

func (co *conditionalOption) Rule(options ...conditionalRuleOption) conditionalOption {
	return func(cf *ConditionalFormat) {
		cf.rules = append(cf.rules, newConditionalRule(options...))
	}
}

//private method used to unpack ConditionalFormat
func fromConditionalFormat(f *ConditionalFormat) (*ml.ConditionalFormatting, []*StyleFormat) {
	if len(f.rules) == 0 {
		return nil, nil
	}

	rules := make([]*ml.ConditionalRule, len(f.rules))
	styles := make([]*StyleFormat, len(f.rules))

	for i, r := range f.rules {
		rules[i] = r.rule
		styles[i] = r.style
	}

	f.info.Rules = rules
	return f.info, styles
}
