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

	var result []string

	if reserved.InnerXML != nil {
		result = append(result, reserved.InnerXML.XML)
	} else {
		result = append(result, "")
	}

	for _, attr := range reserved.Attrs {
		result = append(result,
			attr.Name.Space,
			attr.Name.Local,
			attr.Value,
		)
	}

	return Key(strings.Join(result, ":"))
}
