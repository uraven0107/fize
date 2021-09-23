package view

import (
	"errors"
	"os"

	"github.com/uraven0107/fize/app/controller"
)

type RootViewModel struct {
	left  ViewModel
	right ViewModel
}

type ViewModel struct {
	dirPath   string
	fileInfos []os.FileInfo
}

func NewViewModel(dirPath string) (*ViewModel, error) {
	if dirPath == "" {
		return nil, errors.New("Empty string is not allowed in the path string")
	}

	vm := new(ViewModel)
	err := vm.init(dirPath)
	if err != nil {
		return nil, err
	}

	return vm, nil
}

func (vm *ViewModel) init(dirPath string) error {
	builder := controller.NewParameterBuilder()
	param := builder.Path(dirPath).Build()

	result := controller.GetFileInfos(param)
	if err := result.GetError(); err != nil {
		return err
	}

	vm.dirPath = dirPath
	vm.fileInfos = result.GetFileInfos()

	return nil
}
