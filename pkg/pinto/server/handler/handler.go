package handler

import (
	"bjdaos/pegasus/pkg/common/api/pinto"
	"bjdaos/pegasus/pkg/pinto/server/db"
	"encoding/json"
	httputil "github.com/1851616111/util/http"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
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

func GetExaminationStatus(m map[string][]string) map[string]int {
	booknos, ok := m["booknos"]
	if !ok {
		return nil
	}

	if len(booknos) == 0 {
		return nil
	}

	rows, err := db.GetReadDB().Query(`SELECT b.bookno, COALESCE(e.status, 0)
	FROM book_record b LEFT JOIN examination e
	ON b.examination_no=e.examination_no
	WHERE b.bookno IN (%s)"`, strings.Join(booknos, ","))
	if err != nil {
		glog.Errorln("pinto.GetExaminationStatus GetExamStatusHandler err ", err)
		return nil
	}
	defer rows.Close()

	ret := map[string]int{}
	for rows.Next() {
		var bookno string
		var status int
		if err = rows.Scan(&bookno, &status); err != nil {
			return nil
		}

		ret[bookno] = status
	}

	return ret
}
