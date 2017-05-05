//==================================================================
//创建时间：2017-4-26 首次创建
//功能描述：调用公司成熟的APi接口获取数据
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package api

import (
	"time"

	"net/http"
	"net/url"
	"strings"

	"encoding/json"
	"github.com/golang/glog"

	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/types"
)

var BaseUrl = ""

//获取操作医生信息
func Login(username string, password string) (types.Manager, error) {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var err error
	var req *http.Request
	var rsp *http.Response
	loginMap := make(map[string]interface{})
	//loginData := make([]string, 0, 3)
	if req, err = http.NewRequest("POST", BaseUrl+"/ma_public/third/rest/operator/validateManager/"+username+"/"+password, nil); err != nil {
		glog.Errorln("newrequest err", err)
		return types.Manager{}, err
	}
	req.Header.Set("serviceCode", "service-1611210001")
	req.Header.Set("serviceAuthorizationCode", "647D42A4CC9CBB7B0EE2B154C28FC043")
	if rsp, err = client.Do(req); err != nil {
		glog.Errorln("newrequest err", err)
		return types.Manager{}, err
	}
	defer rsp.Body.Close()
	json.NewDecoder(rsp.Body).Decode(&loginMap)
	mapItem, ok := loginMap["item"].(map[string]interface{})
	if !ok {
		glog.Errorln("mapItem err", err)
		return types.Manager{}, err
	}
	account := mapItem["account"].(string)
	userName := mapItem["username"].(string)
	managerCode := mapItem["managerCode"].(string)
	hosCode := mapItem["hosCode"].(string)
	orgCode := mapItem["orgCode"].(string)
	iRet := types.Manager{
		AccountName: account,
		UserName:    userName,
		ManagerCode: managerCode,
		HosCode:     hosCode,
		OrgCode:     orgCode,
	}
	return iRet, nil
}

//获取病号信息
func GetPersonInfo(queryCode string) (types.Person, error) {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response
	var err error
	personMap := make(map[string]interface{})
	if req, err = http.NewRequest("POST", BaseUrl+"/ma_public/third/rest/examperson/getPersonInfo/"+queryCode, nil); err != nil {
		glog.Errorln("newrequest err", err)
		return types.Person{}, err
	}
	req.Header.Set("serviceCode", "service-1611210001")
	req.Header.Set("serviceAuthorizationCode", "647D42A4CC9CBB7B0EE2B154C28FC043")
	if rsp, err = client.Do(req); err != nil {
		glog.Errorln("newrequest err", err)
		return types.Person{}, err
	}
	defer rsp.Body.Close()
	json.NewDecoder(rsp.Body).Decode(&personMap)
	mapItem, ok := personMap["item"].(map[string]interface{})
	if !ok {
		glog.Errorln("mapItem err", err)
		return types.Person{}, err
	}
	examinationNo := mapItem["examinationNo"].(string)
	name := mapItem["name"].(string)
	sex := mapItem["sex"].(interface{})
	age := mapItem["age"].(interface{})
	checkupDate := mapItem["checkupdate"].(string)
	iRet := types.Person{
		Examination_no: examinationNo,
		Name:           name,
		Sex:            sex,
		Age:            age,
		CheckupDate:    checkupDate,
	}
	return iRet, nil
}

//更改检查状态
func UpdateCheckStatus4Departments(Ex_No string, checkCode string, MangerCode string, dManagerCode string, status string) int {
	var req *http.Request
	var rsp *http.Response
	var err error
	reqParams := url.Values{}
	reqParams.Add("examinationNo", Ex_No)
	reqParams.Add("checkupCode", checkCode)
	reqParams.Add("checkMangerCode", MangerCode)
	reqParams.Add("diagnoseManagerCode", dManagerCode)
	reqParams.Add("checkStatus", status)

	if req, err = http.NewRequest("POST", BaseUrl+"/ma_public/third/rest/examperson/updateCheckStatus4Departments", strings.NewReader(reqParams.Encode())); err != nil {
		glog.Errorln("newrequest err", err)
		return types.UpdateCheckStatus4DepartmentsErr
	}
	req.Header.Set("serviceCode", "service-1611210001")
	req.Header.Set("serviceAuthorizationCode", "647D42A4CC9CBB7B0EE2B154C28FC043")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if rsp, err = http.DefaultClient.Do(req); err != nil {
		glog.Errorln("UpdateCheckStatus4Departments err", err)
		return types.UpdateCheckStatus4DepartmentsErr
	}
	defer rsp.Body.Close()
	//result:= make(map[string]interface{})
	//json.NewDecoder(rsp.Body).Decode(&result)
	//fmt.Println("status",rsp.Status)
	//fmt.Println(result)
	return types.UpdateCheckStatus4DepartmentsSucc
}

//操作结果数据
func UpdateCheckupResult(Ex_No string, checkCode string, MangerCode string, dManagerCode string, iRet string, itemvalues string, img string) int {
	var req *http.Request
	var rsp *http.Response
	var err error
	reqParams := url.Values{}
	reqParams.Add("examinationNo", Ex_No)
	reqParams.Add("checkupCode", checkCode)
	reqParams.Add("checkMangerCode", MangerCode)
	reqParams.Add("diagnoseManagerCode", dManagerCode)
	reqParams.Add("diagnoseResult", iRet)
	reqParams.Add("itemValues", itemvalues)
	reqParams.Add("images", img)
	if req, err = http.NewRequest("POST", BaseUrl+"/ma_public/third/rest/examperson/updateCheckupResult", strings.NewReader(reqParams.Encode())); err != nil {
		glog.Errorln("newrequest err", err)
		return types.UpdateCheckupResultErr
	}
	req.Header.Set("serviceCode", "service-1611210001")
	req.Header.Set("serviceAuthorizationCode", "647D42A4CC9CBB7B0EE2B154C28FC043")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if rsp, err = http.DefaultClient.Do(req); err != nil {
		glog.Errorln("UpdateCheckupResult err", err)
		return types.UpdateCheckupResultErr
	}
	defer rsp.Body.Close()
	//result:= make(map[string]interface{})
	//json.NewDecoder(rsp.Body).Decode(&result)
	//fmt.Println("status",rsp.Status)
	//fmt.Println(result)
	return types.UpdateCheckupResultErrSucc
}

/*
func GetQueue() {
	var   req   *http.Request
	var   rsp   *http.Response
	var   err   error
	queueMap := make(map[string]interface{})
	//reqParams := url.Values{}
	//reqParams.Add("examinationNo", examination_no)

	if req, err = http.NewRequest("POST", "http://10.1.0.190:8080/autoaim/autoaim/queue?examinationNo='0001160001045'", nil); err != nil {
		glog.Errorln("newrequest err", err)
	}
	//req.Header.Set("serviceCode", "service-1611210001")
	//req.Header.Set("serviceAuthorizationCode", "647D42A4CC9CBB7B0EE2B154C28FC043")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if rsp, err = http.DefaultClient.Do(req); err != nil {
		glog.Errorln("getQueue err", err)
	}
	defer rsp.Body.Close()
	json.NewDecoder(rsp.Body).Decode(&queueMap)
	//mapItem, ok  := queueMap["item"].(map[string]interface{})
	//if !ok {
	//	glog.Errorln("mapItem err", err)
	//}
	fmt.Println(rsp.Status)
	fmt.Println(queueMap)
	fmt.Println(rsp.Body)
}
*/
