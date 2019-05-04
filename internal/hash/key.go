package hash

import (
	"hash/fnv"
)

//Key is alias type for string to provide Hash method
type Key string

//Hash returns FNV1 hash of key
func (k Key) Hash() uint64 {
	h := fnv.New64()
	h.Write([]byte(string(k)))
	return h.Sum64()
}
