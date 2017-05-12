package vote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	httputil "github.com/1851616111/util/http"
	"github.com/1851616111/util/image"
	"github.com/1851616111/util/rand"
	"github.com/1851616111/util/weichat/handler"
	apiotoken "github.com/1851616111/util/weichat/util/api-token"
	"github.com/1851616111/util/weichat/util/sign"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"encoding/base64"
	"bytes"
)

var DBI DBInterface
var APPID string
var iosImagePath string

const DEFAULT_PAGE_SIZE = "100"

func AddRouter(r *httprouter.Router, dist string) {
	r.GET("/api/basic/signature", handler.DeveloperValidater)
	r.POST("/api/basic/signature", handler.EventAction)

	r.GET("/api/activity/voters", ListVotersHandler)
	r.GET("/api/activity/voter", GetVoterByOpenIDHandler)
	r.POST("/api/activity/voter", RegisterVoterHandler)
	r.POST("/api/activity/voter/:voterid/vote", VoteHandler)
	r.POST("/api/ios/image", RegisterImageHandler)
	r.GET("/api/ios/image", GetImageHandler)

	r.GET("/api/activity/jsconfig", ExchangeJSConfigHandler)
	r.POST("/api/base64/person", EncodingBase64Handler)
	r.GET("/api/base64/person", DecodingBase64Handler)

	var err error
	dist, err = filepath.Abs(dist)
	if err != nil {
		panic(err)
	}

	iosImagePath = fmt.Sprintf("%s/voterimages", dist)

	r.ServeFiles("/dist/activity/*filepath", http.Dir(dist))
}

func RegisterImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	imageName := fmt.Sprintf("%d_%s.jpg", time.Now().UnixNano(), rand.String(10))
	imagePath := iosImagePath + "/" + imageName
	data := r.FormValue("data")

	idx := strings.Index(data, ",")
	data = data[idx+1:]

	if err := image.GenImageFromBase64([]byte(data), imagePath); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.Response(w, 200, imageName)
	return
}

func GetImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	image := r.FormValue("image")
	if len(image) == 0 {
		httputil.Response(w, 400, "param image not found")
		return
	}

	target := fmt.Sprintf("%s/%s.jpg", iosImagePath, image)
	data, err := ioutil.ReadFile(target)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	w.WriteHeader(200)
	w.Write(data)
	return
}


func ExchangeJSConfigHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reqPage := r.FormValue("reqpage")
	if len(reqPage) == 0 {
		httputil.Response(w, 400, "reqPage not found")
		return
	}

	ticket := apiotoken.TokenCtrl.GetTicket()
	nonce := strings.ToLower(rand.String(5))
	timeStamp := time.Now().Unix()
	//http://mp.weixin.qq.com?params=value
	signStr := sign.SignJsTicket(ticket, nonce, reqPage, timeStamp)
	m := map[string]interface{}{
		"appid":     APPID,
		"timestamp": timeStamp,
		"noncestr":  nonce,
		"signature": signStr,
	}

	httputil.ResponseJson(w, 200, m)
	return
}

func RegisterVoterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	v := &Voter{}
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		glog.Errorf("register voter, decode json object err %v\n", err)
		httputil.Response(w, 400, err)
		return
	}
	if err := v.ValidateRegister(); err != nil {
		glog.Errorf("register voter, validate json object err %v\n", err)
		httputil.Response(w, 400, err)
		return
	}
	v.Complete()

	if CH_CACHE_IMAGES != nil {
		if !strings.HasSuffix(v.Image, ".jpg") {
			CH_CACHE_IMAGES <- v.Image //android
		} else {
			v.imageCached = true //ios
		}
	}

	if err := DBI.Register(v); err != nil {
		glog.Errorf("register voter, database operate err %v\n", err)
		httputil.Response(w, 400, err)
		return
	}

	httputil.Response(w, 200, "ok")
}

func VoteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	openid := r.FormValue("openid")
	voteID := ps.ByName("voterid")

	if len(voteID) == 0 || len(voteID) > 50 {
		httputil.Response(w, 400, "voterid invalid")
		return
	}

	if err := DBI.Vote(openid, voteID); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.Response(w, 200, "ok")
}

func ListVotersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	indexStr, sizeStr, keyStr := r.FormValue("index"), r.FormValue("size"), r.FormValue("key")
	if len(indexStr) == 0 {
		indexStr = "1"
	}

	if len(sizeStr) == 0 {
		sizeStr = DEFAULT_PAGE_SIZE
	}

	if len(indexStr) > 5 || len(sizeStr) > 5 || len(keyStr) > 50 {
		httputil.Response(w, 400, "invalid params")
		return
	}

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	var searchKey interface{}
	if len(keyStr) > 0 {
		searchVoterID, err := strconv.Atoi(keyStr)
		if err == nil {
			searchKey = searchVoterID
		} else {
			searchKey = keyStr
		}
	}

	l, err := DBI.ListVoters(searchKey, index, size)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.ResponseJson(w, 200, l)
	return
}

func GetVoterByOpenIDHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	openid := r.FormValue("openid")
	if len(openid) == 0 {
		httputil.Response(w, 400, "param openid not found")
		return
	}

	v, err := DBI.GetVoterStatus(openid)
	if err != nil {
		httputil.ResponseJson(w, 400, err)
		return
	}

	httputil.ResponseJson(w, 200, v)
	return
}

func SetDB(i DBInterface) {
	DBI = i
}

var baseEncoder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

//curl http://127.0.0.1:8000/api/base64/person -d '{"name":"中国人", "age":10, "sex":"man"}'
func EncodingBase64Handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}
	defer r.Body.Close()

	m := map[string]interface{}{}
	if err := json.Unmarshal(data, &m); err != nil {
		httputil.Response(w, 400, "person data invalid")
		return
	}

	dst := make([]byte, 2* len(data))
	baseEncoder.Encode(dst, data)

	ret := fmt.Sprintf("person=%s", string(dst))

	httputil.Response(w, 200, ret)
	return
}

//curl http://127.0.0.1:8000/api/base64/person?person=eyJuYW1lIjoi5Lit5Zu95Lq6IiwgImFnZSI6MTAsICJzZXgiOiJtYW4ifQ==
func DecodingBase64Handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := r.FormValue("person")
	if len(person) == 0 {
		httputil.Response(w, 400, "param person not found")
		return
	}

	decodeBytes, err := baseEncoder.DecodeString(person)
	if err != nil {
		httputil.Response(w, 400, "param person not found")
		return
	}

	dst := bytes.NewBuffer(decodeBytes)

	m := map[string]interface{}{}
	if err := json.NewDecoder(dst).Decode(&m); err != nil {
		httputil.Response(w, 400, "parse data err")
		return
	}

	httputil.ResponseJson(w, 200, m)
	return
}