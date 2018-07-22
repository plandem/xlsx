package format

import (
	"github.com/plandem/xlsx/internal/ml"
)

//List of all possible values for FontFamilyType
const (
	_ ml.FontFamilyType = iota
	FontFamilyRoman
	FontFamilySwiss
	FontFamilyModern
	FontFamilyScript
	FontFamilyDecorative
)
