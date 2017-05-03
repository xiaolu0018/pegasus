package sign

import (
	"testing"
)

//signature=04cd5a1e9bb8aab6533164f7640cbd1c8bbcbe62&echostr=1703099520105042710&timestamp=1487913728&nonce=285423680
func TestSignToken(t *testing.T) {
	timestamp, nonce, token := "1487913728", "285423680", "bjdaos"

	signature := SignToken(token, timestamp, nonce)
	if string(signature) != "04cd5a1e9bb8aab6533164f7640cbd1c8bbcbe62" {
		t.Fatal(string(signature))
	}
}

func TestSignJsTicket(t *testing.T) {
	targetSign := "9836edd65dcdb926f8de2d3c3aaf65439c3d76f3"

	if str := SignJsTicket("kgt8ON7yVITDhtdwci0qeYgh_SptvWn_34kRNSvKy3Tt985LgwE7e-XlNgaS2eGfsPnUoH8L2l0MC9PjIPcAgQ",
		"qwe", "1493718359", "www.elepick.com"); str != targetSign {
		t.Fatal(str)
	}
}
