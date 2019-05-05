package format

import "github.com/plandem/xlsx/internal/ml"

//ConditionalFormat is objects that holds combined information about cell conditional format
type ConditionalFormat struct {
	info *ml.ConditionalFormatting
}

type conditionalOption func(o *ConditionalFormat)

//NewConditions creates and returns ConditionalFormat object with requested options
func NewConditions(options ...conditionalOption) *ConditionalFormat {
	s := &ConditionalFormat{}

	s.Set(options...)
	return s
}

//Set sets new options for conditional
func (s *ConditionalFormat) Set(options ...conditionalOption) {
	for _, o := range options {
		o(s)
	}
}
