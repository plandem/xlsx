package rule

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"strconv"
	"strings"
)

type textRule struct {
	baseRule
}

var Text textRule

func (x textRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Text
		r.rule = &ml.ConditionalRule{}
	}
}

func (x textRule) setValue(r *Info, value string, formula string, t primitives.ConditionType, operator primitives.ConditionOperatorType, settings []interface{}) {
	x.initIfRequired(r)
	r.rule.Type = t
	r.rule.Operator = operator
	r.rule.Text = value

	formula = strings.ReplaceAll(formula, ":length:", strconv.FormatInt(int64(len(value)), 10))
	formula = fmt.Sprintf(formula, x.escape(value))
	r.rule.Formula = []ml.Formula{ml.Formula(formula)}

	for _, p := range settings {
		switch pv := p.(type) {
		case *styles.Info:
			r.style = pv
		}
	}
}

func (x textRule) Contains(s string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, s, `NOT(ISERROR(SEARCH("%s",:cell:)))`, primitives.ConditionTypeContainsText, primitives.ConditionOperatorContainsText, settings)
	}
}

func (x textRule) NotContains(s string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, s, `ISERROR(SEARCH("%s",:cell:))`, primitives.ConditionTypeNotContainsText, primitives.ConditionOperatorNotContains, settings)
	}
}

func (x textRule) BeginsWith(s string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, s, `LEFT(:cell:,:length:)="%s"`, primitives.ConditionTypeBeginsWith, primitives.ConditionOperatorBeginsWith, settings)
	}
}

func (x textRule) EndsWith(s string, settings ...interface{}) Option {
	return func(r *Info) {
		x.setValue(r, s, `RIGHT(:cell:,:length:)="%s"`, primitives.ConditionTypeEndsWith, primitives.ConditionOperatorEndsWith, settings)
	}
}

func (x textRule) Validate(r *Info) error {
	if len(r.rule.Text) == 0 {
		return errors.New("text: no text to look for")
	}

	return nil
}
