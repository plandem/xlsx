package format

import (
	"github.com/plandem/xlsx/internal/ml/styles"
)

//List of all possible values for FontFamilyType
const (
	_ styles.FontFamilyType = iota
	FontFamilyRoman
	FontFamilySwiss
	FontFamilyModern
	FontFamilyScript
	FontFamilyDecorative
)
