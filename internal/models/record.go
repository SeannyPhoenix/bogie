package models

import (
	"time"

	"github.com/google/uuid"
)

type Record struct {
	Id        uuid.UUID  `json:"id" dynamodbav:"id"`
	Type      string     `json:"type" dynamodbav:"t"`
	Status    string     `json:"status" dynamodbav:"s"`
	CreatedAt time.Time  `json:"createdAt" dynamodbav:"ca"`
	UpdatedAt time.Time  `json:"updatedAt" dynamodbav:"ua"`
	User      *uuid.UUID `json:"user,omitempty" dynamodbav:"u,omitempty"`
}

func newRecord(t string, user *uuid.UUID) Record {
	now := time.Now()
	return Record{
		Id:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Status:    DocStatusActive,
		Type:      t,
		User:      user,
	}
}

func deactivateRecord(r Record) Record {
	r.Status = DocStatusInactive
	r.UpdatedAt = time.Now()
	return r
}

func activateRecord(r Record) Record {
	r.Status = DocStatusActive
	r.UpdatedAt = time.Now()
	return r
}

func updateRecord(r Record) Record {
	r.UpdatedAt = time.Now()
	return r
}

func updateRecordUser(r Record, user *uuid.UUID) Record {
	r.User = user
	r.UpdatedAt = time.Now()
	return r
}
