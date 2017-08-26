package format

//FontVAlignType is a type to encode XSD ST_VerticalAlignRun
type FontVAlignType byte

const (
	_ FontVAlignType = iota
	FontVAlignBaseline
	FontVAlignSuperscript
	FontVAlignSubscript
)

func (v FontVAlignType) String() string {
	var s string

	switch v {
	case FontVAlignBaseline:
		s = "baseline"
	case FontVAlignSuperscript:
		s = "superscript"
	case FontVAlignSubscript:
		s = "subscript"
	}

	return s
}
