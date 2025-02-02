package models

import (
	"time"
)

type Event struct {
	Record

	Agency        string     `json:"agency,omitempty"`        // a
	Route         string     `json:"route,omitempty"`         // r
	Trip          string     `json:"trip,omitempty"`          // tr
	UnitID        string     `json:"unitID,omitempty"`        // u
	UnitCount     *int       `json:"unitCount,omitempty"`     // uc
	UnitPosition  *int       `json:"unitPosition,omitempty"`  // up
	DepartureStop string     `json:"departureStop,omitempty"` // ds
	ArrivalStop   string     `json:"arrivalStop,omitempty"`   // as
	DepartureTime *time.Time `json:"departureTime,omitempty"` // dt
	ArrivalTime   *time.Time `json:"arrivalTime,omitempty"`   // at
	Notes         []string   `json:"notes,omitempty"`         // n
}
