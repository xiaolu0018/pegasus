package common

import (
	"192.168.199.199/bjdaos/pegasus/pkg/appoint"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"time"
)

func Go_Through_Http(method, url string, userid string) ([]byte, int, error) {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response

	var err error
	if req, err = http.NewRequest(method, AppointServe+url, nil); err != nil {
		glog.Errorln("newrequest err", err)
		return nil, 400, err
	}
	req.SetBasicAuth(userid, appoint.PASSWORD)
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

func Go_Through_HttpWithBody(method, url string, userid string, i interface{}) ([]byte, int, error) {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response

	var err error
	var rbuf bytes.Buffer
	json.NewEncoder(&rbuf).Encode(i)
	if req, err = http.NewRequest(method, AppointServe+url, &rbuf); err != nil {
		glog.Errorln("newrequest err", err)
		return nil, 400, err
	}
	req.SetBasicAuth(userid, appoint.PASSWORD)
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
