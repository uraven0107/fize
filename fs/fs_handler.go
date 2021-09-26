package fs

import (
	"io/ioutil"
	"os"
)

func GetFileInfosUnderDir(dirPath string) ([]os.FileInfo, error) {
	fileInfos, err := ioutil.ReadDir(dirPath)
	return fileInfos, err
}
