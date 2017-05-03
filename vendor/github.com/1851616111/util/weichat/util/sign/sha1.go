package sign

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

// sort by value
func SignToken(nonce, timestamp, token string) string {
	ps := []string{nonce, timestamp, token}
	sort.Strings(ps)

	return Sha1(ps[0] + ps[1] + ps[2])
}

func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func SignJsTicket(ticket, noncestr, url string, timestamp int64) string {
	target := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, noncestr, timestamp, url)
	return Sha1(target)
}
