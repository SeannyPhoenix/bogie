package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {

	t.Parallel()
	assert := assert.New(t)

	v := 42
	ptr := Ptr(v)
	assert.Equal(&v, ptr)
}

func TestNilPtr(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	ptr := NilPtr[int]()
	assert.Nil(ptr)
}
