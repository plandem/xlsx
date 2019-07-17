// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/number_format"
)

//NumberFormat is option to update Info with provided custom format of number, but respecting built-in number formats
func NumberFormat(format string) func(*Info) {
	return func(s *Info) {
		*s.styleInfo.NumberFormat = number.New(-1, format)
	}
}

//NumberFormatID is option to update Info with provided id of already existing or built-in number format
func NumberFormatID(id int) func(*Info) {
	return func(s *Info) {
		*s.styleInfo.NumberFormat = number.New(id, "")
	}
}
