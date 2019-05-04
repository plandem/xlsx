package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strings"
)

//StringItem return string with values of ml.StringItem
func StringItem(si *ml.StringItem) Key {
	if si == nil {
		si = &ml.StringItem{}
	}

	result := []string{
		string(si.Text),
		string(Reserved(si.PhoneticPr)),
	}

	if si.RPh != nil {
		for _, r := range *si.RPh {
			result = append(result, string(Reserved(r)))
		}
	}

	if si.RichText != nil {
		for _, part := range *si.RichText {
			result = append(result, string(part.Text))

			if part.Font != nil {
				font := ml.Font(*part.Font)
				result = append(result, string(Font(&font)))
			}
		}
	}

	return Key(strings.Join(result, ":"))
}
