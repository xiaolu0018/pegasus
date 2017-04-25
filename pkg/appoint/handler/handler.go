package handler

import (
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"

	httputil "github.com/1851616111/util/http"

	org "192.168.199.199/bjdaos/pegasus/pkg/appoint/organization"
)

func CreateBasicHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := ps.ByName("code")
	if len(code) == 0 {
		glog.Errorln("orgnization.CreateBasicHandler organization code not found")
		httputil.Response(rw, 400, "organization code not found")
		return
	}

	cfg := org.Config_Basic{}
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		glog.Errorf("orgnization.CreateBasicHandler decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	cfg.Org_Code = code
	if err := cfg.Validate(); err != nil {
		glog.Errorf("orgnization.CreateBasicHandler validate req params err %v\n", err)
		httputil.Response(rw, 400, cfg.Validate())
		return
	}

	if err := cfg.Create(); err != nil {
		glog.Errorf("orgnization.CreateBasicHandler Create err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	httputil.Response(rw, 200, "ok")
}

//创建 特殊项目的可预约人数
func CreateSpecialHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := ps.ByName("code")
	if code == "" {
		glog.Errorln("orgnization.CreateSpecialHandler organization code not found")
		httputil.Response(rw, 400, "organization code not found")
		return
	}

	cfg := org.Config_Special{}
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		glog.Errorln("orgnization.CreateSpecialHandler Decode ", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	cfg.Org_Code = code
	if err := cfg.Validate(); err != nil {
		glog.Errorf("orgnization.CreateSpecialHandler Validate req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	if err := cfg.Create(); err != nil {
		glog.Errorln("Orgnization CreateSpecialHandler Create", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	httputil.Response(rw, 200, "ok")
}

func ListHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	index, size := r.FormValue("index"), r.FormValue("size")
	if len(size) == 0 {
		size = "10"
	}
	if len(index) == 0 {
		index = "0"
	}

	var page_index, page_size int
	var err error
	if page_index, err = strconv.Atoi(index); err != nil {
		glog.Errorln("Orgnization ListHandler page_index", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	if page_size, err = strconv.Atoi(size); err != nil {
		glog.Errorln("Orgnization ListHandler page_size", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	if org, err := org.ListOrganizations(page_index, page_size); err != nil {
		glog.Errorln("Orgnization ListHandler ListOrgConfig", err.Error())
		httputil.Response(rw, 400, err)
		return
	} else {
		httputil.ResponseJson(rw, 200, org)
	}
}
