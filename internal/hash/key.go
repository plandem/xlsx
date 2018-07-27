package hash

import (
	"crypto/md5"
	"fmt"
	"io"
)

//Key is alias type for string to provide Hash method
type Key string

//Hash returns md5 hash of key
func (k Key) Hash() string {
	h := md5.New()
	io.WriteString(h, string(k))
	return fmt.Sprintf("%x", h.Sum(nil))
}
