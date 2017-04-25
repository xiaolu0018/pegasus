package sign

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

// sort by value
func Sign(nonce, timestamp, token string) string {
	ps := []string{nonce, timestamp, token}
	sort.Strings(ps)

	return Sha1(ps[0] + ps[1] + ps[2])
}

func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
