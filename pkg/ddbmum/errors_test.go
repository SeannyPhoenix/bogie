package ddbmum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsupportedTypeError(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	err := &UnsupportedTypeError{}
	assert.Equal("unsupported type: invalid", err.Error())
}
