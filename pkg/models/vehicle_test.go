package models

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewVehicle(t *testing.T) {
	vehicle := NewVehicle(VehicleTypeCar, "agency123")

	assert.NotEqual(t, uuid.Nil, vehicle.ID)
	assert.Equal(t, VehicleTypeCar, vehicle.Type)
	assert.Equal(t, "agency123", vehicle.AgencyID)
	assert.NotNil(t, vehicle.Registrations)
	assert.False(t, vehicle.CreatedAt.IsZero())
	assert.False(t, vehicle.UpdatedAt.IsZero())
}

func TestUUIDv7Ordering(t *testing.T) {
	// Create multiple vehicles and verify UUID v7 time ordering
	v1 := NewVehicle(VehicleTypeCar, "agency1")
	time.Sleep(2 * time.Millisecond)
	v2 := NewVehicle(VehicleTypeCar, "agency1")
	time.Sleep(2 * time.Millisecond)
	v3 := NewVehicle(VehicleTypeCar, "agency1")

	// UUID v7 should be time-ordered
	assert.True(t, v1.ID.String() < v2.ID.String())
	assert.True(t, v2.ID.String() < v3.ID.String())
}

func TestCarWithVINAndLicensePlate(t *testing.T) {
	// Example: US car with VIN and license plate
	car := NewVehicle(VehicleTypeCar, "transit-agency-1")
	car.Make = "Toyota"
	car.Model = "Camry"
	car.Year = 2023

	// Add VIN (permanent identifier)
	vin := NewPermanentRegistration(
		"1HGBH41JXMN109186",
		"US",
		"VIN",
	)
	car.AddRegistration(vin)

	// Add license plate (temporary registration)
	issued := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)
	expiry := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	plate := NewTemporaryRegistration(
		"ABC-1234",
		"California, US",
		"License Plate",
		&issued,
		&expiry,
	)
	car.AddRegistration(plate)

	assert.Equal(t, 2, len(car.Registrations))
	assert.NotNil(t, car.GetPermanentID())
	assert.Equal(t, "1HGBH41JXMN109186", car.GetPermanentID().Value)
	assert.NotNil(t, car.GetCurrentRegistration())
	assert.Equal(t, "ABC-1234", car.GetCurrentRegistration().Value)
}

func TestUKCarWithPermanentLicensePlate(t *testing.T) {
	// Example: UK car where license plate is permanent
	car := NewVehicle(VehicleTypeCar, "uk-transit-1")
	car.Make = "Vauxhall"
	car.Model = "Astra"
	car.Year = 2022

	// In UK, license plate stays with vehicle for life (permanent)
	plate := NewPermanentRegistration(
		"AB12 CDE",
		"UK",
		"Registration Number",
	)
	car.AddRegistration(plate)

	// VIN is also permanent
	vin := NewPermanentRegistration(
		"W0L0AHL0855123456",
		"UK",
		"VIN",
	)
	car.AddRegistration(vin)

	assert.Equal(t, 2, len(car.Registrations))
	// Both are permanent, so GetPermanentID returns the first one
	assert.NotNil(t, car.GetPermanentID())
}

func TestBusWithFleetNumber(t *testing.T) {
	// Example: Transit bus with multiple identifiers
	bus := NewVehicle(VehicleTypeBus, "metro-transit")
	bus.Make = "New Flyer"
	bus.Model = "XD40"
	bus.Year = 2021
	bus.Capacity = 40

	// VIN
	vin := NewPermanentRegistration(
		"5FZACABE5GJ123456",
		"US",
		"VIN",
	)
	bus.AddRegistration(vin)

	// Fleet number (operational identifier)
	fleet := NewOperationalRegistration(
		"2145",
		"Fleet Number",
	)
	bus.AddRegistration(fleet)

	// License plate
	issued := time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC)
	expiry := time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC)
	plate := NewTemporaryRegistration(
		"TRN-2145",
		"New York, US",
		"Commercial Plate",
		&issued,
		&expiry,
	)
	bus.AddRegistration(plate)

	assert.Equal(t, 3, len(bus.Registrations))
	assert.NotNil(t, bus.GetPermanentID())
	assert.NotNil(t, bus.GetOperationalID())
	assert.Equal(t, "2145", bus.GetOperationalID().Value)
	assert.NotNil(t, bus.GetCurrentRegistration())
}

func TestSubwayCarWithCarNumber(t *testing.T) {
	// Example: Subway/train car with car number
	traincar := NewVehicle(VehicleTypeTrainCar, "mta-nyc")
	traincar.Make = "Kawasaki"
	traincar.Model = "R211"
	traincar.Year = 2022
	traincar.Capacity = 180

	// Serial number (permanent)
	serial := NewPermanentRegistration(
		"R211-4023",
		"US",
		"Serial Number",
	)
	traincar.AddRegistration(serial)

	// Car number (operational)
	carNum := NewOperationalRegistration(
		"4023",
		"Car Number",
	)
	traincar.AddRegistration(carNum)

	assert.Equal(t, 2, len(traincar.Registrations))
	assert.NotNil(t, traincar.GetPermanentID())
	assert.NotNil(t, traincar.GetOperationalID())
	assert.Equal(t, "4023", traincar.GetOperationalID().Value)
}

func TestAirplaneWithTailNumber(t *testing.T) {
	// Example: Commercial airplane
	airplane := NewVehicle(VehicleTypeAirplane, "united-airlines")
	airplane.Make = "Boeing"
	airplane.Model = "737-800"
	airplane.Year = 2019
	airplane.Capacity = 175

	// Serial number (permanent - manufacturer serial)
	serial := NewPermanentRegistration(
		"42560",
		"US",
		"Manufacturer Serial Number",
	)
	airplane.AddRegistration(serial)

	// Tail number/registration (in some jurisdictions this is permanent)
	tail := NewPermanentRegistration(
		"N27733",
		"US (FAA)",
		"Tail Number",
	)
	airplane.AddRegistration(tail)

	// Fleet number (operational)
	fleet := NewOperationalRegistration(
		"733",
		"Fleet Number",
	)
	airplane.AddRegistration(fleet)

	assert.Equal(t, 3, len(airplane.Registrations))
	assert.NotNil(t, airplane.GetPermanentID())
	assert.NotNil(t, airplane.GetOperationalID())
}

func TestSeaVesselWithHullNumber(t *testing.T) {
	// Example: Ferry boat
	vessel := NewVehicle(VehicleTypeSeaVessel, "staten-island-ferry")
	vessel.Make = "Eastern Shipbuilding"
	vessel.Model = "Ferry"
	vessel.Year = 2015
	vessel.Capacity = 4400

	// Hull Identification Number (permanent)
	hin := NewPermanentRegistration(
		"ESG12345A616",
		"US",
		"Hull Identification Number",
	)
	vessel.AddRegistration(hin)

	// Official number (permanent - Coast Guard)
	official := NewPermanentRegistration(
		"1234567",
		"US (USCG)",
		"Official Number",
	)
	vessel.AddRegistration(official)

	// Vessel name (operational)
	name := NewOperationalRegistration(
		"Spirit of America",
		"Vessel Name",
	)
	vessel.AddRegistration(name)

	assert.Equal(t, 3, len(vessel.Registrations))
	assert.NotNil(t, vessel.GetPermanentID())
	assert.NotNil(t, vessel.GetOperationalID())
}

func TestHelicopterWithRegistration(t *testing.T) {
	// Example: Helicopter
	helicopter := NewVehicle(VehicleTypeHelicopter, "air-ambulance-1")
	helicopter.Make = "Eurocopter"
	helicopter.Model = "EC135"
	helicopter.Year = 2020
	helicopter.Capacity = 6

	// Serial number
	serial := NewPermanentRegistration(
		"EC135-0987",
		"US",
		"Serial Number",
	)
	helicopter.AddRegistration(serial)

	// Registration number (tail number)
	reg := NewPermanentRegistration(
		"N911AE",
		"US (FAA)",
		"Registration Number",
	)
	helicopter.AddRegistration(reg)

	// Call sign (operational)
	callsign := NewOperationalRegistration(
		"LifeFlight-1",
		"Call Sign",
	)
	helicopter.AddRegistration(callsign)

	assert.Equal(t, 3, len(helicopter.Registrations))
	assert.NotNil(t, helicopter.GetPermanentID())
	assert.NotNil(t, helicopter.GetOperationalID())
	assert.Equal(t, "LifeFlight-1", helicopter.GetOperationalID().Value)
}

func TestMultipleLicensePlateChanges(t *testing.T) {
	// Example: Car with multiple license plate changes over time
	car := NewVehicle(VehicleTypeCar, "agency1")

	// VIN
	vin := NewPermanentRegistration(
		"1HGBH41JXMN109186",
		"US",
		"VIN",
	)
	car.AddRegistration(vin)

	// First plate
	issued1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	expiry1 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	plate1 := NewTemporaryRegistration(
		"OLD-1234",
		"California, US",
		"License Plate",
		&issued1,
		&expiry1,
	)
	car.AddRegistration(plate1)

	// Second plate (renewal/change)
	issued2 := time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)
	expiry2 := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
	plate2 := NewTemporaryRegistration(
		"NEW-5678",
		"California, US",
		"License Plate",
		&issued2,
		&expiry2,
	)
	car.AddRegistration(plate2)

	// GetCurrentRegistration should return the most recent one
	current := car.GetCurrentRegistration()
	assert.NotNil(t, current)
	assert.Equal(t, "NEW-5678", current.Value)
}

func TestRetiredVehicle(t *testing.T) {
	vehicle := NewVehicle(VehicleTypeBus, "agency1")

	assert.Nil(t, vehicle.RetiredAt)

	// Retire the vehicle
	retiredTime := time.Now()
	vehicle.RetiredAt = &retiredTime

	assert.NotNil(t, vehicle.RetiredAt)
	assert.False(t, vehicle.RetiredAt.IsZero())
}
