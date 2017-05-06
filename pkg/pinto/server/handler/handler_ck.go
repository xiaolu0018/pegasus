package handler

import (
	"bjdaos/pegasus/pkg/common/api/pinto"
	"bjdaos/pegasus/pkg/reporter/db"
	"encoding/json"
	httputil "github.com/1851616111/util/http"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetCheckupCodesBySaleCodesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	m := map[string][]string{}
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	if _, exist := m["salecodes"]; !exist {
		httputil.Response(w, 400, "req param salecodes not found")
		return
	}

	cks, err := pinto.GetCheckupCodesBySaleCodes(db.GetReadDB(), m["salecodes"])
	if err != nil {
		httputil.Response(w, 400, err)
		return

	}
	httputil.ResponseJson(w, 200, map[string][]string{"salecodes": cks})
	return
}
