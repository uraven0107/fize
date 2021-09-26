package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ResolveParentDirPath(t *testing.T) {
	t.Run("GetParentDirPath", func(*testing.T) {
		assert := assert.New(t)
		assert.Equal("/home", ResolveParentDirPath("/home/hoge"))
		assert.Equal("/", ResolveParentDirPath("/home"))
		assert.Equal("/", ResolveParentDirPath("/"))
	})
}
