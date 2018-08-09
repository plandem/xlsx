package format

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for FontFamilyType
const (
	_ primitives.FontFamilyType = iota
	FontFamilyRoman
	FontFamilySwiss
	FontFamilyModern
	FontFamilyScript
	FontFamilyDecorative
)
