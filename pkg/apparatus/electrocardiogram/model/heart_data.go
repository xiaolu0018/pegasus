//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：获取数据
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package model

import (
	"fmt"
	"github.com/golang/glog"

	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/api"
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/types"
)

func GetHeartData(queryCode string) ([]types.HeartData, error) {
	var heartData []types.HeartData
	hrows, herr := DB.Query(fmt.Sprintf(`SELECT watch_mark, examination_no, heart_rate, time_pd_rr, time_pd_pr, qrs_front, pd_qrs, pd_qt, pd_qtc, wave_rs, wave_rv5, wave_rv6, heart_analysis, gread, discode, miscode, confirm, annotation, result
		FROM heart_data WHERE examination_no='%s'`, queryCode))
	if herr != nil {
		glog.Errorf("GetPersonHeartData: sql return err %v\n", herr)
		return nil, herr
	}
	for hrows.Next() {
		var hrow = new(types.HeartData)
		if herr = hrows.Scan(&hrow.Watch_mark, &hrow.Examination_no, &hrow.Heart_rate, &hrow.Time_pd_rr, &hrow.Time_pd_pr, &hrow.Qrs_front, &hrow.Pd_qrs, &hrow.Pd_qt, &hrow.Pd_qtc,
			&hrow.Wave_rs, &hrow.Wave_rv5, &hrow.Wave_rv6, &hrow.Heart_analysis, &hrow.Gread, &hrow.Discode, &hrow.Miscode, &hrow.Confirm, &hrow.Annotation, &hrow.Result); herr != nil {
			return nil, herr
		}
		heartData = append(heartData, *hrow)
	}
	if hrows.Err() != nil {
		glog.Errorf("GetPersonHeartData: sql rows.Err() %v\n", herr)
		return nil, hrows.Err()
	}
	return heartData, nil
}

func GetPersonHeartData(queryCode string) (*types.Data, error) {
	var personData types.Person
	var heartData []types.HeartData
	personData, err1 := api.GetPersonInfo(queryCode)
	if err1 != nil {
		glog.Errorf("GetPersonHeartData: err %v\n", err1)
		return nil, err1
	}
	heartData, err := GetHeartData(queryCode)
	if err != nil {
		glog.Errorf("GetPersonHeartData: err %v\n", err)
		return nil, err
	}
	iRets := &types.Data{
		Personinfo: personData,
		Heartdata:  heartData,
	}
	return iRets, nil
}
