// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conditional

import (
	"fmt"
	"github.com/plandem/xlsx/format/conditional/rule"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"

	// to link unexported
	_ "unsafe"
)

//
////go:linkname fromRule github.com/plandem/xlsx/format/conditional/rule.fromRule
//func fromRule(info *rule.Info) (*ml.ConditionalRule, *styles.Info)

//Info is objects that holds combined information about cell conditional format
type Info struct {
	info  *ml.ConditionalFormatting
	rules []*rule.Info
}

//Option is helper type to set options for conditional formatting
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

//Validate the conditional formatting information
func (f *Info) Validate() error {
	if len(f.info.Bounds) == 0 {
		return fmt.Errorf("no any refs for conditional formatting")
	}

	if len(f.rules) == 0 {
		return fmt.Errorf("no any rules for conditional formatting")
	}

	for i, r := range f.rules {
		rInfo, _ := rule.FromRule(r)
		if rInfo.Type == 0 {
			return fmt.Errorf("conditional rule#%d: no type", i)
		}

		if rInfo.Priority < 1 {
			return fmt.Errorf("conditional rule#%d: priority(%d) can't be higher thatn 1", i, rInfo.Priority)
		}

		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

//Pivot sets pivot flag of conditional formatting
func Pivot(cf *Info) {
	cf.info.Pivot = true
}

//Refs sets references that will be used for this conditional formatting
func Refs(refs ...primitives.Ref) Option {
	return func(cf *Info) {
		for _, ref := range refs {
			cf.info.Bounds.Add(ref)
		}
	}
}

//AddRule adds another rule to conditional formatting
func AddRule(options ...rule.Option) Option {
	return func(cf *Info) {
		r := rule.New(options...)

		rInfo, _ := rule.FromRule(r)
		rInfo.Priority = len(cf.rules) + 1
		cf.rules = append(cf.rules, r)
	}
}

//private method used to unpack Info
func From(f *Info) (*ml.ConditionalFormatting, []*styles.Info) {
	if len(f.rules) == 0 {
		return nil, nil
	}

	allRules := make([]*ml.ConditionalRule, len(f.rules))
	allStyles := make([]*styles.Info, len(f.rules))

	for i, r := range f.rules {
		rInfo, sInfo := rule.FromRule(r)

		allRules[i] = rInfo
		allStyles[i] = sInfo
	}

	f.info.Rules = allRules
	return f.info, allStyles
}
