package utils

import (
	"path/filepath"
	"strings"
)

const DS = string(filepath.Separator)

func ResolvePath(dir string, target string) string {
	return dir + DS + target
}

func ResolveRootDirPath(current string) string {
	dirs := strings.Split(current, DS)
	var dirPath = ""
	for i := 0; i < len(dirs)-1; i++ {
		dirPath = "/" + dirs[i]
	}
	if dirPath == "" {
		dirPath = "/"
	}
	return dirPath
}
