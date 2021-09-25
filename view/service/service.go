package service

import (
	"os"

	"github.com/uraven0107/fize/app/controller"
)

func FetchFileInfos(dirPath string) ([]os.FileInfo, error) {
	builder := controller.NewParameterBuilder()
	param := builder.Path(dirPath).Build()

	result := controller.GetFileInfos(param)
	if err := result.GetError(); err != nil {
		return nil, err
	}

	return result.GetFileInfos(), nil
}
