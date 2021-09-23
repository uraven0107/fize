package view

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewViewModel(t *testing.T) {
	t.Run("EmptyStringReturnError", func(t *testing.T) {
		assert := assert.New(t)
		emptyPath := ""
		_, err := NewViewModel(emptyPath)
		assert.EqualError(err, "🚨 Empty string is not allowed in the path string")
	})
}

func Test_Init(t *testing.T) {
	t.Run("CanInitViewModel", func(t *testing.T) {
		assert := assert.New(t)
		vm, err := NewViewModel(".")
		assert.Nil(err, fmt.Sprintf("⚠️ NewViewModel() return error!!! error = %v", err))
		f, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Printf("⚠️ Error has occured at ioutil.ReadDir(%v)!!! error = %v", ".", err)
		}
		expected := len(f)
		assert.Equal(expected, len(vm.fileInfos), "🚨 Length of ViewModel.fileInfos hasn't expected!!!")
		assert.Equal(".", vm.dirPath, "🚨 ViewModel.dirPath hasn")
	})
}
