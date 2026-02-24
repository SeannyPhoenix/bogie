package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/seannyphoenix/bogie/pkg/models"
)

func main() {
	// Example 1: US Car
	fmt.Println("=== Example 1: US Car ===")
	car := models.NewVehicle(models.VehicleTypeCar, "transit-agency-1")
	car.Make = "Toyota"
	car.Model = "Camry"
	car.Year = 2023
	car.Capacity = 5

	vin := models.NewPermanentRegistration("1HGBH41JXMN109186", "US", "VIN")
	car.AddRegistration(vin)

	issued := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)
	expiry := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	plate := models.NewTemporaryRegistration("ABC-1234", "California, US", "License Plate", &issued, &expiry)
	car.AddRegistration(plate)

	printJSON(car)

	// Example 2: Transit Bus
	fmt.Println("\n=== Example 2: Transit Bus ===")
	bus := models.NewVehicle(models.VehicleTypeBus, "metro-transit")
	bus.Make = "New Flyer"
	bus.Model = "XD40"
	bus.Year = 2021
	bus.Capacity = 40

	busVin := models.NewPermanentRegistration("5FZACABE5GJ123456", "US", "VIN")
	bus.AddRegistration(busVin)

	fleet := models.NewOperationalRegistration("2145", "Fleet Number")
	bus.AddRegistration(fleet)

	printJSON(bus)

	// Example 3: Subway Car
	fmt.Println("\n=== Example 3: Subway Car ===")
	traincar := models.NewVehicle(models.VehicleTypeTrainCar, "mta-nyc")
	traincar.Make = "Kawasaki"
	traincar.Model = "R211"
	traincar.Year = 2022
	traincar.Capacity = 180

	serial := models.NewPermanentRegistration("R211-4023", "US", "Serial Number")
	traincar.AddRegistration(serial)

	carNum := models.NewOperationalRegistration("4023", "Car Number")
	traincar.AddRegistration(carNum)

	printJSON(traincar)

	// Example 4: Airplane
	fmt.Println("\n=== Example 4: Commercial Airplane ===")
	airplane := models.NewVehicle(models.VehicleTypeAirplane, "united-airlines")
	airplane.Make = "Boeing"
	airplane.Model = "737-800"
	airplane.Year = 2019
	airplane.Capacity = 175

	airSerial := models.NewPermanentRegistration("42560", "US", "Manufacturer Serial Number")
	airplane.AddRegistration(airSerial)

	tail := models.NewPermanentRegistration("N27733", "US (FAA)", "Tail Number")
	airplane.AddRegistration(tail)

	airFleet := models.NewOperationalRegistration("733", "Fleet Number")
	airplane.AddRegistration(airFleet)

	printJSON(airplane)
}

func printJSON(v *models.Vehicle) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(string(data))
}
