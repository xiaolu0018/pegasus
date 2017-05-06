//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：自定义日志接口
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package log

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/glog"

	"bjdaos/pegasus/pkg/instrument/heart/env"
	"bjdaos/pegasus/pkg/instrument/heart/types"
)

func WriteLog(logtype, errcontent string) int {
	if logtype != "Info" && logtype != "Debug" && logtype != "Err" {
		glog.Error("this is not a logtype ")
		return types.WriteLogErr
	}
	data := time.Now().Format("20060102")
	logFilea := env.GetConLogPath()
	logPath := logFilea + "/" + data + "_" + logtype + ".log"
	errcontent = "[" + errcontent + "]"
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("open file error=%s\r\n", err.Error())
		return types.WriteLogErr
	}
	logger := log.New(logFile, "{"+logtype+"} ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(errcontent)
	return types.WriteLogSucc
}
