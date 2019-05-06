package conditional

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/format/conditional/rule"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	_ "unsafe"
)

//go:linkname fromRuleInfo github.com/plandem/xlsx/format/conditional/rule.from
func fromRuleInfo(r *rule.Info) (*ml.ConditionalRule, *styles.Info)

//Info is objects that holds combined information about cell conditional format
type Info struct {
	info  *ml.ConditionalFormatting
	rules []*rule.Info
}

type Option func(o *Info)

//New creates and returns Info object with requested options
func New(options ...Option) *Info {
	f := &Info{
		info:  &ml.ConditionalFormatting{},
		rules: []*rule.Info{},
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
		ri, _ := fromRuleInfo(r)

		if ri.Type == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no type", i))
		}

		if ri.Priority < 1 {
			return errors.New(fmt.Sprintf("conditional rule#%d: priority(%d) can't be higher thatn 1", i, ri.Priority))
		}

		if ri.Type == rule.TypeCellIs && ri.Operator == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no operator", i))
		}

		if ri.Type == rule.TypeTop10 && ri.Rank == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: wrong rank", i))
		}

		if ri.Type == rule.TypeContainsText && len(ri.Text) == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no text", i))
		}

		if ri.Type == rule.TypeTimePeriod && ri.TimePeriod == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no time period", i))
		}

		if ri.ColorScale != nil {
			if len(ri.ColorScale.Values) != len(ri.ColorScale.Colors) {
				return errors.New(fmt.Sprintf("conditional rule#%d: color scale should have equal numbers of colors and values", i))
			}

			if len(ri.ColorScale.Values) < 2 {
				return errors.New(fmt.Sprintf("conditional rule#%d: color scale should have at least 2 values", i))
			}
		}

		if ri.IconSet != nil && (len(ri.IconSet.Values) < 2) {
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

func AddRule(options ...rule.Option) Option {
	return func(cf *Info) {
		cf.rules = append(cf.rules, rule.New(options...))
	}
}

//private method used to unpack Info
func from(f *Info) (*ml.ConditionalFormatting, []*styles.Info) {
	if len(f.rules) == 0 {
		return nil, nil
	}

	allRules := make([]*ml.ConditionalRule, len(f.rules))
	allStyles := make([]*styles.Info, len(f.rules))

	for i, r := range f.rules {
		ri, si := fromRuleInfo(r)

		allRules[i] = ri
		allStyles[i] = si
	}

	f.info.Rules = allRules
	return f.info, allStyles
}
