package common

import (
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"time"
)

func Go_Through_Http(url string) ([]byte, int, error) {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response

	var err error
	if req, err = http.NewRequest("GET", AppointServe+url, nil); err != nil {
		glog.Errorln("newrequest err", err)
		return nil, 400, err
	}

	if rsp, err = client.Do(req); err != nil {
		glog.Errorln("newrequest err", err)
		return nil, 400, err
	}

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("err", err.Error())
		return nil, 400, err
	}
	defer rsp.Body.Close()
	return buf, rsp.StatusCode, nil
}
