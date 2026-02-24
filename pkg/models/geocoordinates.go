// Package models provides domain model types for the Bogie transit tracking system.
//
// This package contains core value objects and entities used throughout the application,
// particularly for geographic data modeling. Types in this package are designed to be
// immutable after construction and provide validation at creation time.
//
// Design Philosophy:
//   - Value objects are created through constructor functions that validate invariants
//   - Zero values are invalid and should never be used directly
//   - JSON marshaling/unmarshaling maintains validation constraints
//   - Types are optimized for correctness over convenience
package models

import (
	"encoding/json"
	"fmt"
)

type GeoCoordinates struct {
	latitude  Latitude
	longitude Longitude

	valid bool
}

func NewGeoCoordinates(lat, lon float64) (GeoCoordinates, error) {
	var c GeoCoordinates

	latT, err := ParseLatitude(lat)
	if err != nil {
		return c, err
	}
	c.latitude = latT

	lonT, err := ParseLongitude(lon)
	if err != nil {
		return c, err
	}
	c.longitude = lonT

	c.valid = true
	return c, nil
}

func MustNewGeoCoordinates(lat, lon float64) GeoCoordinates {
	c, err := NewGeoCoordinates(lat, lon)
	if err != nil {
		panic(fmt.Sprintf("must new coordinates: %v", err))
	}
	return c
}

func (c GeoCoordinates) Lat() float64 { return float64(c.latitude) }

func (c GeoCoordinates) Lon() float64 { return float64(c.longitude) }

func (c GeoCoordinates) LatLon() (float64, float64) { return float64(c.latitude), float64(c.longitude) }

func (c GeoCoordinates) IsValid() bool { return c.valid }

type geoCoordinatesJSON struct {
	Lat *Latitude  `json:"lat"`
	Lon *Longitude `json:"lon"`
}

func (c GeoCoordinates) MarshalJSON() ([]byte, error) {
	if !c.valid {
		return json.Marshal(nil)
	}

	m := geoCoordinatesJSON{
		Lat: &c.latitude,
		Lon: &c.longitude,
	}

	return json.Marshal(m)
}

func (c *GeoCoordinates) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*c = GeoCoordinates{}
		return nil
	}

	// if err := validate.RequiredFields(data, "lat", "lon"); err != nil {
	// 	return fmt.Errorf("unmarshal geo coordinates: %w", err)
	// }

	var m geoCoordinatesJSON
	if err := json.Unmarshal(data, &m); err != nil {
		return fmt.Errorf("unmarshal geo coordinates: %w", err)
	}

	if m.Lat == nil || m.Lon == nil {
		return fmt.Errorf("unmarshal geo coordinates: missing required fields")
	}

	c.latitude = *m.Lat
	c.longitude = *m.Lon
	c.valid = true
	return nil
}
