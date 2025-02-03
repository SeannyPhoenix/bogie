package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Document

	Agency        string     `json:"agency,omitempty"`
	Route         string     `json:"route,omitempty"`
	Trip          string     `json:"trip,omitempty"`
	UnitID        string     `json:"unitID,omitempty"`
	UnitCount     *int       `json:"unitCount,omitempty"`
	UnitPosition  *int       `json:"unitPosition,omitempty"`
	DepartureStop string     `json:"departureStop,omitempty"`
	ArrivalStop   string     `json:"arrivalStop,omitempty"`
	DepartureTime *time.Time `json:"departureTime,omitempty"`
	ArrivalTime   *time.Time `json:"arrivalTime,omitempty"`
	Notes         []string   `json:"notes,omitempty"`
}

func NewEventDocuemnt(user *uuid.UUID) Document {
	return newDocument(DocTypeEvent, user)
}
