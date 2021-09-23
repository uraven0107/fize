package fs

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_DIR = "." + DS + "test"

func Test_GetFileInfosUnderDir(t *testing.T) {
	t.Run("NotExistDirReturnError", func(t *testing.T) {
		assert := assert.New(t)
		_, err := GetFileInfosUnderDir("./notexist")
		assert.EqualError(err, "open ./notexist: no such file or directory")
	})

	t.Run("PathForFileReturnError", func(t *testing.T) {
		filePath := TEST_DIR + DS + "test1"
		assert := assert.New(t)
		_, err := GetFileInfosUnderDir(filePath)
		assert.EqualError(err, "readdirent "+filePath+": not a directory")
	})

	t.Run("CanGetFileInfos", func(t *testing.T) {
		filePath := TEST_DIR + DS + "test1"

		f, err := ioutil.ReadDir(filePath)
		if err != nil {
			fmt.Printf("⚠️ Error has occured at ioutil.ReadDir(%v)!!! error = %v", filePath, err)
			return
		}
		expected := len(f)

		assert := assert.New(t)
		fileInfos, err := GetFileInfosUnderDir(filePath)

		assert.Equal(expected, len(fileInfos), "🚨 Length of []os.FileInfo doesn't equal!!!")

	})
}
