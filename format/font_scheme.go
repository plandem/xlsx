package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for FontSchemeType
const (
	FontSchemeNone  primitives.FontSchemeType = "none"
	FontSchemeMinor primitives.FontSchemeType = "minor"
	FontSchemeMajor primitives.FontSchemeType = "major"
)
