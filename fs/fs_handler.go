package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const DS = string(filepath.Separator)

func GetFileInfosUnderDir(dirPath string) ([]os.FileInfo, error) {
	fileInfos, err := ioutil.ReadDir(dirPath)
	return fileInfos, err
}
