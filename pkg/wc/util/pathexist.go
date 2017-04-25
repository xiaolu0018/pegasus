package util

import (
	"os"
	"path/filepath"
)

//判断文件或文件夹是否存在

func PathExist(path string) (bool, error) {
	_, err := os.Stat(filepath.Clean(path))
	if err == nil {
		return true, nil //存在
	}

	if os.IsNotExist(err) {
		return false, nil //不存在
	}
	return false, err
}
