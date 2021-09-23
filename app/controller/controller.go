package controller

import (
	"os"

	"github.com/uraven0107/fize/fs"
)

type Result struct {
	err       error
	fileInfos []os.FileInfo
}

func newResult(err error, fileInfos []os.FileInfo) Result {
	return Result{
		err,
		fileInfos,
	}
}

func (result Result) HasError() bool {
	return result.err != nil
}

func (result Result) GetError() error {
	return result.err
}

func (result Result) GetFileInfos() []os.FileInfo {
	return result.fileInfos
}

func GetFileInfos(param Parameter) Result {
	dirPath := param.path

	fileInfos, err := fs.GetFileInfosUnderDir(dirPath)
	if err != nil {
		return newResult(err, nil)
	}

	return newResult(nil, fileInfos)
}
