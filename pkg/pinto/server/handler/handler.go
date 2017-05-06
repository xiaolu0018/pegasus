package handler

import (
	"encoding/json"
	"net/http"

	"bjdaos/pegasus/pkg/common/api/pinto"
	"bjdaos/pegasus/pkg/pinto/server/db"

	httputil "github.com/1851616111/util/http"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
)

func CreateBookRecordHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		glog.Errorf("pinto.handler CreateBookRecord Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	bookRecord := pinto.MapToBookRecord(result)

	err := pinto.InsertBookRecord(db.GetWriteDB(), &bookRecord)
	if err != nil {
		glog.Errorf("pinto.handler CreateBookRecord InsertBookRecord params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	data := map[string]string{}
	data["bookno"] = bookRecord.BookNo
	httputil.ResponseJson(rw, 200, data)
}

func CreateExamsHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		glog.Errorf("pinto.handler CreateExamsHandler Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	exam, exam_checkups, exam_sales, br, person := pinto.MapToExams(result)
	err := pinto.SaveExaminations(db.GetWriteDB(), exam, exam_checkups, exam_sales, br, person)
	if err != nil {
		glog.Errorf("pinto.handler CreateExamsHandler SaveExaminations  err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	data := map[string]string{}
	data["bookno"] = br.BookNo
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
