package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	httputil "github.com/1851616111/util/http"
	"github.com/1851616111/util/message"
	"github.com/golang/glog"

	"github.com/julienschmidt/httprouter"

	"192.168.199.199/bjdaos/pegasus/pkg/reporter/model"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/util"
	//"192.168.199.199/bjdaos/pegasus/pkg/common/util/safe"
)

func AuthHandler(handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//user, passwd, ok := r.BasicAuth()
		//if ! ok {
		// httputil.Response(w, 401, http.StatusText(401))
		// return
		//}
		r.ParseForm()
		url_ := r.Form
		user := url_.Get("username")
		passwd := url_.Get("password")

		if hosCode, err := model.Auth(user, passwd); err != nil {
			glog.Errorf("auth reporter err %v\n", err)
			httputil.Response(w, 401, http.StatusText(401))
			return
		} else {
			handler(w, r, util.AddParam(ps, "hos_code", fmt.Sprintf("%s", hosCode)))
		}
	}
}

func GetReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ex_no := r.FormValue("examination_no")
	if ex_no == "" {
		message.ParamNotFound(w, "examination_no")
		return
	}

	//status, err := model.GetExaminationStatus(ex_no)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		httputil.Response(w, 404, "Not Found")
	//	} else {
	//		httputil.Response(w, 404, err)
	//	}
	//	return
	//}
	//
	//var ifSync bool
	//switch status {
	//case 1999, 1123, 1120, 1112, 1111, 1110, 1100, 1090, 1080:
	//	//ifSync = false
	//	ifSync = true
	//default:
	//	ifSync = true
	//}

	report, err := model.GetReporterByExNo(ex_no, true)
	if err != nil {
		if err == sql.ErrNoRows {
			httputil.Response(w, 404, "Not Found")
		} else {
			httputil.Response(w, 400, err)
		}
		return
	}
	//b, err := json.Marshal(report)
	//if err != nil {
	//	httputil.Response(w, 400, err)
	//	return
	//}

	//b, err = safe.RsaEncrypt(b)
	//if err != nil {
	//	httputil.Response(w, 400, err)
	//	return
	//}
	if err := json.NewEncoder(w).Encode(report); err != nil {
		glog.Error("[Handler] GetReport err %v\n", err)
		message.InnerError(w)
	}
	return

}

func ReportListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page_no := r.FormValue("page_no")

	if page_no == "" {
		page_no = "1"
	}

	var page_index int
	var err error
	if page_index, err = strconv.Atoi(page_no); err != nil {
		glog.Errorln("Orgnization ListHandler page_index", err.Error())
		httputil.Response(w, 400, err)
		return
	}

	var ifAlreadyReported bool = false
	if state_no := r.FormValue("state_no"); state_no == "1" {
		ifAlreadyReported = true
	}

	ex_no := r.FormValue("examination_no")
	name := r.FormValue("name")
	begintime := r.FormValue("begintime")
	endtime := r.FormValue("endtime")
	sex := r.FormValue("sex")

	Rets, err := model.GetQueryAll(page_index, ex_no, name, sex, ifAlreadyReported, begintime, endtime, ps.ByName("hos_code"))
	if err != nil {
		if err == sql.ErrNoRows {
			httputil.Response(w, 404, "Not Found")
		} else {
			httputil.Response(w, 400, err)
		}
		return
	}

	jsonerr := json.NewEncoder(w).Encode(Rets)
	if jsonerr != nil {
		glog.Errorf("[ReportListHandler] err  %v\n", jsonerr)
		message.InnerError(w)
	}
	return
}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ex_no, status := r.FormValue("examination_no"), r.FormValue("status")
	if len(ex_no) == 0 || len(ex_no) > 50 {
		httputil.Response(w, 400, "examination_no format wrong")
		return
	}
	if len(status) == 0 || len(status) > 50 {
		httputil.Response(w, 400, "status format wrong")
		return
	}

	if err := model.UpdateStatus(ex_no, status); err != nil {
		httputil.Response(w, 400, fmt.Sprintf("update status err %v\n", err))
		return
	}

	httputil.Response(w, 200, "ok")
}
