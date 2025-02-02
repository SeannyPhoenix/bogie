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

	r := newRecord(DocTypeEvent, nil)
	r = updateRecord(r)

	ca := r.CreatedAt

	assert.NotEqual(ca, r.UpdatedAt)
}

func TestDeactivateRecord(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := newRecord(DocTypeEvent, nil)
	r = deactivateRecord(r)

	assert.Equal(DocStatusInactive, r.Status)
}

func TestActivateRecord(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r := newRecord(DocTypeEvent, nil)
	r = deactivateRecord(r)
	r = activateRecord(r)

	assert.Equal(DocStatusActive, r.Status)
}

func TestUpdateRecordUser(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	u := uuid.New()
	r := newRecord(DocTypeEvent, &u)

	u = uuid.New()
	r = updateRecordUser(r, &u)

	assert.Equal(&u, r.User)
}

func BenchmarkNewRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newRecord(DocTypeEvent, nil)
	}
}

func BenchmarkUpdateRecord(b *testing.B) {
	r := newRecord(DocTypeEvent, nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateRecord(r)
	}
}

func BenchmarkDeactivateRecord(b *testing.B) {
	r := newRecord(DocTypeEvent, nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		deactivateRecord(r)
	}
}

func BenchmarkActivateRecord(b *testing.B) {
	r := newRecord(DocTypeEvent, nil)
	r = deactivateRecord(r)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		activateRecord(r)
	}
}

func BenchmarkUpdateRecordUser(b *testing.B) {
	u := uuid.New()
	r := newRecord(DocTypeEvent, &u)
	u = uuid.New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateRecordUser(r, &u)
	}
}
