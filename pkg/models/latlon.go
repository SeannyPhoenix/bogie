package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	InvalidLatitude  = errors.New("invalid latitude: must be between -90 and 90")
	InvalidLongitude = errors.New("invalid longitude: must be between -180 and 180")
)

type Latitude float64
type Longitude float64

func ParseLatitude(v float64) (Latitude, error) {
	if v < -90 || v > 90 {
		return 0, InvalidLatitude
	}
	return Latitude(v), nil
}

func ParseLongitude(v float64) (Longitude, error) {
	if v < -180 || v > 180 {
		return 0, InvalidLongitude
	}
	return Longitude(v), nil
}

func (l *Latitude) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return InvalidLatitude
	}
	var v float64
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal latitude: %w", err)
	}
	parsed, err := ParseLatitude(v)
	if err != nil {
		return err
	}
	*l = parsed
	return nil
}

func (lo *Longitude) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return InvalidLongitude
	}
	var v float64
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal longitude: %w", err)
	}
	parsed, err := ParseLongitude(v)
	if err != nil {
		return err
	}
	*lo = parsed
	return nil
}
