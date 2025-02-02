package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewEventRecord(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	u := uuid.New()
	r := NewEventRecord(&u)

	assert.Equal(DocTypeEvent, r.Type)
	assert.Equal(&u, r.User)
}
