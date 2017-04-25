package sign

import (
	"fmt"
	"testing"
)

//signature=04cd5a1e9bb8aab6533164f7640cbd1c8bbcbe62&echostr=1703099520105042710&timestamp=1487913728&nonce=285423680
func TestSign(t *testing.T) {
	timestamp, nonce, token := "1487913728", "285423680", "bjdaos"

	signature := Sign(token, timestamp, nonce)
	fmt.Println(string(signature))
	if string(signature) != "04cd5a1e9bb8aab6533164f7640cbd1c8bbcbe62" {
		t.Fatal(string(signature))
	}
}
