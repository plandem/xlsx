// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"strings"
)

type baseRule byte

func (x baseRule) StopIfTrue(r *Info) {
	r.rule.StopIfTrue = true
}

func (x baseRule) escape(value string) string {
	// TODO: escape text value of formula (research what kind of escaping Excel is expecting)
	return strings.Replace(value, `"`, `""`, -1)
}

func (x baseRule) Validate(r *Info) error {
	return nil
}
