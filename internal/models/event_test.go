package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewEventDocument(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	u := uuid.New()
	r := NewEventDocuemnt(&u)

	assert.Equal(DocTypeEvent, r.Type)
	assert.Equal(&u, r.User)
}
