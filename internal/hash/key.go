// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hash

import (
	"hash/fnv"
)

//Key is string type to provide Hash method
type Key string

//Code is alias type for result of Hash method to unify logic around
type Code = uint64

//Hash returns FNV1 hash of key
func (k Key) Hash() Code {
	h := fnv.New64a()
	h.Write([]byte(string(k)))
	return Code(h.Sum64())
}
