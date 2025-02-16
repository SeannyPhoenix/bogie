package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	Id        uuid.UUID  `json:"id" dynamodbav:"id"`
	Type      string     `json:"type" dynamodbav:"t"`
	Status    string     `json:"status" dynamodbav:"s"`
	CreatedAt time.Time  `json:"createdAt" dynamodbav:"ca"`
	UpdatedAt time.Time  `json:"updatedAt" dynamodbav:"ua"`
	User      *uuid.UUID `json:"user,omitempty" dynamodbav:"u,omitempty"`
}

func newDocument(t string, user *uuid.UUID) Document {
	now := time.Now()
	return Document{
		Id:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Status:    DocStatusActive,
		Type:      t,
		User:      user,
	}
}

func deactivateDocument(r Document) Document {
	r.Status = DocStatusInactive
	r.UpdatedAt = time.Now()
	return r
}

func activateDocument(r Document) Document {
	r.Status = DocStatusActive
	r.UpdatedAt = time.Now()
	return r
}

func updateDocuemnt(r Document) Document {
	r.UpdatedAt = time.Now()
	return r
}

func updateDocuemntUser(r Document, user *uuid.UUID) Document {
	r.User = user
	r.UpdatedAt = time.Now()
	return r
}
