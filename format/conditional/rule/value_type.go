// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for ConditionValueType
const (
	_ primitives.ConditionValueType = iota
	ValueTypeNumber
	ValueTypePercent
	ValueTypeHighest
	ValueTypeLowest
	ValueTypeFormula
	ValueTypePercentile
)

func init() {
	primitives.FromConditionValueType = map[primitives.ConditionValueType]string{
		ValueTypeNumber:     "num",
		ValueTypePercent:    "percent",
		ValueTypeHighest:    "max",
		ValueTypeLowest:     "min",
		ValueTypeFormula:    "formula",
		ValueTypePercentile: "percentile",
	}

	primitives.ToConditionValueType = make(map[string]primitives.ConditionValueType, len(primitives.FromConditionValueType))
	for k, v := range primitives.FromConditionValueType {
		primitives.ToConditionValueType[v] = k
	}
}
