package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"time"
)

const PASSWORD = "58f06cdfa46d12688c23405b"

func Go_Through_Http(method, appointServe, url string, userid string) ([]byte, int, error) {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response

	var err error
	if req, err = http.NewRequest(method, appointServe+url, nil); err != nil {
		glog.Errorln("newrequest err", err)
		return nil, 400, err
	}
	req.SetBasicAuth(userid, PASSWORD)
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

func Go_Through_HttpWithBody(method, appointServe, url string, userid string, i interface{}) ([]byte, int, error) {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response

	var err error
	var rbuf bytes.Buffer
	json.NewEncoder(&rbuf).Encode(i)
	if req, err = http.NewRequest(method, appointServe+url, &rbuf); err != nil {
		glog.Errorln("common.Go_Through_HttpWithBody.newrequest err", err)
		return nil, 400, err
	}
	//req.SetBasicAuth(userid, appoint.PASSWORD)
	if rsp, err = client.Do(req); err != nil {
		glog.Errorln("common.Go_Through_HttpWithBody.client.Do err", err)
		return nil, 400, err
	}

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		glog.Errorln("common.Go_Through_HttpWithBody.ReadAll err", err)
		return nil, 400, err
	}
	defer rsp.Body.Close()
	return buf, rsp.StatusCode, nil
}
