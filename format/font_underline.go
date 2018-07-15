package format

//UnderlineType is a type to encode XSD CT_UnderlineProperty
type UnderlineType byte

//List of all possible values for UnderlineType
const (
	_ UnderlineType = iota
	UnderlineTypeSingle
	UnderlineTypeDouble
	UnderlineTypeSingleAccounting
	UnderlineTypeDoubleAccounting
	UnderlineTypeNone
)

func (v UnderlineType) String() string {
	var s string

	switch v {
	case UnderlineTypeSingle:
		s = "single"
	case UnderlineTypeDouble:
		s = "double"
	case UnderlineTypeSingleAccounting:
		s = "singleAccounting"
	case UnderlineTypeDoubleAccounting:
		s = "doubleAccounting"
	case UnderlineTypeNone:
		s = "none"
	}

	return s
}
