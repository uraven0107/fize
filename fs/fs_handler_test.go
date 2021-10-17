package fs

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uraven0107/fize/utils"
)

const TEST_DIR = "." + utils.DS + "test"

func Test_GetFileInfosUnderDir(t *testing.T) {
	t.Run("NotExistDirReturnError", func(t *testing.T) {
		assert := assert.New(t)
		_, err := GetFileInfosUnderDir("./notexist")
		assert.EqualError(err, "open ./notexist: no such file or directory")
	})

	t.Run("PathForFileReturnError", func(t *testing.T) {
		filePath := TEST_DIR + utils.DS + "test1"
		assert := assert.New(t)
		_, err := GetFileInfosUnderDir(filePath)
		assert.EqualError(err, "readdirent "+filePath+": not a directory")
	})

	t.Run("CanGetFileInfos", func(t *testing.T) {
		filePath := TEST_DIR

		f, err := os.ReadDir(filePath)
		if err != nil {
			fmt.Printf("⚠️ Error has occured at os.ReadDir(%v)!!! error = %v", filePath, err)
			return
		}
		expected := len(f)

		assert := assert.New(t)
		fileInfos, _ := GetFileInfosUnderDir(filePath)

		assert.Equal(expected, len(fileInfos), "🚨 Length of []os.FileInfo doesn't equal!!!")

	})

	t.Run("FileInfoShouldn'tNil", func(t *testing.T) {
		assert := assert.New(t)
		filePath := TEST_DIR
		fileInfos, err := GetFileInfosUnderDir(filePath)
		assert.Nil(err)
		for _, fileInfo := range fileInfos {
			assert.NotNil(fileInfo)
		}
	})
}
