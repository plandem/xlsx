package format

//FontSchemeType is a type to encode XSD ST_FontScheme
type FontSchemeType byte

const (
	_ FontSchemeType = iota
	FontSchemeNone
	FontSchemeMinor
	FontSchemeMajor
)

func (v FontSchemeType) String() string {
	var s string

	switch v {
	case FontSchemeNone:
		s = "none"
	case FontSchemeMinor:
		s = "minor"
	case FontSchemeMajor:
		s = "major"
	}

	return s
}
