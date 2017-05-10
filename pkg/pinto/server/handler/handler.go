package handler

import (
	"encoding/json"
	"net/http"

	"bjdaos/pegasus/pkg/common/api/pinto"
	"bjdaos/pegasus/pkg/pinto/server/db"

	httputil "github.com/1851616111/util/http"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"time"
)

func BookPersonHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := pinto.Appointment{}
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		glog.Errorf("pinto.handler CreateBookRecord Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	result.TimeNow = time.Now()

	bookRecord := pinto.FilterBookRecordByAppoint(&result)

	err := pinto.InsertBookRecord(db.GetWriteDB(), bookRecord)
	if err != nil {
		glog.Errorf("pinto.handler CreateBookRecord InsertBookRecord params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	data := map[string]string{}
	data["bookno"] = bookRecord.BookNo
	httputil.ResponseJson(rw, 200, data)
}

func BookPlanHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := pinto.Appointment{}
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		glog.Errorf("pinto.handler CreateExamsHandler Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	result.TimeNow = time.Now()
	e_all, err := pinto.FilterExamsAll(db.GetWriteDB(), &result)
	if err != nil {
		glog.Errorf("pinto.handler FilterExamsAll err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	err = pinto.SaveExaminations(db.GetWriteDB(), e_all)
	if err != nil {
		glog.Errorf("pinto.handler CreateExamsHandler SaveExaminations  err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	data := map[string]string{}
	data["bookno"] = e_all.B.BookNo
	httputil.ResponseJson(rw, 200, data)

}

func CancelExamHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		glog.Errorf("pinto.handler CancelExamHandler Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	var bookno string
	var withplan bool

	if bookno_, ok := result["bookno"]; ok {
		bookno = bookno_.(string)
	}

	if withplan_, ok := result["withplan"]; ok {
		withplan = withplan_.(bool)
	}
	var err error
	if withplan && bookno != "" {
		err = pinto.UpdateBookRecordWithExamToInvalid(db.GetWriteDB(), bookno)
	} else if !withplan && bookno != "" {
		err = pinto.UpdateBookRecordToInvalid(db.GetWriteDB(), bookno)
	}
	if err != nil {
		glog.Errorf("pinto.handler CancelExamHandler update err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	httputil.ResponseJson(rw, 200, nil)
	return
}

func GetExamStatusHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := make(map[string][]string)
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		glog.Errorf("pinto.handler CreateExamsHandler Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	mapresult := GetExaminationStatus(result)
	httputil.ResponseJson(rw, 200, mapresult)
	return
}
