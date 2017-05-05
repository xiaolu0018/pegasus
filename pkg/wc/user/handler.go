package user

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"

	httputil "github.com/1851616111/util/http"

	"bjdaos/pegasus/pkg/wc/common"
)

//保存或更新用户的个人基本信息
// curl --data '{"mobile":"12345678910","idcard":"34052419800101001X","name":"张三"}' http://www.elepick.com/api/user
func UpsertInfoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u := User{}
	//u := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		glog.Errorf("UpsertInfoHandler err", err.Error())
		httputil.ResponseJson(w, 400, "params invalid"+err.Error())
		return
	}

	if err := u.CreateValidate(); err != nil {
		glog.Errorf("CreateValidate err", err.Error())
		httputil.ResponseJson(w, 400, err.Error())
		return
	}
	u.ID = ps.ByName(common.AuthHeaderKey)
	glog.Errorln("upsert u,id", ps.ByName(common.AuthHeaderKey))
	if err := u.Upsert(); err != nil {
		glog.Errorf("user.UpsertInfoHandler: user(%v) err %v\n", u, err)
		httputil.ResponseJson(w, 400, err)
		return
	}

	httputil.ResponseJson(w, 200, "ok")
}

//保存或更新用户的label
// curl --data '{"bingshi":{"gaoxueya","gaoxuezhi"},"jiazushi":{"guanxinbing","naogeng"}}' http://www.elepick.com/api/user/12345678910/health
func UpdateLabelHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userid := ps.ByName(common.AuthHeaderKey)
	h := Health{}
	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		glog.Errorln("user.UpdateLabelHandler err", err.Error())
		httputil.ResponseJson(w, 404, "params invalid")
		return
	}

	if err := h.Upsert(userid); err != nil {
		glog.Errorf("user.UpdateLabelHandler: updatelabel(%v) err %v\n", userid, err)
		httputil.ResponseJson(w, 400, err)
		return
	}
	httputil.ResponseJson(w, 200, "ok")
}

func GetLabelHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	userid := ps.ByName(common.AuthHeaderKey)
	if h, err := GetHealth(userid); err != nil {
		glog.Errorf("user.GetLabelHandler:err %v\n", err)
		httputil.ResponseJson(w, 400, err)
		return
	} else {
		json.NewEncoder(w).Encode(h)
	}
	return
}

//得到用户信息
//curl http://www.elepick.com/api/user/12345678910/health
func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userid := ps.ByName(common.AuthHeaderKey)

	if u, err := GetUserByid(userid); err != nil {
		glog.Errorf("user.GetHandler:err %v\n", err)
		httputil.ResponseJson(w, 400, err)
		return
	} else {
		httputil.ResponseJson(w, 200, u)
	}
}
