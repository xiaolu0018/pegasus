//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：xml解析处理
//创建人：张志浩
//修改记录：若要修改请记录 郭世江去除以前的方法重写
//==================================================================
package parsexml

import (
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/types"
	"encoding/xml"
	"github.com/golang/glog"
	"io/ioutil"
)

func ParseXml(xmlPath string, xmlName string) (types.Result, error) {
	result := types.Result{}
	buffer, err := ioutil.ReadFile(xmlPath + xmlName)
	if err != nil {
		glog.Errorln("ParseXml err", err)
		return types.Result{}, nil
	}
	xml.Unmarshal(buffer, &result)
	return result, nil
}
