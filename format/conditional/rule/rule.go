// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
)

type validator interface {
	Validate(r *Info) error
}

//Info is objects that holds combined information about conditional rule
type Info struct {
	initialized bool
	rule        *ml.ConditionalRule
	style       *styles.Info
	validator   validator
}

//Option is helper type to set options for rule
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

//Validate validate rule if required
func (r *Info) Validate() error {
	if r.validator != nil {
		return r.validator.Validate(r)
	}

	return nil
}

//private method used to access private fields that we want to protect from direct mutating
func FromRule(r *Info) (*ml.ConditionalRule, *styles.Info) {
	return r.rule, r.style
}
