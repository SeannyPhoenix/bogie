package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Document

	EventType string    `json:"eventType"`
	Timestamp time.Time `json:"timestamp"`
	Exact     bool      `json:"exact"`

	Location string   `json:"location,omitempty"`
	Agency   string   `json:"agency,omitempty"`
	Route    string   `json:"route,omitempty"`
	Vehicle  *Vehicle `json:"vehicle,omitempty"`
	Run      string   `json:"run,omitempty"`
}

func NewEventDocuemnt(user *uuid.UUID) Event {
	return Event{
		Document: NewDocument(DocTypeEvent, user),
	}
}
