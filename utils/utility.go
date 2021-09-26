package utils

import (
	"path/filepath"
)

const DS = string(filepath.Separator)

func ResolvePath(dir string, target string) string {
	return dir + DS + target
}

func ResolveParentDirPath(current string) string {
	return filepath.Dir(current)
}
