package vote

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"net/http"
	"database/sql"
	"encoding/json"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	httputil "github.com/1851616111/util/http"
	"github.com/1851616111/util/rand"
	"github.com/1851616111/util/weichat/handler"
	apiotoken "github.com/1851616111/util/weichat/util/api-token"
	"github.com/1851616111/util/weichat/util/sign"
)

var dbI DBInterface
var URL_REGISTER_HTML string
var APPID string

const DEFAULT_PAGE_SIZE = "8"

func AddRouter(r *httprouter.Router, dist string) {

	r.GET("/api/basic/signature", handler.DeveloperValidater)
	r.POST("/api/basic/signature", handler.EventAction)

	r.GET("/api/activity/voters", ListVotersHandler)
	r.GET("/api/activity/voter", GetVoterByOpenIDHandler)
	r.POST("/api/activity/voter", RegisterVoterHandler)
	r.POST("/api/activity/voter/:voterid/vote", VoteHandler)

	r.GET("/api/activity/jsconfig", ExchangeJSConfigHandler)

	var err error
	dist, err = filepath.Abs(dist)
	if err != nil {
		panic(err)
	}

	r.ServeFiles("/dist/activity/*filepath", http.Dir(dist))
}

func ExchangeJSConfigHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	openid := r.FormValue("openid")
	if len(openid) == 0 {
		httputil.Response(w, 400, "openid not found")
		return
	}

	ticket := apiotoken.TokenCtrl.GetTicket()
	nonce := strings.ToLower(rand.String(5))
	timeStamp := time.Now().Unix()
	//http://mp.weixin.qq.com?params=value
	signStr := sign.SignJsTicket(ticket, nonce, fmt.Sprintf("%s?openid=%s", URL_REGISTER_HTML, openid), timeStamp)
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
		CH_CACHE_IMAGES <- v.Image
	}

	if err := dbI.Register(v); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint ") {
			glog.Errorf("weichat regist voter err %v\n", err)
			httputil.Response(w, 409, "user duplicate")
			return
		}
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

	if err := dbI.Vote(openid, voteID); err != nil {
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


	l, err := dbI.ListVoters(searchKey, index, size)
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

	v, err := dbI.GetVoter(openid)
	if err != nil {
		if err == sql.ErrNoRows {
			httputil.Response(w, 404, "not found")
			return
		}
		httputil.Response(w, 400, err)
		return
	}

	httputil.ResponseJson(w, 200, v)
	return
}

func SetDB(i DBInterface) {
	dbI = i
}
