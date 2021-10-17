package fs

import (
	"os"
)

func GetFileInfosUnderDir(dirPath string) ([]os.FileInfo, error) {
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	fileInfos := make([]os.FileInfo, len(dirEntries))
	for i, dirEntry := range dirEntries {
		fileInfo, err := dirEntry.Info()
		if err != nil {
			return nil, err
		} else {
			fileInfos[i] = fileInfo
		}
	}
	return fileInfos, nil
}
