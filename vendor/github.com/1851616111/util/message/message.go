package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	httputil "github.com/1851616111/util/http"
)

var bytesPool *sync.Pool

func init() {
	bytesPool = &sync.Pool{}
	bytesPool.New = func() interface{} {
		return &bytes.Buffer{}
	}
}

var codeMsgM map[int]interface{} = map[int]interface{}{
	1000: "success",
	1001: ERR_SERVER_INNER_ERROR,
	1004: ERR_REQ_NOT_FOUND_ERROR,
}

//向用户返回错误信息
var ERR_SERVER_INNER_ERROR error = errors.New("Internal Server Error")
var ERR_REQ_NOT_FOUND_ERROR error = errors.New("Request Not Found")

const _Inner_Error = `{"code":1001,"message":"Internal Server Error"}`
const _Req_Not_Find = `{"code":1004,"message":"Request Not Found"}`
const _Param_Not_Find = `{"code":1004,"message":"Param %s Not Found"}`

type msg struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

func Success(w http.ResponseWriter, data []byte) {
	httputil.Response(w, 200, generateData(data))
}

func SuccessI(w http.ResponseWriter, obj interface{}) error {
	return json.NewEncoder(w).Encode(obj)
}

func InnerError(w http.ResponseWriter) {
	httputil.Response(w, 400, message(_Inner_Error))
}

func NotFoundError(w http.ResponseWriter) {
	httputil.Response(w, 404, message(_Req_Not_Find))
}

func ParamNotFound(w http.ResponseWriter, param string) {
	httputil.Response(w, 400, message(fmt.Sprintf(_Param_Not_Find, param)))
}

//{"code":1000,"data":null,"message":"success"}
func generateData(data []byte) *bytes.Buffer {
	buf := message(`{"code":1000,"data":`)
	buf.Write(data)
	buf.WriteString(`,"message":"success"}`)

	return buf
}

func ReadBody(body io.Reader) ([]byte, error) {
	tmp := bytesPool.Get().(*bytes.Buffer)
	if _, err := tmp.ReadFrom(body); err != nil {
		return nil, err
	}

	data := tmp.Bytes()
	if tmp.Len() > 0 && data[tmp.Len()-1] == 0xa {
		data = data[:tmp.Len()-1]
	}

	return data, nil
}

func message(msgs ...string) *bytes.Buffer {
	buf := bytesPool.Get().(*bytes.Buffer)
	if len(msgs) > 0 {
		for _, m := range msgs {
			buf.WriteString(m)
		}
	}

	return buf
}
