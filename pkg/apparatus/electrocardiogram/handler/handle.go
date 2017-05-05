//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述:处理QT前端发过来的数据
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package handler

import (
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/api"
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/model"
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/parsexml"
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/types"
	"encoding/json"
	"github.com/1851616111/util/message"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

var ImgsPathUrl = ""
var XmlPathUrl = ""

func GetAccountInfo(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	name := r.FormValue("username")
	password := r.FormValue("password")
	if name == "" && password == "" {
		glog.Errorf("name and password is nil")
		return
	}
	ManageInfo, err := api.Login(name, password)
	if err != nil {
		glog.Errorf("Login err")
		return
	}
	if err := json.NewEncoder(w).Encode(ManageInfo); err != nil {
		message.InnerError(w)
	}
	return
}

func GetQueryCode(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	queryCode := r.FormValue("queryCode")
	if queryCode == " " {
		glog.Errorf("queryCode is nil")
		return
	}
	Ex_No_s, err := api.GetPersonInfo(queryCode)
	if err != nil {
		glog.Errorf("Login err")
		return
	}
	if err := json.NewEncoder(w).Encode(Ex_No_s); err != nil {
		message.InnerError(w)
	}
	return
}

func GetHeartData(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//homePath := os.Getenv("HOME")
	xmlNum := parsexml.XmlHearData(XmlPathUrl, ".xml")

	if xmlNum != 2008 {
		glog.Errorf("value is nil")
		return
	}
	Ex_No := r.FormValue("examination_no")
	if Ex_No == "" {
		glog.Errorf("value is nil")
		return
	}

	phData, err := model.GetPersonHeartData(Ex_No)
	if err != nil {
		glog.Errorf("GetPersonHeartData err")
		return
	}
	if err := json.NewEncoder(w).Encode(phData); err != nil {
		message.InnerError(w)
	}
	return
}

func UpdateResult(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	Ex_No := r.FormValue("examination_no")
	checkCode := r.FormValue("checkCode")
	MangerCode := r.FormValue("MangerCode")
	dManagerCode := r.FormValue("dManagerCode")
	iResult := r.FormValue("diagnoseResult")

	//心电相关参数
	time_pdpr := r.FormValue("timepdpr")   //P波
	pd_qrs := r.FormValue("pdqrs")         //QRS波
	pd_qt := r.FormValue("pdqt")           //ST段T波
	pd_qtc := r.FormValue("pdqtc")         //Q—T间期
	qrs_front := r.FormValue("qrsfront")   //心电轴
	heart_rate := r.FormValue("heartrate") //房性心率
	if Ex_No == "" && checkCode == "" && MangerCode == "" && dManagerCode == "" && iResult == "" {
		glog.Errorf("examination_no, checkCode, MangerCode and diagnoseResult is nil")
		return
	}
	errUcsd := api.UpdateCheckStatus4Departments(Ex_No, checkCode, MangerCode, dManagerCode, "1")
	if errUcsd == 1000 {
		glog.Errorf("UpdateCheckStatus4Departments is fail")
		return
	}
	homePath := os.Getenv("HOME")
	imgsName := parsexml.GetImgNameActExNo(homePath+ImgsPathUrl, Ex_No, "jpg")
	itemValues := "[{" + types.ItemCode + ":" + types.PCode + "," + types.ItemValue + ":" + time_pdpr + "}" + "," +
		"{" + types.ItemCode + ":" + types.QRSCode + "," + types.ItemValue + ":" + pd_qrs + "}" + "," +
		"{" + types.ItemCode + ":" + types.STTCode + "," + types.ItemValue + ":" + pd_qt + "}" + "," +
		"{" + types.ItemCode + ":" + types.QTCode + "," + types.ItemValue + ":" + pd_qtc + "}" + "," +
		"{" + types.ItemCode + ":" + types.HDZCode + "," + types.ItemValue + ":" + qrs_front + "}" + "," +
		"{" + types.ItemCode + ":" + types.HDWCode + "," + types.ItemValue + ":" + heart_rate + "}]"

	imageUrl := homePath + ImgsPathUrl + imgsName
	imgUrl := "\"" + imageUrl + "\""
	images := "[" + "{" + types.ImgUrl + ":" + imgUrl + "}" + "]"

	errUcr := api.UpdateCheckupResult(Ex_No, checkCode, MangerCode, dManagerCode, iResult, itemValues, images)
	if errUcr == 1001 {
		glog.Errorf("UpdateCheckupResult is fail")
		return
	}

	if err := json.NewEncoder(w).Encode("修改成功"); err != nil {
		message.InnerError(w)
	}
	return
}

/*
func GetNextRoom(){

}

func GetPersonCheckUpItem(){

}
*/
