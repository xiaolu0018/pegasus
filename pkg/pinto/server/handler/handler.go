package handler

import (
	"net/http"
	"encoding/json"

	"bjdaos/pegasus/pkg/common/api/pinto"
	"bjdaos/pegasus/pkg/pinto/server/db"

	"github.com/golang/glog"
	httputil "github.com/1851616111/util/http"
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
