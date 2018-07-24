package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
)

//List of all possible values for FontSchemeType
const (
	FontSchemeNone  styles.FontSchemeType = "none"
	FontSchemeMinor styles.FontSchemeType = "minor"
	FontSchemeMajor styles.FontSchemeType = "major"
)
