package format

import (
	"strings"
)

//ARGB is a type to encode XSD ST_UnsignedIntHex, hexBinary[4]
type ARGB string

//ColorToARGB check if color in #RGB format and convert it into ARGB format
func ColorToARGB(color string) ARGB {
	if len(color) > 1 {
		if color[0] == '#' {
			color = "FF" + color[1:]
		}
	}

	return ARGB(strings.ToUpper(color))
}

//ToIndex convert to index if it's possible, or return -1 in other case
func (c ARGB) ToIndex() int {
	for i, color := range indexedColors {
		if color == c {
			return i
		}
	}

	return -1
}
