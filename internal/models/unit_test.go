package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnitRecord(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := NewUnitRecord()

	assert.Equal(DocTypeUnit, r.Type)
	assert.Nil(r.User)
}
