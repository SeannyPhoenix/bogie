package models

import (
	"time"

	"github.com/google/uuid"
)

// NewVehicle creates a new Vehicle with a UUID v7 identifier
func NewVehicle(vehicleType VehicleType, agencyID string) *Vehicle {
	now := time.Now()
	return &Vehicle{
		ID:            uuid.Must(uuid.NewV7()), // UUID v7 for time-ordered sorting
		Type:          vehicleType,
		AgencyID:      agencyID,
		Registrations: make([]Registration, 0),
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// AddRegistration adds a new registration to the vehicle
func (v *Vehicle) AddRegistration(reg Registration) {
	v.Registrations = append(v.Registrations, reg)
	v.UpdatedAt = time.Now()
}

// NewPermanentRegistration creates a permanent identifier registration (VIN, serial number, etc.)
func NewPermanentRegistration(value, jurisdiction, description string) Registration {
	return Registration{
		Type:         RegistrationTypePermanent,
		Value:        value,
		Jurisdiction: jurisdiction,
		Description:  description,
	}
}

// NewTemporaryRegistration creates a temporary registration (license plate, etc.)
func NewTemporaryRegistration(value, jurisdiction, description string, issued, expiry *time.Time) Registration {
	return Registration{
		Type:         RegistrationTypeTemporary,
		Value:        value,
		Jurisdiction: jurisdiction,
		Description:  description,
		IssuedDate:   issued,
		ExpiryDate:   expiry,
	}
}

// NewOperationalRegistration creates an operational identifier (fleet number, car number, etc.)
func NewOperationalRegistration(value, description string) Registration {
	return Registration{
		Type:        RegistrationTypeOperational,
		Value:       value,
		Description: description,
	}
}
