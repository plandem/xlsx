// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hash

import (
	"github.com/plandem/ooxml/ml"
	"strings"
)

//Reserved return string with values of reserved
func Reserved(reserved *ml.Reserved) Key {
	if reserved == nil {
		reserved = &ml.Reserved{}
	}

	result := make([]string, 0, len(reserved.Attrs))
	result = append(result, reserved.InnerXML)

	for _, attr := range reserved.Attrs {
		result = append(result,
			attr.Name.Space,
			attr.Name.Local,
			attr.Value,
		)
	}

	return Key(strings.Join(result, ":"))
}
