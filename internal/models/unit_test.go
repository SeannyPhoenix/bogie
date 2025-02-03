package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnitDocuemnt(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := NewUnitDocument()

	assert.Equal(DocTypeUnit, r.Type)
	assert.Nil(r.User)
}
