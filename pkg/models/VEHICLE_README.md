# Vehicle Type System

This package provides a flexible vehicle type system that handles various vehicle types and their different registration/identification requirements.

## Overview

The system is designed to handle the complexity of vehicle identification across different jurisdictions and vehicle types. Key features:

- **UUID v7 identifiers**: Each vehicle gets a time-ordered UUID v7 for consistent, sortable identification
- **Multiple registration types**: Permanent (VIN, serial numbers), Temporary (license plates), and Operational (fleet numbers, car numbers)
- **Flexible registration tracking**: Supports multiple registrations per vehicle with jurisdiction and validity tracking
- **Type-safe vehicle categories**: Enumerated vehicle types for cars, buses, trains, aircraft, vessels, and helicopters

## Vehicle Types

- `VehicleTypeCar` - Automobiles
- `VehicleTypeBus` - Buses
- `VehicleTypeTrainCar` - Railway cars/subway cars
- `VehicleTypeSeaVessel` - Boats, ships, ferries
- `VehicleTypeAirplane` - Fixed-wing aircraft
- `VehicleTypeHelicopter` - Rotary-wing aircraft

## Registration Types

### Permanent Identifiers
Identifiers that stay with the vehicle for its entire lifetime:
- Vehicle Identification Numbers (VIN)
- Manufacturer serial numbers
- Hull Identification Numbers (boats)
- Aircraft serial numbers
- UK license plates (they stay with the vehicle)

### Temporary Registrations
Registrations that can change over time:
- US license plates (stay with owner, not vehicle)
- Registration renewals
- Temporary permits

### Operational Identifiers
Identifiers assigned by operators:
- Fleet numbers
- Car numbers (subway/train)
- Call signs
- Vessel names

## Usage Examples

### US Car with VIN and License Plate

```go
import (
    "time"
    "github.com/seannyphoenix/bogie/pkg/models"
)

// Create a new car
car := models.NewVehicle(models.VehicleTypeCar, "transit-agency-1")
car.Make = "Toyota"
car.Model = "Camry"
car.Year = 2023
car.Color = "Silver"
car.Capacity = 5

// Add VIN (permanent identifier)
vin := models.NewPermanentRegistration(
    "1HGBH41JXMN109186",
    "US",
    "VIN",
)
car.AddRegistration(vin)

// Add license plate (temporary - changes when sold)
issued := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)
expiry := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
plate := models.NewTemporaryRegistration(
    "ABC-1234",
    "California, US",
    "License Plate",
    &issued,
    &expiry,
)
car.AddRegistration(plate)

// Retrieve identifiers
permanentID := car.GetPermanentID() // Returns VIN
currentPlate := car.GetCurrentRegistration() // Returns license plate
```

### UK Car with Permanent License Plate

In the UK, license plates stay with vehicles for life, making them permanent identifiers:

```go
car := models.NewVehicle(models.VehicleTypeCar, "uk-transit-1")
car.Make = "Vauxhall"
car.Model = "Astra"
car.Year = 2022

// UK registration number is permanent
plate := models.NewPermanentRegistration(
    "AB12 CDE",
    "UK",
    "Registration Number",
)
car.AddRegistration(plate)

// VIN is also permanent
vin := models.NewPermanentRegistration(
    "W0L0AHL0855123456",
    "UK",
    "VIN",
)
car.AddRegistration(vin)
```

### Transit Bus with Multiple Identifiers

```go
bus := models.NewVehicle(models.VehicleTypeBus, "metro-transit")
bus.Make = "New Flyer"
bus.Model = "XD40"
bus.Year = 2021
bus.Capacity = 40

// VIN (permanent)
vin := models.NewPermanentRegistration(
    "5FZACABE5GJ123456",
    "US",
    "VIN",
)
bus.AddRegistration(vin)

// Fleet number (operational - used by transit agency)
fleet := models.NewOperationalRegistration(
    "2145",
    "Fleet Number",
)
bus.AddRegistration(fleet)

// License plate (temporary)
issued := time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC)
expiry := time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC)
plate := models.NewTemporaryRegistration(
    "TRN-2145",
    "New York, US",
    "Commercial Plate",
    &issued,
    &expiry,
)
bus.AddRegistration(plate)

// Access different identifiers
fleetNum := bus.GetOperationalID().Value // "2145"
```

### Subway/Train Car

```go
traincar := models.NewVehicle(models.VehicleTypeTrainCar, "mta-nyc")
traincar.Make = "Kawasaki"
traincar.Model = "R211"
traincar.Year = 2022
traincar.Capacity = 180

// Serial number (permanent)
serial := models.NewPermanentRegistration(
    "R211-4023",
    "US",
    "Serial Number",
)
traincar.AddRegistration(serial)

// Car number (operational - displayed on the car)
carNum := models.NewOperationalRegistration(
    "4023",
    "Car Number",
)
traincar.AddRegistration(carNum)
```

### Commercial Airplane

```go
airplane := models.NewVehicle(models.VehicleTypeAirplane, "united-airlines")
airplane.Make = "Boeing"
airplane.Model = "737-800"
airplane.Year = 2019
airplane.Capacity = 175

// Manufacturer serial number (permanent)
serial := models.NewPermanentRegistration(
    "42560",
    "US",
    "Manufacturer Serial Number",
)
airplane.AddRegistration(serial)

// Tail number/N-number (permanent in US)
tail := models.NewPermanentRegistration(
    "N27733",
    "US (FAA)",
    "Tail Number",
)
airplane.AddRegistration(tail)

// Fleet number (operational)
fleet := models.NewOperationalRegistration(
    "733",
    "Fleet Number",
)
airplane.AddRegistration(fleet)
```

### Ferry/Sea Vessel

```go
vessel := models.NewVehicle(models.VehicleTypeSeaVessel, "staten-island-ferry")
vessel.Make = "Eastern Shipbuilding"
vessel.Model = "Ferry"
vessel.Year = 2015
vessel.Capacity = 4400

// Hull Identification Number (permanent)
hin := models.NewPermanentRegistration(
    "ESG12345A616",
    "US",
    "Hull Identification Number",
)
vessel.AddRegistration(hin)

// Coast Guard official number (permanent)
official := models.NewPermanentRegistration(
    "1234567",
    "US (USCG)",
    "Official Number",
)
vessel.AddRegistration(official)

// Vessel name (operational)
name := models.NewOperationalRegistration(
    "Spirit of America",
    "Vessel Name",
)
vessel.AddRegistration(name)
```

### Helicopter

```go
helicopter := models.NewVehicle(models.VehicleTypeHelicopter, "air-ambulance-1")
helicopter.Make = "Eurocopter"
helicopter.Model = "EC135"
helicopter.Year = 2020
helicopter.Capacity = 6

// Serial number (permanent)
serial := models.NewPermanentRegistration(
    "EC135-0987",
    "US",
    "Serial Number",
)
helicopter.AddRegistration(serial)

// Registration number (permanent)
reg := models.NewPermanentRegistration(
    "N911AE",
    "US (FAA)",
    "Registration Number",
)
helicopter.AddRegistration(reg)

// Call sign (operational)
callsign := models.NewOperationalRegistration(
    "LifeFlight-1",
    "Call Sign",
)
helicopter.AddRegistration(callsign)
```

## Handling Registration Changes

When a vehicle's registration changes (e.g., license plate renewal or transfer):

```go
car := models.NewVehicle(models.VehicleTypeCar, "agency1")

// Add VIN
vin := models.NewPermanentRegistration("1HGBH41JXMN109186", "US", "VIN")
car.AddRegistration(vin)

// First plate
issued1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
expiry1 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
plate1 := models.NewTemporaryRegistration(
    "OLD-1234",
    "California, US",
    "License Plate",
    &issued1,
    &expiry1,
)
car.AddRegistration(plate1)

// New plate after sale/transfer
issued2 := time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)
expiry2 := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
plate2 := models.NewTemporaryRegistration(
    "NEW-5678",
    "California, US",
    "License Plate",
    &issued2,
    &expiry2,
)
car.AddRegistration(plate2)

// GetCurrentRegistration returns the most recent temporary registration
current := car.GetCurrentRegistration() // Returns "NEW-5678"
```

## UUID v7 Benefits

The system uses UUID v7 for vehicle identifiers, which provides:

- **Time-ordered**: UUIDs can be sorted chronologically
- **Globally unique**: No coordination needed between systems
- **Database-friendly**: Better for indexing than random UUIDs
- **Future-proof**: 128-bit identifier space

```go
v1 := models.NewVehicle(models.VehicleTypeCar, "agency1")
time.Sleep(2 * time.Millisecond)
v2 := models.NewVehicle(models.VehicleTypeCar, "agency1")

// v1.ID < v2.ID (time-ordered)
```

## Retirement Tracking

Track when vehicles are retired from service:

```go
vehicle := models.NewVehicle(models.VehicleTypeBus, "agency1")

// ... use the vehicle ...

// Retire the vehicle
retiredTime := time.Now()
vehicle.RetiredAt = &retiredTime
```

## JSON Serialization

All types are JSON-serializable:

```go
import "encoding/json"

vehicle := models.NewVehicle(models.VehicleTypeCar, "agency1")
// ... configure vehicle ...

jsonData, err := json.Marshal(vehicle)
// Use in APIs, store in databases, etc.
```

## Design Rationale

### Why separate Registration types?

Different jurisdictions and contexts require different approaches:
- US cars: VIN is permanent, license plate changes with ownership
- UK cars: Registration number is permanent
- Aircraft: Multiple permanent identifiers (serial, tail number)
- Transit: Operational identifiers (fleet numbers) separate from legal registrations

### Why store all registrations?

Maintaining a history of registrations allows:
- Tracking license plate changes over time
- Associating vehicles with historical data
- Audit trails for regulatory compliance
- Supporting different identification schemes simultaneously

### Why UUID v7?

- Time-ordered UUIDs improve database performance
- No central coordination required
- Future-proof identifier space
- Compatible with distributed systems
