// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//Info hold advanced settings of sheet.
// N.B.: You should NOT mutate any value directly.
type Info struct {
	Visibility primitives.VisibilityType
}

//Option is helper type to set options for sheet
type Option func(co *Info)

//New create and returns option set for sheet
func New(settings ...Option) *Info {
	i := &Info{}
	i.Set(settings...)
	return i
}

//Set sets new options for option set
func (i *Info) Set(settings ...Option) {
	for _, o := range settings {
		o(i)
	}
}

//Visibility sets flag indicating if the affected column are hidden on this worksheet.
func Visibility(visibility primitives.VisibilityType) Option {
	return func(i *Info) {
		i.Visibility = visibility
	}
}
