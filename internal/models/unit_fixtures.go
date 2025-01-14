package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/internal/db"
	"github.com/seannyphoenix/bogie/internal/fixtures"
)

var (
	defaultUnitNotes = []string{
		"Updated display",
		"",
	}
)

func GetExampleUnit(id uuid.UUID) Unit {
	fixtures.MaybeRefreshUUID(&id)
	t := time.Now().Truncate(time.Second)

	return Unit{
		Id:        id,
		Type:      db.DocTypeUnit,
		Status:    db.StatusActive,
		CreatedAt: &t,
		UpdatedAt: &t,
		Agency:    defaultAgency,
		UnitID:    defaultUnitID,
		Notes:     defaultUnitNotes,
	}
}

func GetExampleUnitArray(count int) []Unit {
	units := make([]Unit, count)
	for i := 0; i < count; i++ {
		units[i] = GetExampleUnit(uuid.Nil)
	}
	return units
}
