package hash

import (
	"crypto/md5"
	"io"
	"fmt"
)

type Key string

func (k Key) Hash() string {
	h := md5.New()
	io.WriteString(h, string(k))
	return fmt.Sprintf("%x", h.Sum(nil))
}

