package conditional

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//Info is objects that holds combined information about cell conditional format
type Info struct {
	info  *ml.ConditionalFormatting
	rules []*ruleInfo
}

type Option func(o *Info)

//New creates and returns Info object with requested options
func New(options ...Option) *Info {
	f := &Info{
		info:  &ml.ConditionalFormatting{},
		rules: []*ruleInfo{},
	}

	f.Set(options...)
	return f
}

//Set sets new options for conditional
func (f *Info) Set(options ...Option) {
	for _, o := range options {
		o(f)
	}
}

func (f *Info) Validate() error {
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

		if r.rule.Type == TypeCellIs && r.rule.Operator == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no operator", i))
		}

		if r.rule.Type == TypeTop10 && r.rule.Rank == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: wrong rank", i))
		}

		if r.rule.Type == TypeContainsText && len(r.rule.Text) == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no text", i))
		}

		if r.rule.Type == TypeTimePeriod && r.rule.TimePeriod == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no time period", i))
		}

		if r.rule.ColorScale != nil {
			if len(r.rule.ColorScale.Values) != len(r.rule.ColorScale.Colors) {
				return errors.New(fmt.Sprintf("conditional rule#%d: color scale should have equal numbers of colors and values", i))
			}

			if len(r.rule.ColorScale.Values) < 2 {
				return errors.New(fmt.Sprintf("conditional rule#%d: color scale should have at least 2 values", i))
			}
		}

		if r.rule.IconSet != nil && (len(r.rule.IconSet.Values) < 2) {
			return errors.New(fmt.Sprintf("conditional rule#%d: icon set should have at least 2 values", i))
		}
	}

	return nil
}

func Pivot(cf *Info) {
	cf.info.Pivot = true
}

func Refs(refs ...primitives.Ref) Option {
	return func(cf *Info) {
		for _, ref := range refs {
			cf.info.Bounds.Add(ref)
		}
	}
}

func AddRule(options ...ruleOption) Option {
	return func(cf *Info) {
		cf.rules = append(cf.rules, newRule(options...))
	}
}

//private method used to unpack Info
func from(f *Info) (*ml.ConditionalFormatting, []*styles.Info) {
	if len(f.rules) == 0 {
		return nil, nil
	}

	rules := make([]*ml.ConditionalRule, len(f.rules))
	si := make([]*styles.Info, len(f.rules))

	for i, r := range f.rules {
		rules[i] = r.rule
		si[i] = r.style
	}

	f.info.Rules = rules
	return f.info, si
}
