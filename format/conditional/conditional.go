// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

//go:linkname fromRule github.com/plandem/xlsx/format/conditional/rule.fromRule
func fromRule(info *rule.Info) (*ml.ConditionalRule, *styles.Info)

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
		rInfo, _ := fromRule(r)
		if rInfo.Type == 0 {
			return errors.New(fmt.Sprintf("conditional rule#%d: no type", i))
		}

		if rInfo.Priority < 1 {
			return errors.New(fmt.Sprintf("conditional rule#%d: priority(%d) can't be higher thatn 1", i, rInfo.Priority))
		}

		if err := r.Validate(); err != nil {
			return err
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
		r := rule.New(options...)

		rInfo, _ := fromRule(r)
		rInfo.Priority = len(cf.rules) + 1
		cf.rules = append(cf.rules, r)
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
		rInfo, sInfo := fromRule(r)

		allRules[i] = rInfo
		allStyles[i] = sInfo
	}

	f.info.Rules = allRules
	return f.info, allStyles
}
