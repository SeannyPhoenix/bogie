package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewRecord(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := newDocument(DocTypeEvent, nil)

	ca := r.CreatedAt

	assert.NotEqual(r.Id, uuid.Nil)
	assert.Equal(DocTypeEvent, r.Type)
	assert.Equal(DocStatusActive, r.Status)
	assert.NotNil(r.CreatedAt)
	assert.NotNil(r.UpdatedAt)
	assert.Equal(ca, r.UpdatedAt)
	assert.Nil(r.User)
}

func TestUpdateRecord(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := newDocument(DocTypeEvent, nil)
	r = updateDocuemnt(r)

	ca := r.CreatedAt

	assert.NotEqual(ca, r.UpdatedAt)
}

func TestDeactivateRecord(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := newDocument(DocTypeEvent, nil)
	r = deactivateDocument(r)

	assert.Equal(DocStatusInactive, r.Status)
}

func TestActivateRecord(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := newDocument(DocTypeEvent, nil)
	r = deactivateDocument(r)
	r = activateDocument(r)

	assert.Equal(DocStatusActive, r.Status)
}

func TestUpdateRecordUser(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	u := uuid.New()
	r := newDocument(DocTypeEvent, &u)

	u = uuid.New()
	r = updateDocuemntUser(r, &u)

	assert.Equal(&u, r.User)
}

func BenchmarkNewRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newDocument(DocTypeEvent, nil)
	}
}

func BenchmarkUpdateRecord(b *testing.B) {
	r := newDocument(DocTypeEvent, nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateDocuemnt(r)
	}
}

func BenchmarkDeactivateRecord(b *testing.B) {
	r := newDocument(DocTypeEvent, nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		deactivateDocument(r)
	}
}

func BenchmarkActivateRecord(b *testing.B) {
	r := newDocument(DocTypeEvent, nil)
	r = deactivateDocument(r)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		activateDocument(r)
	}
}

func BenchmarkUpdateRecordUser(b *testing.B) {
	u := uuid.New()
	r := newDocument(DocTypeEvent, &u)
	u = uuid.New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateDocuemntUser(r, &u)
	}
}
