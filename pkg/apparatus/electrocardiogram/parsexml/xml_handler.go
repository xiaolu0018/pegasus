//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：文件处理
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package parsexml

import (
	"fmt"
	"os"
	"regexp"

	"io/ioutil"
	"strings"

	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/model"
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/types"
	"github.com/golang/glog"
)

//遍历一个目录下的所有文件和目录,该函数只在本文件中使用
func getListDir(dirPth string) (currentFiles []string, nextDirfiles []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			nextDirfiles = append(nextDirfiles, dirPth+PthSep+fi.Name())
			getListDir(dirPth + PthSep + fi.Name())
		} else {
			currentFiles = append(currentFiles, dirPth+PthSep+fi.Name())
		}
	}
	return currentFiles, nextDirfiles, nil
}

//根据体检单号获取图片的名字
func GetImgNameActExNo(path string, queryCode string, str string) string {
	files, _, _ := getListDir(path)
	fileLen := len(files)
	var imgName string
	reg_front := regexp.MustCompile("\\d{8,13}")
	reg_end := regexp.MustCompile("\\d{14}")
	if str == ".jpg" {
		for i := 0; i < fileLen; i++ {
			data_front := reg_front.FindString(files[i])
			date_end := reg_end.FindString(files[i])
			imgNameReg := data_front + "_" + date_end
			if strings.Contains(imgNameReg, queryCode) {
				imgName = imgNameReg + str
			}
		}
	}
	return imgName
}

//获取指定目录下一个文件和截取体检单号或者腕表号
func getAllFileName(path string, str string) (int, []string, []string) {
	files, _, _ := getListDir(path)
	fileLen := len(files)
	fileSlice := make([]string, 0, fileLen)
	queryCode := make([]string, 0, fileLen)
	reg_front := regexp.MustCompile("\\d{8,16}")
	reg_end := regexp.MustCompile("\\d{14}.xml")
	if str == ".xml" {
		for i := 0; i < fileLen; i++ {
			data_front := reg_front.FindString(files[i])
			date_end := reg_end.FindString(files[i])
			imgName := data_front + "_" + date_end
			fileSlice = append(fileSlice, imgName)
			queryCode = append(queryCode, data_front)
		}
	}
	return fileLen, queryCode, fileSlice
}

//删除指定名字的文件
func RemoveFile(path string, name string) int {
	err := os.Remove(path + name)
	if err != nil {
		glog.Errorln("RemoveFile err", err)
		return types.RemoveFileErr
	}
	return types.RemoveFileSucc
}

func XmlHearData(filePath string, fileType string) int {
	//获取XML文件
	fileLen, queryCode, fileSlice := getAllFileName(filePath, fileType)
	if fileLen == 0 {
		glog.Errorf("XmlHearData err")
	}
	fmt.Println(fileSlice)
	//解析XML文件
	for i := 0; i < fileLen; i++ {
		iRet, err := ParseXml(filePath, fileSlice[i])
		if err != nil {
			glog.Errorf("ParseXml fail")
		}
		//将数据插入到表heart_data中
		SuccCode := model.InsertXmlDataToDB(queryCode[i], iRet)
		if SuccCode != 2007 {
			glog.Errorf("InsertXmlDataToDB fail")
		}
		//删除文件
		removeCode := RemoveFile(filePath, fileSlice[i])
		if removeCode != 2003 {
			glog.Errorf("RemoveFile fail")
		}
	}
	return types.XmlHearDataSucc
}
