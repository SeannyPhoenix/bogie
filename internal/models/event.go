package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Document

	EventType string    `json:"eventType,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Exact     bool      `json:"exact,omitempty"`

	Location string   `json:"location,omitempty"`
	Agency   string   `json:"agency,omitempty"`
	Route    string   `json:"route,omitempty"`
	Vehicle  *Vehicle `json:"vehicle,omitempty"`
	Run      string   `json:"run,omitempty"`

	// Trip         string `json:"trip,omitempty"`
	// UnitID       string `json:"unitID,omitempty"`
	// UnitCount    *int   `json:"unitCount,omitempty"`
	// UnitPosition *int   `json:"unitPosition,omitempty"`
}

func NewEventDocuemnt(user *uuid.UUID) Event {
	return Event{
		Document: NewDocument(DocTypeEvent, user),
	}
}
