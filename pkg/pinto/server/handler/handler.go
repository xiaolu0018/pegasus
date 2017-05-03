package handler

import (
	"192.168.199.199/bjdaos/pegasus/pkg/common/api/pinto"
	"192.168.199.199/bjdaos/pegasus/pkg/pinto/server/db"
	"encoding/json"
	httputil "github.com/1851616111/util/http"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CreateBookRecordHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		glog.Errorf("pinto.handker CreateBookRecord Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	bookRecord := pinto.MapToBookRecord(result)

	err := pinto.InsertBookRecord(db.GetWriteDB(), &bookRecord)
	if err != nil {
		glog.Errorf("pinto.handker CreateBookRecord InsertBookRecord params err %v\n", err.Error())
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
		glog.Errorf("pinto.handker CreateExamsHandler Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	exam, exam_checkups, exam_sales, br, person := pinto.MapToExams(result)
	err := pinto.SaveExaminations(db.GetWriteDB(), exam, exam_checkups, exam_sales, br, person)
	if err != nil {
		glog.Errorf("pinto.handker CreateExamsHandler SaveExaminations  err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	data := map[string]string{}
	data["bookno"] = br.BookNo
	httputil.ResponseJson(rw, 200, data)

}
