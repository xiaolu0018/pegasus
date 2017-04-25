package image

import (
	"fmt"
	httputil "github.com/1851616111/util/http"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const rootpath string = "/home/administration/" //将文件储存的位置 ,最后必须带 /

func SaveImageHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := r.ParseMultipartForm(int64(10 << 20)); err != nil { //暂时将最大可存储大小设置为10M
		httputil.ResponseJson(rw, 400, fmt.Errorf("file too large", err.Error()))
		return
	}

	var file multipart.File
	var fh *multipart.FileHeader
	var err error
	defer file.Close()

	if file, fh, err = r.FormFile("uploadfile"); err != nil {
		glog.Errorf("uploadfile err", err.Error())
		httputil.ResponseJson(rw, 400, "paras invlid")
		return
	}
	//临时文件存储路径， 创建文件时以存储当天作为其文件夹，避免过多的重名文件
	//tempath := filepath.Clean(rootpath + time.Now().Format("2006-01-02"))

	//if b, err := util.PathExist(tempath); !b && err == nil { //不存在
	//if err = os.Mkdir(tempath, 0666); err != nil {
	//	glog.Errorf("make direct err", err.Error())
	//	httputil.ResponseJson(rw, 400, "paras invlid")
	//	return
	//}
	//}
	tempath := filepath.Clean(rootpath + fmt.Sprintf("%d", time.Now().Unix()) + fh.Filename)
	if f, err := os.OpenFile(tempath, 1, 0666); err != nil {
		glog.Errorf("openfile err", err.Error())
		httputil.ResponseJson(rw, 400, "paras invlid")
		f.Close()
		return

	} else {
		if _, err = io.Copy(f, file); err != nil {
			glog.Errorf("copyfile err", err.Error())
			httputil.ResponseJson(rw, 400, "paras invlid")
			f.Close()
			return
		} else {
			httputil.ResponseJson(rw, 200, tempath)
			f.Close()
			return
		}
	}
}
