package branch

import (
	"net/http"

	"encoding/json"

	"bjdaos/pegasus/pkg/wc/common"
	"bjdaos/pegasus/pkg/wc/db"
	httputil "github.com/1851616111/util/http"
	"github.com/julienschmidt/httprouter"
)

//curl -u michael:123456 -XPOST 192.168.199.198:9000/api/manage/branch  -d '{ "name": "北京第一体检中心", "desc": "北京第一体检中心成立于1980年, 位于北京市海淀区。", "address": { "province": "北京", "city": "北京", "district": "海淀", "details": "金源小区91#1-1-2"}, "tel": "010-55555555", "avatar":"123"}'
func CreateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	br := Branch{}
	if err := json.NewDecoder(r.Body).Decode(&br); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	if err := br.Validate(); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	if err := br.Create(db.Branch()); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.Response(w, 200, "ok")
	return
}

// curl 192.168.199.168:9000/api/manager/branches
func ListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//l, err := ListBranches(db.Branch())
	//if err != nil {
	//	httputil.Response(w, 400, err)
	//	return
	//}
	//
	//if err := json.NewEncoder(w).Encode(l); err != nil {
	//	httputil.Response(w, 400, err)
	//	return
	//}
	//return
	rspbyte, statuscode, err := common.Go_Through_Http("GET", "/api/organizations/wc", "")
	if statuscode != 200 {
		httputil.ResponseJson(w, 400, err)
	}
	if _, err := w.Write(rspbyte); err != nil {
		httputil.ResponseJson(w, 400, err)
		return
	}

	return

	//http.Redirect(w, r, common.AppointServe+"/api/organizations/wc", 301)
}

//curl -u michael:123456 -XPUT 192.168.199.198:9000/api/manage/branch/58e757f08fe64213cadb1f73  -d '{ "name": "北京第一体检中心2", "desc": "北京第一体检中心成立于1980年, 位于北京市海淀区。", "address": { "province": "北京", "city": "北京", "district": "海淀", "details": "金源小区91#1-1-2" }, "tel": "010-588888"}'
func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	br := Branch{}
	if err := json.NewDecoder(r.Body).Decode(&br); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	if err := br.Update(db.Branch(), id); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.Response(w, 200, "ok")
}
