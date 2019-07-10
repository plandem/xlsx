// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for UnderlineType
const (
	UnderlineTypeSingle           primitives.UnderlineType = "single"
	UnderlineTypeDouble           primitives.UnderlineType = "double"
	UnderlineTypeSingleAccounting primitives.UnderlineType = "singleAccounting"
	UnderlineTypeDoubleAccounting primitives.UnderlineType = "doubleAccounting"
	UnderlineTypeNone             primitives.UnderlineType = "none"
)
