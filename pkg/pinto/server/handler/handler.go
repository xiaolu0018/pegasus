package handler

import (
	"bjdaos/pegasus/pkg/common/api/pinto"
	"bjdaos/pegasus/pkg/pinto/server/db"
	"encoding/json"
	"fmt"
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
	result := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		glog.Errorf("pinto.handler CreateExamsHandler Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	mapresult := GetExaminationStatus(result)
	httputil.ResponseJson(rw, 200, mapresult)
	return
}

func GetExaminationStatus(result map[string]interface{}) []map[string]interface{} {
	var booknosStrs []string
	booknos, ok := result["bookno"]
	if ok {
		for _, val := range booknos.([]interface{}) {
			booknosStrs = append(booknosStrs, val.(string))
		}
	}

	if len(booknosStrs) == 0 {
		return nil
	}

	itmeStr := make([]string, len(booknosStrs))
	for k, salecode := range booknosStrs {
		itmeStr[k] = fmt.Sprintf(`'%s'`, salecode)
	}
	sqlStr := fmt.Sprintf("SELECT b.bookno,COALESCE(b.examination_no,''), COALESCE(e.status, 0) FROM book_record b LEFT JOIN examination e ON b.examination_no=e.examination_no WHERE b.bookno IN (%s)", strings.Join(itmeStr, ","))

	rows, err := db.GetReadDB().Query(sqlStr)
	if err != nil {
		glog.Errorln("pinto.GetExaminationStatus GetExamStatusHandler err ", err)
		return nil
	}

	results := make([]map[string]interface{}, 0)

	defer rows.Close()
	var bookno, exam_no string
	var status int
	for rows.Next() {
		if err = rows.Scan(&bookno, &exam_no, &status); err != nil {
			return nil
		}
		result_ := make(map[string]interface{})
		result_["bookno"] = bookno
		result_["examno"] = exam_no
		result_["status"] = status
		results = append(results, result_)
	}

	if rows.Err() != nil {
		glog.Errorln("pinto.GetExaminationStatus rows.Err() ", err)
		return nil
	}
	return results
}
