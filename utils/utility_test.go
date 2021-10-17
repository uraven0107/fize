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

func Test_IsAccessableDir(t *testing.T) {
	t.Run("CanJudgeIsAssesable", func(*testing.T) {
		assert := assert.New(t)

		b1, err := IsAccessableDir("/home/uraven/.zshrc")
		assert.NotNil(err)
		assert.False(b1)

		b2, err := IsAccessableDir("/home/uraven")
		assert.Nil(err)
		assert.True(b2)

		b3, err := IsAccessableDir("/root")
		assert.Nil(err)
		assert.False(b3)

	})
}
