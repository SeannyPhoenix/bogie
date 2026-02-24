package main

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/seannyphoenix/binarytime/pkg/binarytime"
)

type event struct {
	ID        uuid.UUID       `json:"id"`
	Previous  *uuid.UUID      `json:"previous,omitempty"`
	Begin     string          `json:"begin,omitempty"`
	EventType string          `json:"eventType,omitempty"`
	Timestamp binarytime.Date `json:"timestamp"`
	Location  location        `json:"location"`
	Run       run             `json:"run"`
	Notes     []string        `json:"notes,omitempty"`
}

type location struct {
	Name     string `json:"name,omitempty"`
	Platform string `json:"platform,omitempty"`
	StopID   string `json:"stopId,omitempty"`
	Agency   string `json:"agency,omitempty"`
}

type run struct {
	Route     string `json:"route,omitempty"`
	Run       string `json:"run,omitempty"`
	Operator  string `json:"operator,omitempty"`
	VehicleID string `json:"vehicleId,omitempty"`
	Units     []unit `json:"units,omitempty"`
}

type unit struct {
	ID          string `json:"id,omitempty"`
	Orientation string `json:"orientation,omitempty"`
}

func getUnits(e EventV2CSV) []unit {
	units := []unit{
		{e.Unit1ID, e.Unit1Orientation},
		{e.Unit2ID, e.Unit2Orientation},
		{e.Unit3ID, e.Unit3Orientation},
		{e.Unit4ID, e.Unit4Orientation},
		{e.Unit5ID, e.Unit5Orientation},
		{e.Unit6ID, e.Unit6Orientation},
		{e.Unit7ID, e.Unit7Orientation},
		{e.Unit8ID, e.Unit8Orientation},
		{e.Unit9ID, e.Unit9Orientation},
	}

	c, err := strconv.Atoi(e.Count)
	if err != nil || c < 0 {
		c = 0
	}

	return units[:c]
}

func getNotes(e EventV2CSV) []string {
	var notes []string

	if e.Note0 != "" {
		notes = append(notes, e.Note0)
	}
	if e.Note1 != "" {
		notes = append(notes, e.Note1)
	}
	if e.Note2 != "" {
		notes = append(notes, e.Note2)
	}
	if e.Note3 != "" {
		notes = append(notes, e.Note3)
	}
	if e.Note4 != "" {
		notes = append(notes, e.Note4)
	}

	return notes
}
