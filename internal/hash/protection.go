// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//Protection return string with all values of protection
func Protection(protection *ml.CellProtection) Key {
	if protection == nil {
		protection = &ml.CellProtection{}
	}

	return Key(strings.Join([]string{
		strconv.FormatBool(protection.Locked),
		strconv.FormatBool(protection.Hidden),
	}, ":"))
}
