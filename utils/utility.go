package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const DS = string(filepath.Separator)

func ResolvePath(dir string, target string) string {
	return dir + DS + target
}

func ResolveParentDirPath(current string) string {
	return filepath.Dir(current)
}

func IsAccessableDir(dirPath string) (bool, error) {
	fileInfo, e := os.Stat(dirPath)
	if e != nil {
		return false, e
	}

	if !fileInfo.IsDir() {
		return false, errors.New(fmt.Sprintf("Arg is not directory. %v", fileInfo))
	}

	_, err := os.ReadDir(dirPath)
	if err != nil {
		if os.IsPermission(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}
