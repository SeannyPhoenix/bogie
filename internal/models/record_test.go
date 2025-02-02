package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewRecord(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := newRecord(DocTypeEvent, nil)

	assert.NotEqual(r.Id, uuid.Nil)
	assert.Equal(DocTypeEvent, r.Type)
	assert.Equal(DocStatusActive, r.Status)
	assert.NotNil(r.CreatedAt)
	assert.NotNil(r.UpdatedAt)
	assert.Nil(r.User)
}
