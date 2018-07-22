package format

import (
	"github.com/plandem/xlsx/internal/ml"
)

//List of all possible values for FontSchemeType
const (
	FontSchemeNone  ml.FontSchemeType = "none"
	FontSchemeMinor ml.FontSchemeType = "minor"
	FontSchemeMajor ml.FontSchemeType = "major"
)