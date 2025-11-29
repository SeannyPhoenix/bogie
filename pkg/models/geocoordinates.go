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
	"errors"
	"fmt"

	"github.com/seannyphoenix/bogie/pkg/validate"
)

var (
	ErrInvalidLatitude  = errors.New("invalid latitude: must be between -90 and 90")
	ErrInvalidLongitude = errors.New("invalid longitude: must be between -180 and 180")
)

// GeoCoordinates represents a geographic location with latitude and longitude.
//
// The zero value is invalid and should not be used directly.
// Use NewGeoCoordinates to create valid instances.
//
// Latitude must be between -90 and 90 degrees (inclusive).
// Longitude must be between -180 and 180 degrees (inclusive).
//
// JSON representation: {"lat":45.0,"lon":-93.0}
// Invalid coordinates marshal to JSON null.
type GeoCoordinates struct {
	latitude  float64
	longitude float64

	valid bool
}

// NewGeoCoordinates creates a new GeoCoordinates instance with the given latitude and longitude.
// It returns a partial GeoCoordinates instance and an error if the latitude or longitude are out of valid ranges.
func NewGeoCoordinates(lat, lon float64) (GeoCoordinates, error) {
	var c GeoCoordinates

	if lat < -90 || lat > 90 {
		return c, ErrInvalidLatitude
	}
	c.latitude = lat

	if lon < -180 || lon > 180 {
		return c, ErrInvalidLongitude
	}
	c.longitude = lon

	c.valid = true
	return c, nil
}

// MustNewGeoCoordinates is like [NewGeoCoordinates] but panics if the
// latitude or longitude are out of valid ranges.
func MustNewGeoCoordinates(lat, lon float64) GeoCoordinates {
	c, err := NewGeoCoordinates(lat, lon)
	if err != nil {
		panic(fmt.Sprintf("must new coordinates: %v", err))
	}
	return c
}

// Lat returns the latitude of the GeoCoordinates.
func (c GeoCoordinates) Lat() float64 {
	return c.latitude
}

// Lon returns the longitude of the GeoCoordinates.
func (c GeoCoordinates) Lon() float64 {
	return c.longitude
}

// LatLon returns the latitude and longitude of the GeoCoordinates, in that order.
func (c GeoCoordinates) LatLon() (float64, float64) {
	return c.latitude, c.longitude
}

// IsValid returns true if the GeoCoordinates was created successfully via [NewGeoCoordinates].
// The zero value of GeoCoordinates is invalid.
func (c GeoCoordinates) IsValid() bool {
	return c.valid
}

type geoCoordinatesJSON struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// MarshalJSON implements json.Marshaler.
// An invalid GeoCoordinates will be marshaled to JSON null.
func (c GeoCoordinates) MarshalJSON() ([]byte, error) {
	if !c.valid {
		return json.Marshal(nil)
	}

	m := geoCoordinatesJSON{
		Lat: c.latitude,
		Lon: c.longitude,
	}

	return json.Marshal(m)
}

// UnmarshalJSON implements json.Unmarshaler.
// A JSON null will be unmarshaled to an invalid GeoCoordinates.
// Otherwise, the latitude and longitude will be validated.
// If they are out of valid ranges, an error will be returned.
func (c *GeoCoordinates) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*c = GeoCoordinates{}
		return nil
	}

	err := validate.RequiredFields(data, "lat", "lon")
	if err != nil {
		return fmt.Errorf("unmarshal geo coordinates: %w", err)
	}

	// Now unmarshal into the typed struct
	var m geoCoordinatesJSON
	if err := json.Unmarshal(data, &m); err != nil {
		return fmt.Errorf("unmarshal geo coordinates: %w", err)
	}

	newC, err := NewGeoCoordinates(m.Lat, m.Lon)
	if err != nil {
		return fmt.Errorf("unmarshal geo coordinates: %w", err)
	}

	*c = newC
	return nil
}
