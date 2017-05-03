package vote

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fmt"
	httputil "github.com/1851616111/util/http"
	"github.com/1851616111/util/rand"
	"github.com/1851616111/util/weichat/handler"
	apiotoken "github.com/1851616111/util/weichat/util/api-token"
	"github.com/1851616111/util/weichat/util/sign"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"path/filepath"
	"strings"
	"time"
)

var dbI DBInterface
var URL_REGISTER_HTML string
var APPID string

const DEFAULT_PAGE_SIZE = "20"

func AddRouter(r *httprouter.Router, dist string) {

	r.GET("/api/basic/signature", handler.DeveloperValidater)
	r.POST("/api/basic/signature", handler.EventAction)

	r.GET("/api/activity/voters", ListVotersHandler)
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

	if err := dbI.Register(v); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint ") {
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
	indexS, sizeS := r.FormValue("index"), r.FormValue("size")
	if len(indexS) == 0 {
		indexS = "1"
	}

	if len(sizeS) == 0 {
		sizeS = DEFAULT_PAGE_SIZE
	}

	index, err := strconv.Atoi(indexS)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	size, err := strconv.Atoi(sizeS)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	l, err := dbI.ListVoters(index, size)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.ResponseJson(w, 200, l)
	return
}

func SetDB(i DBInterface) {
	dbI = i
}
