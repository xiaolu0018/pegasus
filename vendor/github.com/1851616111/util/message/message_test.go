package message

import (
	httputil "github.com/1851616111/util/http"
	"testing"
)

func Test_SuccessMessage(t *testing.T) {
	rsp, _ := httputil.Send(&httputil.HttpSpec{
		URL: "http://127.0.0.1:1080/controller/reporter/api",
		BasicAuth: &httputil.BasicAuth{
			User:     "michael",
			Password: "6485C640-E349-40D8-BC71-F5D4FBE7B6C1",
		},
	})

	generateData(rsp.Body)
}
