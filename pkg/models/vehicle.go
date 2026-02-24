package models

import (
	"time"

	"github.com/google/uuid"
)

// VehicleType represents the category of vehicle
type VehicleType string

const (
	VehicleTypeCar        VehicleType = "car"
	VehicleTypeBus        VehicleType = "bus"
	VehicleTypeTrainCar   VehicleType = "traincar"
	VehicleTypeSeaVessel  VehicleType = "sea_vessel"
	VehicleTypeAirplane   VehicleType = "airplane"
	VehicleTypeHelicopter VehicleType = "helicopter"
)

// RegistrationType distinguishes between permanent identifiers and temporary registrations
type RegistrationType string

const (
	// Permanent identifiers that stay with the vehicle for its lifetime
	RegistrationTypePermanent RegistrationType = "permanent" // VIN, serial number, hull number, tail number

	// Temporary registrations that can change (license plates in US, call signs, etc.)
	RegistrationTypeTemporary RegistrationType = "temporary"

	// Operational identifiers assigned by operators (subway car numbers, fleet numbers)
	RegistrationTypeOperational RegistrationType = "operational"
)

// Registration represents a vehicle identifier/registration
type Registration struct {
	Type         RegistrationType `json:"type"`
	Value        string           `json:"value"`
	Jurisdiction string           `json:"jurisdiction,omitempty"` // Country, state, or authority
	IssuedDate   *time.Time       `json:"issued_date,omitempty"`
	ExpiryDate   *time.Time       `json:"expiry_date,omitempty"`
	Description  string           `json:"description,omitempty"` // e.g., "VIN", "License Plate", "Tail Number"
}

// Vehicle represents a physical vehicle with various identifiers
type Vehicle struct {
	// Primary UUID identifier (v7 recommended for time-ordered sorting)
	ID uuid.UUID `json:"id"`

	// Vehicle type
	Type VehicleType `json:"type"`

	// Agency or operator ID
	AgencyID string `json:"agency_id"`

	// All registrations/identifiers for this vehicle
	// Typically includes at least one permanent identifier (VIN, serial number, etc.)
	// May include temporary registrations (license plates, etc.)
	Registrations []Registration `json:"registrations,omitempty"`

	// Vehicle details
	Make     string `json:"make,omitempty"`
	Model    string `json:"model,omitempty"`
	Year     int    `json:"year,omitempty"`
	Color    string `json:"color,omitempty"`
	Capacity int    `json:"capacity,omitempty"` // passenger capacity

	// Timestamps
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	RetiredAt *time.Time `json:"retired_at,omitempty"`
}

// GetPermanentID returns the first permanent identifier (VIN, serial number, etc.)
func (v *Vehicle) GetPermanentID() *Registration {
	for i := range v.Registrations {
		if v.Registrations[i].Type == RegistrationTypePermanent {
			return &v.Registrations[i]
		}
	}
	return nil
}

// GetCurrentRegistration returns the most recent temporary registration (license plate, etc.)
func (v *Vehicle) GetCurrentRegistration() *Registration {
	var latest *Registration
	for i := range v.Registrations {
		if v.Registrations[i].Type == RegistrationTypeTemporary {
			if latest == nil ||
				(v.Registrations[i].IssuedDate != nil &&
					latest.IssuedDate != nil &&
					v.Registrations[i].IssuedDate.After(*latest.IssuedDate)) {
				latest = &v.Registrations[i]
			}
		}
	}
	return latest
}

// GetOperationalID returns the operational identifier (fleet number, car number, etc.)
func (v *Vehicle) GetOperationalID() *Registration {
	for i := range v.Registrations {
		if v.Registrations[i].Type == RegistrationTypeOperational {
			return &v.Registrations[i]
		}
	}
	return nil
}
