package pinto

import "github.com/1851616111/util/http"

const JSON_KEY_CHECKUP_CODES = "chechops"

type SDK interface {
	GetCheckupCodesBySaleCodes(codes []string) ([]string, error)
}

type sdkImpl struct {
	addr string
	auth http.BasicAuth
	//cheiupPath "sdfsdfasd/asdfzsdf/"

}
