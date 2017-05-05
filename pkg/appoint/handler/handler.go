package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"

	httputil "github.com/1851616111/util/http"

	"bjdaos/pegasus/pkg/appoint/appointment"
	"bjdaos/pegasus/pkg/appoint/db"
	org "bjdaos/pegasus/pkg/appoint/organization"
	"bjdaos/pegasus/pkg/common/api/pinto"
)

func CreateBasicHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := ps.ByName("code")
	if len(code) == 0 {
		glog.Errorln("orgnization.CreateBasicHandler organization code not found")
		httputil.Response(rw, 400, "organization code not found")
		return
	}

	//cfg := org.Config_Basic{}
	cfg := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		glog.Errorf("orgnization.CreateBasicHandler decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	glog.Errorln("cfg ___,", cfg)
	//cfg.Org_Code = code
	//if err := cfg.Validate(); err != nil {
	//	glog.Errorf("orgnization.CreateBasicHandler validate req params err %v\n", err)
	//	httputil.Response(rw, 400, cfg.Validate())
	//	return
	//}
	//
	//if err := cfg.Create(); err != nil {
	//	glog.Errorf("orgnization.CreateBasicHandler Create err %v\n", err.Error())
	//	httputil.Response(rw, 400, err)
	//	return
	//}

	httputil.Response(rw, 200, "ok")
}

func GetBasicHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := ps.ByName("code")
	if len(code) == 0 {
		glog.Errorln("orgnization.GetBasicHandler organization code not found")
		httputil.Response(rw, 400, "organization code not found")
		return
	}
	org_basic := &org.Config_Basic{}
	var err error
	if org_basic, err = org.GetConfigBasic(code); err != nil {
		glog.Errorln("orgnization.GetBasicHandler GetConfigBasic err ", err.Error())
		httputil.Response(rw, 400, err.Error())
		return
	}
	if err := json.NewEncoder(rw).Encode(org_basic); err != nil {
		httputil.Response(rw, 400, err)
		return
	}
	return
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

func ListHandlerWC(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var org_wcs []org.Org_WC
	var org_wc org.Org_WC
	if orgs, err := org.ListOrganizationsForWC(); err != nil {
		glog.Errorln("Orgnization ListHandler ListOrgConfig", err.Error())
		httputil.Response(rw, 400, err)
		return
	} else {
		for _, org := range orgs {
			org_wc.Name = org.Name
			org_wc.OrgCode = org.Code
			org_wc.ImageUrl = org.ImageUrl
			org_wc.DetailsUrl = org.DetailsUrl
			org_wcs = append(org_wcs, org_wc)
		}
	}
	httputil.ResponseJson(rw, 200, &org_wcs)
	return
}

//todo 这块直接用的是pinto的接口，最后要改
func GetListCheckupsHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	checkups, err := pinto.ListCheckups(db.GetDB())
	if err != nil {
		glog.Errorln("Orgnization GetListCheckupsHandler ListCheckups err ", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	httputil.ResponseJson(rw, 200, checkups)
}

func GetPlansHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var plans []appointment.Plan
	var err error
	if plans, err = appointment.GetPlans(); err == nil {
		if err = json.NewEncoder(rw).Encode(&plans); err != nil {
			glog.Errorln("appoint.GetPlansHandler", err.Error())
			httputil.ResponseJson(rw, 400, "not found")
		}
	}
	return
}

func GetBannersHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var banners []appointment.Banner
	var err error
	if banners, err = appointment.GetBanners(); err == nil {
		httputil.ResponseJson(rw, 200, &banners)
		return
	}
	glog.Errorln("appoint.GetBannersHandler Err  ", err.Error())
	httputil.ResponseJson(rw, 400, err)
}

func GetOffDayHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	org_code := ps.ByName("code")
	offdays, err := appointment.GetOffDay(org_code)
	if err != nil {
		glog.Errorln("appoint.GetOffDayHandler Err ", err.Error())
	}
	if err = json.NewEncoder(rw).Encode(offdays); err != nil {
		httputil.ResponseJson(rw, 400, err)
	}
	return
}
