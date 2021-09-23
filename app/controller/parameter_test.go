package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParameterBuilder(t *testing.T) {
	t.Run("CanBuileParameters", func(t *testing.T) {
		assert := assert.New(t)

		builder := NewParameterBuilder()
		param := builder.Path("test").Build()

		assert.Equal("test", param.path, "🚨 Parameter.path hasn't expected!!!")
	})
}
