package primitives

import (
	"strings"
	"unsafe"
)

//RefList is a type to encode XSD ST_Sqref, a reference that identifies list of cells or ranges of cells. E.g.: "N28 B5:N10 C10"
type RefList string

//ToRefs returns refs of RefList
func (r RefList) ToRefs() []Ref {
	//N.B.: even it's unsafe way, it still should be safe, because Ref is string
	var refs = strings.Split(string(r), " ")
	return *(*[]Ref)(unsafe.Pointer(&refs))
}

//ToBoundsList returns related bounds of RefList
func (r RefList) ToBoundsList() BoundsList {
	return BoundsListFromRefs(r.ToRefs()...)
}

//RefListFromRefs returns RefList for refs
func RefListFromRefs(refs ...Ref) RefList {
	//N.B.: even it's unsafe way, it still should be safe, because Ref is string
	list := *(*[]string)(unsafe.Pointer(&refs))
	return RefList(strings.Join(list, " "))
}
